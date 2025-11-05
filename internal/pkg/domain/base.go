package domain

type Base struct {
	id      string
	name    string
	cpu     uint64
	ram     uint64
	storage uint64
}

func NewBase(id, name string, cpu, ram, storage uint64) *Base {
	return &Base{
		id:      id,
		name:    name,
		cpu:     cpu,
		ram:     ram,
		storage: storage,
	}
}

func (b *Base) ID() string {
	return b.id
}

func (b *Base) Name() string {
	return b.name
}

func (b *Base) CPU() uint64 {
	return b.cpu
}

func (b *Base) RAM() uint64 {
	return b.ram
}

func (b *Base) Storage() uint64 {
	return b.storage
}
