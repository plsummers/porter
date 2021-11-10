package kube_events

import (
	"net/http"
	"strings"
	"time"

	"github.com/porter-dev/porter/api/server/authz"
	"github.com/porter-dev/porter/api/server/handlers"
	"github.com/porter-dev/porter/api/server/shared"
	"github.com/porter-dev/porter/api/server/shared/apierrors"
	"github.com/porter-dev/porter/api/server/shared/config"
	"github.com/porter-dev/porter/api/types"
	"github.com/porter-dev/porter/internal/integrations/slack"
	"github.com/porter-dev/porter/internal/models"
)

type CreateKubeEventHandler struct {
	handlers.PorterHandlerReadWriter
	authz.KubernetesAgentGetter
}

func NewCreateKubeEventHandler(
	config *config.Config,
	decoderValidator shared.RequestDecoderValidator,
	writer shared.ResultWriter,
) *CreateKubeEventHandler {
	return &CreateKubeEventHandler{
		PorterHandlerReadWriter: handlers.NewDefaultPorterHandler(config, decoderValidator, writer),
		KubernetesAgentGetter:   authz.NewOutOfClusterAgentGetter(config),
	}
}

func (c *CreateKubeEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	proj, _ := r.Context().Value(types.ProjectScope).(*models.Project)
	cluster, _ := r.Context().Value(types.ClusterScope).(*models.Cluster)

	request := &types.CreateKubeEventRequest{}

	if ok := c.DecodeAndValidate(w, r, request); !ok {
		return
	}

	// Look for an event matching by the name, namespace, and was last updated within the
	// grouping threshold time. If so, we append a subevent to the existing event.
	kubeEvent, err := c.Repo().KubeEvent().ReadEventByGroup(proj.ID, cluster.ID, &types.GroupOptions{
		Name:          request.Name,
		Namespace:     request.Namespace,
		ResourceType:  request.ResourceType,
		ThresholdTime: time.Now().Add(-15 * time.Minute),
	})

	foundMatchedEvent := kubeEvent != nil

	if !foundMatchedEvent {
		kubeEvent, err = c.Repo().KubeEvent().CreateEvent(&models.KubeEvent{
			ProjectID:    proj.ID,
			ClusterID:    cluster.ID,
			ResourceType: request.ResourceType,
			Name:         request.Name,
			OwnerType:    request.OwnerType,
			OwnerName:    request.OwnerName,
			Namespace:    request.Namespace,
		})

		if err != nil {
			c.HandleAPIError(w, r, apierrors.NewErrInternal(err))
			return
		}
	}

	// append the subevent to the event
	err = c.Repo().KubeEvent().AppendSubEvent(kubeEvent, &models.KubeSubEvent{
		EventType: request.EventType,
		Message:   request.Message,
		Reason:    request.Reason,
		Timestamp: request.Timestamp,
	})

	if err != nil {
		c.HandleAPIError(w, r, apierrors.NewErrInternal(err))
		return
	}

	w.WriteHeader(http.StatusCreated)

	if strings.ToLower(string(request.EventType)) == "critical" && strings.ToLower(request.ResourceType) == "pod" {
		err := notifyPodCrashing(c.Config(), proj, cluster, request)

		if err != nil {
			c.HandleAPIErrorNoWrite(w, r, apierrors.NewErrInternal(err))
		}
	}
}

func notifyPodCrashing(
	config *config.Config,
	project *models.Project,
	cluster *models.Cluster,
	event *types.CreateKubeEventRequest,
) error {
	slackInts, _ := config.Repo.SlackIntegration().ListSlackIntegrationsByProjectID(project.ID)

	notifier := slack.NewSlackNotifier(&types.NotificationConfig{
		Enabled: true,
		Success: true,
		Failure: true,
	}, slackInts...)

	notifyOpts := &slack.NotifyOpts{
		ProjectID:   cluster.ProjectID,
		ClusterID:   cluster.ID,
		ClusterName: cluster.Name,
		Name:        event.OwnerName,
		Namespace:   event.Namespace,
		URL:         config.ServerConf.ServerURL,
	}

	notifyOpts.Status = slack.StatusPodCrashed

	return notifier.Notify(notifyOpts)
}
