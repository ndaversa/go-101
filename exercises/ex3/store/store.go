package store

type Keyer interface {
	Key() uint64
}

type InMemory struct {
	m map[uint64]Keyer
}

func New() *InMemory {
	return &InMemory{
		m: make(map[uint64]Keyer),
	}
}

func (s *InMemory) Get(k uint64) Keyer {
	return s.m[k]
}

func (s *InMemory) Put(v Keyer) {
	s.m[v.Key()] = v
}
