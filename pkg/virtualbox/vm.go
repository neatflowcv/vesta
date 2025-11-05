package virtualbox

type VM struct {
	ID     string
	Name   string
	Status string
}

func NewVM(id, name, status string) *VM {
	return &VM{
		ID:     id,
		Name:   name,
		Status: status,
	}
}
