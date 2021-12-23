package types

type Environment struct {
	ID                uint   `json:"id"`
	ProjectID         uint   `json:"project_id"`
	ClusterID         uint   `json:"cluster_id"`
	GitInstallationID uint   `json:"git_installation_id"`
	GitRepoOwner      string `json:"git_repo_owner"`
	GitRepoName       string `json:"git_repo_name"`

	Name string `json:"name"`
}

type CreateEnvironmentRequest struct {
	Name         string `json:"name" form:"required"`
	GitRepoOwner string `json:"git_repo_owner" form:"required"`
	GitRepoName  string `json:"git_repo_name" form:"required"`
}

type GitHubMetadata struct {
	DeploymentID int64		`json:"gh_deployment_id"`
	PRName		 string		`json:"gh_pr_name"`
	RepoName	 string		`json:"gh_repo_name"`
	RepoOwner	 string		`json:"gh_repo_owner"`
	CommitSHA	 string		`json:"gh_commit_sha"`
}

type Deployment struct {
	*GitHubMetadata
	
	ID                 uint   `json:"id"`
	EnvironmentID      uint   `json:"environment_id"`
	Namespace          string `json:"namespace"`
	Status             string `json:"status"`
	Subdomain          string `json:"subdomain"`
	PullRequestID      uint   `json:"pull_request_id"`
}

type CreateGHDeploymentRequest struct {
	Branch   string `json:"branch" form:"required"`
	ActionID uint   `json:"action_id" form:"required"`
}

type CreateDeploymentRequest struct {
	*CreateGHDeploymentRequest
	*GitHubMetadata

	Namespace     string `json:"namespace" form:"required"`
	PullRequestID uint   `json:"pull_request_id" form:"required"`
}

type FinalizeDeploymentRequest struct {
	Namespace string `json:"namespace" form:"required"`
	Subdomain string `json:"subdomain"`
}

type UpdateDeploymentRequest struct {
	*CreateGHDeploymentRequest

	CommitSHA string `json:"commit_sha" form:"required"`	
	Namespace string `json:"namespace" form:"required"`
}

type DeleteDeploymentRequest struct {
	Namespace string `json:"namespace" form:"required"`
}

type GetDeploymentRequest struct {
	Namespace string `schema:"namespace" form:"required"`
}