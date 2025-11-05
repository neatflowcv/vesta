package virtualbox

type VM struct {
	ID      string
	Name    string
	CPU     uint64
	RAM     uint64
	Storage uint64
	Status  string
}

func NewVM(id, name, status string, cpu, ram, storage uint64) *VM {
	return &VM{
		ID:      id,
		Name:    name,
		CPU:     cpu,
		RAM:     ram,
		Storage: storage,
		Status:  status,
	}
}
