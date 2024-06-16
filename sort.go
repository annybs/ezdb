package ezdb

type sortable[T any] struct {
	Key   string
	Value T
}

type keySort struct {
	a []string
	f SortFunc[string]
}

func (s *keySort) Len() int {
	return len(s.a)
}

func (s *keySort) Less(i, j int) bool {
	return s.f(s.a[i], s.a[j])
}

func (s *keySort) Result() []string {
	return s.a
}

func (s *keySort) Swap(i, j int) {
	a := s.a[i]
	b := s.a[j]
	s.a[i] = b
	s.a[j] = a
}

type valueSort[T any] struct {
	a []*sortable[T]
	f SortFunc[T]
}

func (s *valueSort[T]) Len() int {
	return len(s.a)
}

func (s *valueSort[T]) Less(i, j int) bool {
	a := s.a[i]
	b := s.a[j]

	return s.f(a.Value, b.Value)
}

func (s *valueSort[T]) Result() []string {
	k := []string{}
	for _, el := range s.a {
		k = append(k, el.Key)
	}
	return k
}

func (s *valueSort[T]) Swap(i, j int) {
	a := s.a[i]
	b := s.a[j]
	s.a[i] = b
	s.a[j] = a
}

func makeSortable[T any](m map[string]T) []*sortable[T] {
	a := []*sortable[T]{}
	for key, value := range m {
		a = append(a, &sortable[T]{Key: key, Value: value})
	}
	return a
}
