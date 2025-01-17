package actions

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-github/v41/github"

	"gopkg.in/yaml.v2"
)

type EnvOpts struct {
	Client                                  *github.Client
	ServerURL                               string
	PorterToken                             string
	GitRepoOwner, GitRepoName               string
	EnvironmentName                         string
	ProjectID, ClusterID, GitInstallationID uint
}

func SetupEnv(opts *EnvOpts) error {
	// make a best-effort to create a Github environment. this is a non-fatal operation,
	// as the environments API is not enabled for private repositories that don't have
	// github enterprise.
	_, resp, err := opts.Client.Repositories.GetEnvironment(
		context.Background(),
		opts.GitRepoOwner,
		opts.GitRepoName,
		opts.EnvironmentName,
	)

	if resp != nil && resp.StatusCode == http.StatusNotFound {
		opts.Client.Repositories.CreateUpdateEnvironment(
			context.Background(),
			opts.GitRepoOwner,
			opts.GitRepoName,
			opts.EnvironmentName,
			nil,
		)
	}

	// create porter token secret
	err = createGithubSecret(
		opts.Client,
		getPorterTokenSecretName(opts.ProjectID),
		opts.PorterToken,
		opts.GitRepoOwner,
		opts.GitRepoName,
	)

	if err != nil {
		return err
	}

	// get the repository to find the default branch
	repo, _, err := opts.Client.Repositories.Get(
		context.TODO(),
		opts.GitRepoOwner,
		opts.GitRepoName,
	)

	if err != nil {
		return err
	}

	defaultBranch := repo.GetDefaultBranch()

	applyWorkflowYAML, err := getPreviewApplyActionYAML(opts)

	if err != nil {
		return err
	}

	deleteWorkflowYAML, err := getPreviewDeleteActionYAML(opts)

	if err != nil {
		return err
	}

	_, err = commitGithubFile(
		opts.Client,
		fmt.Sprintf("porter_%s_env.yml", strings.ToLower(opts.EnvironmentName)),
		applyWorkflowYAML,
		opts.GitRepoOwner,
		opts.GitRepoName,
		defaultBranch,
		false,
	)

	if err != nil {
		if strings.Contains(err.Error(), "409 Could not create file") {
			// possibly a write-protected branch
			err = createPorterPreviewBranch(opts, defaultBranch)

			if err != nil {
				return fmt.Errorf("write-protected branch %s. Error creating porter-preview branch: %w", defaultBranch, err)
			}

			_, err = commitGithubFile(
				opts.Client,
				fmt.Sprintf("porter_%s_env.yml", strings.ToLower(opts.EnvironmentName)),
				applyWorkflowYAML,
				opts.GitRepoOwner,
				opts.GitRepoName,
				"porter-preview",
				false,
			)

			if err != nil {
				return fmt.Errorf("write-protected branch %s. Error committing to porter-preview branch: %w", defaultBranch, err)
			}

			_, err = commitGithubFile(
				opts.Client,
				fmt.Sprintf("porter_%s_delete_env.yml", strings.ToLower(opts.EnvironmentName)),
				deleteWorkflowYAML,
				opts.GitRepoOwner,
				opts.GitRepoName,
				"porter-preview",
				false,
			)

			if err != nil {
				return fmt.Errorf("write-protected branch %s. Error committing to porter-preview branch: %w", defaultBranch, err)
			}

			pr, _, err := opts.Client.PullRequests.Create(
				context.Background(), opts.GitRepoOwner, opts.GitRepoName, &github.NewPullRequest{
					Title: github.String("Merge Porter preview environment Github Actions workflow files"),
					Base:  github.String(defaultBranch),
					Head:  github.String("porter-preview"),
				},
			)

			if err != nil {
				return err
			}

			return fmt.Errorf("write-protected branch %s. Please merge %s to enable preview environment for your repository", defaultBranch, pr.GetURL())
		}

		return err
	}

	_, err = commitGithubFile(
		opts.Client,
		fmt.Sprintf("porter_%s_delete_env.yml", strings.ToLower(opts.EnvironmentName)),
		deleteWorkflowYAML,
		opts.GitRepoOwner,
		opts.GitRepoName,
		defaultBranch,
		false,
	)

	if err != nil {
		return err
	}

	return err
}

func DeleteEnv(opts *EnvOpts) error {
	// get the repository to find the default branch
	repo, _, err := opts.Client.Repositories.Get(
		context.TODO(),
		opts.GitRepoOwner,
		opts.GitRepoName,
	)

	if err != nil {
		return err
	}

	defaultBranch := repo.GetDefaultBranch()

	// delete GitHub Environment: check that environment exists before deletion

	_, resp, err := opts.Client.Repositories.GetEnvironment(
		context.Background(),
		opts.GitRepoOwner,
		opts.GitRepoName,
		opts.EnvironmentName,
	)

	if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
		_, err = opts.Client.Repositories.DeleteEnvironment(
			context.Background(),
			opts.GitRepoOwner,
			opts.GitRepoName,
			opts.EnvironmentName,
		)

		if err != nil {
			return err
		}
	}

	err = deleteGithubFile(
		opts.Client,
		fmt.Sprintf("porter_%s_env.yml", strings.ToLower(opts.EnvironmentName)),
		opts.GitRepoOwner,
		opts.GitRepoName,
		defaultBranch,
		false,
	)

	if err != nil {
		return err
	}

	return deleteGithubFile(
		opts.Client,
		fmt.Sprintf("porter_%s_delete_env.yml", strings.ToLower(opts.EnvironmentName)),
		opts.GitRepoOwner,
		opts.GitRepoName,
		defaultBranch,
		false,
	)
}

func getPreviewApplyActionYAML(opts *EnvOpts) ([]byte, error) {
	gaSteps := []GithubActionYAMLStep{
		getCheckoutCodeStep(),
		getCreatePreviewEnvStep(
			opts.ServerURL,
			getPorterTokenSecretName(opts.ProjectID),
			opts.ProjectID,
			opts.ClusterID,
			opts.GitInstallationID,
			opts.GitRepoOwner,
			opts.GitRepoName,
			"v0.2.0",
		),
	}

	actionYAML := GithubActionYAML{
		On: map[string]interface{}{
			"workflow_dispatch": map[string]interface{}{
				"inputs": map[string]interface{}{
					"pr_number": map[string]interface{}{
						"description": "Pull request number",
						"type":        "number",
						"required":    true,
					},
					"pr_title": map[string]interface{}{
						"description": "Pull request title",
						"type":        "string",
						"required":    true,
					},
					"pr_branch_from": map[string]interface{}{
						"description": "Pull request head branch",
						"type":        "string",
						"required":    true,
					},
					"pr_branch_into": map[string]interface{}{
						"description": "Pull request base branch",
						"type":        "string",
						"required":    true,
					},
				},
			},
		},
		Name: "Porter Preview Environment",
		Jobs: map[string]GithubActionYAMLJob{
			"porter-preview": {
				RunsOn: "ubuntu-latest",
				Steps:  gaSteps,
			},
		},
	}

	return yaml.Marshal(actionYAML)
}

func getPreviewDeleteActionYAML(opts *EnvOpts) ([]byte, error) {
	gaSteps := []GithubActionYAMLStep{
		getDeletePreviewEnvStep(
			opts.ServerURL,
			getPorterTokenSecretName(opts.ProjectID),
			opts.ProjectID,
			opts.ClusterID,
			opts.GitRepoName,
			"v0.2.0",
		),
	}

	actionYAML := GithubActionYAML{
		On: map[string]interface{}{
			"workflow_dispatch": map[string]interface{}{
				"inputs": map[string]interface{}{
					"environment_id": map[string]interface{}{
						"description": "Environment ID",
						"type":        "number",
						"required":    true,
					},
					"repo_owner": map[string]interface{}{
						"description": "Repository owner",
						"type":        "string",
						"required":    true,
					},
					"repo_name": map[string]interface{}{
						"description": "Repository name",
						"type":        "string",
						"required":    true,
					},
					"pr_number": map[string]interface{}{
						"description": "Pull request number",
						"type":        "number",
						"required":    true,
					},
				},
			},
		},
		Name: "Porter Preview Environment",
		Jobs: map[string]GithubActionYAMLJob{
			"porter-delete-preview": {
				RunsOn: "ubuntu-latest",
				Steps:  gaSteps,
			},
		},
	}

	return yaml.Marshal(actionYAML)
}

func createPorterPreviewBranch(opts *EnvOpts, defaultBranch string) error {
	_, resp, err := opts.Client.Repositories.GetBranch(
		context.Background(), opts.GitRepoOwner, opts.GitRepoName, "porter-preview", false,
	)

	if resp.StatusCode == http.StatusNotFound {
		branch, _, err := opts.Client.Repositories.GetBranch(
			context.Background(), opts.GitRepoOwner, opts.GitRepoName, defaultBranch, false,
		)

		if err != nil {
			return err
		}

		_, _, err = opts.Client.Git.CreateRef(
			context.Background(), opts.GitRepoOwner, opts.GitRepoName, &github.Reference{
				Ref: github.String("refs/heads/porter-preview"),
				Object: &github.GitObject{
					SHA: branch.Commit.SHA,
				},
			},
		)

		if err != nil {
			return err
		}

		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
