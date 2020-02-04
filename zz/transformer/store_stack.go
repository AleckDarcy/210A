package transformer

type StoreStack struct {
	KVs map[string]*StoreMeta

	Parent *StoreStack
}

type StoreType int64

const (
	StoreInteger StoreType = iota
	StoreFloat
	StoreFuncDefinition
)

type StoreMeta struct {
	typ StoreType
}

func (s *StoreStack) Get(name string) (*StoreMeta, bool) {
	m, ok := s.KVs[name]
	if ok {
		return m, true
	}

	if s.Parent != nil {
		return s.Parent.Get(name)
	}

	return nil, false
}
