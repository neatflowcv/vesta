package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/neatflowcv/vesta/api"
)

type Handler struct{}

func NewHandler() http.Handler { //nolint:ireturn
	handler := &Handler{}

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
	panic("unimplemented")
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
