package main

import (
	"github.com/neatflowcv/vesta/api"
	"github.com/neatflowcv/vesta/internal/pkg/domain"
)

func toInstances(instances []*domain.Instance) []api.Instance {
	ret := make([]api.Instance, len(instances))
	for i, instance := range instances {
		ret[i] = toInstance(instance)
	}

	return ret
}

func toInstance(instance *domain.Instance) api.Instance {
	return api.Instance{
		Id:     instance.ID(),
		Name:   instance.Name(),
		Status: api.Status(instance.Status()),
	}
}
