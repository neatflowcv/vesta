package domain

type InstanceStatus string

const (
	InstanceStatusBooting  InstanceStatus = "Booting"
	InstanceStatusRunning  InstanceStatus = "Running"
	InstanceStatusStopped  InstanceStatus = "Stopped"
	InstanceStatusStopping InstanceStatus = "Stopping"
	InstanceStatusUnknown  InstanceStatus = "Unknown"
)

type Instance struct {
	id     string
	name   string
	status InstanceStatus
}

func NewInstance(id, name string, status InstanceStatus) *Instance {
	return &Instance{
		id:     id,
		name:   name,
		status: status,
	}
}

func (i *Instance) ID() string {
	return i.id
}

func (i *Instance) Name() string {
	return i.name
}

func (i *Instance) Status() InstanceStatus {
	return i.status
}
