package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/neatflowcv/vesta/api"
	"github.com/neatflowcv/vesta/internal/app/flow"
)

type Handler struct {
	service *flow.Service
}

func NewHandler(service *flow.Service) http.Handler { //nolint:ireturn
	handler := &Handler{
		service: service,
	}

	return api.HandlerFromMux(api.NewStrictHandler(handler, nil), chi.NewMux())
}

func (h *Handler) DeleteVestaV1BasesId(
	ctx context.Context,
	request api.DeleteVestaV1BasesIdRequestObject,
) (api.DeleteVestaV1BasesIdResponseObject, error) {
	panic("unimplemented")
}

func (h *Handler) DeleteVestaV1InstancesId(
	ctx context.Context,
	request api.DeleteVestaV1InstancesIdRequestObject,
) (api.DeleteVestaV1InstancesIdResponseObject, error) {
	panic("unimplemented")
}

func (h *Handler) ListBases(
	ctx context.Context,
	request api.ListBasesRequestObject,
) (api.ListBasesResponseObject, error) {
	panic("unimplemented")
}

func (h *Handler) ListInstances(
	ctx context.Context,
	request api.ListInstancesRequestObject,
) (api.ListInstancesResponseObject, error) {
	instances, err := h.service.ListInstances(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	if len(instances) == 0 {
		return api.ListInstances204Response{}, nil
	}

	return api.ListInstances200JSONResponse(toInstances(instances)), nil
}

func (h *Handler) PostVestaV1BasesIdClone(
	ctx context.Context,
	request api.PostVestaV1BasesIdCloneRequestObject,
) (api.PostVestaV1BasesIdCloneResponseObject, error) {
	panic("unimplemented")
}

func (h *Handler) PostVestaV1InstancesIdPromote(
	ctx context.Context,
	request api.PostVestaV1InstancesIdPromoteRequestObject,
) (api.PostVestaV1InstancesIdPromoteResponseObject, error) {
	panic("unimplemented")
}

func (h *Handler) PostVestaV1InstancesIdStart(
	ctx context.Context,
	request api.PostVestaV1InstancesIdStartRequestObject,
) (api.PostVestaV1InstancesIdStartResponseObject, error) {
	panic("unimplemented")
}

func (h *Handler) PostVestaV1InstancesIdStop(
	ctx context.Context,
	request api.PostVestaV1InstancesIdStopRequestObject,
) (api.PostVestaV1InstancesIdStopResponseObject, error) {
	panic("unimplemented")
}
