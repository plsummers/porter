package policy

import (
	"encoding/json"
	"net/http"

	"github.com/porter-dev/porter/api/server/handlers"
	"github.com/porter-dev/porter/api/server/shared"
	"github.com/porter-dev/porter/api/server/shared/apierrors"
	"github.com/porter-dev/porter/api/server/shared/config"
	"github.com/porter-dev/porter/api/types"
	"github.com/porter-dev/porter/internal/models"
	"github.com/porter-dev/porter/internal/repository"
)

type PolicyCreateHandler struct {
	handlers.PorterHandlerReadWriter
}

func NewPolicyCreateHandler(
	config *config.Config,
	decoderValidator shared.RequestDecoderValidator,
	writer shared.ResultWriter,
) *PolicyCreateHandler {
	return &PolicyCreateHandler{
		PorterHandlerReadWriter: handlers.NewDefaultPorterHandler(config, decoderValidator, writer),
	}
}

func (p *PolicyCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(types.UserScope).(*models.User)
	proj, _ := r.Context().Value(types.ProjectScope).(*models.Project)

	req := &types.CreatePolicy{}

	if ok := p.DecodeAndValidate(w, r, req); !ok {
		return
	}

	uid, err := repository.GenerateRandomBytes(16)

	if err != nil {
		p.HandleAPIError(w, r, apierrors.NewErrInternal(err))
		return
	}

	policyBytes, err := json.Marshal(req.Policy)

	if err != nil {
		p.HandleAPIError(w, r, apierrors.NewErrInternal(err))
		return
	}

	policy := &models.Policy{
		ProjectID:       proj.ID,
		UniqueID:        uid,
		CreatedByUserID: user.ID,
		Name:            req.Name,
		PolicyBytes:     policyBytes,
	}

	policy, err = p.Repo().Policy().CreatePolicy(policy)

	if err != nil {
		p.HandleAPIError(w, r, apierrors.NewErrInternal(err))
		return
	}

	res, err := policy.ToAPIPolicyType()

	if err != nil {
		p.HandleAPIError(w, r, apierrors.NewErrInternal(err))
		return
	}

	p.WriteResult(w, r, res)
}