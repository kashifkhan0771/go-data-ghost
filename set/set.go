package set

type Set struct {
	data map[interface{}]bool
}

func NewSet() *Set {
	return &Set{data: make(map[interface{}]bool)}
}

func (s *Set) Add(item interface{}) {
	s.data[item] = true
}

func (s *Set) Remove(item interface{}) {
	delete(s.data, item)
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.data[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) Clear() {
	s.data = make(map[interface{}]bool)
}
