package ezdb

type sortable[T any] struct {
	Key   string
	Value T
}

type keySort[T any] struct {
	a []*sortable[T]
	f SortFunc[string]
}

func (s *keySort[T]) Len() int {
	return len(s.a)
}

func (s *keySort[T]) Less(i, j int) bool {
	a := s.a[i]
	b := s.a[j]

	return s.f(a.Key, b.Key)
}

func (s *keySort[T]) Result() map[string]T {
	m := map[string]T{}
	for _, el := range s.a {
		m[el.Key] = el.Value
	}
	return m
}

func (s *keySort[T]) Swap(i, j int) {
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

func (s *valueSort[T]) Result() map[string]T {
	m := map[string]T{}
	for _, el := range s.a {
		m[el.Key] = el.Value
	}
	return m
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
