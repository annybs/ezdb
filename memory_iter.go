package ezdb

import "sort"

type MemoryIterator[T any] struct {
	k []string
	m map[string]T

	pos      int
	released bool

	prev Iterator[T]
}

func (i *MemoryIterator[T]) Count() int {
	return len(i.k)
}

func (i *MemoryIterator[T]) Filter(f FilterFunc[T]) Iterator[T] {
	if i.released {
		return i
	}

	k := []string{}
	m := map[string]T{}
	i.reset()
	for i.Next() {
		key, value, _ := i.Get()
		if f(key, value) {
			k = append(k, key)
			m[key] = value
		}
	}
	return newMemoryIterator(m, k, i)
}

func (i *MemoryIterator[T]) First() bool {
	if i.released {
		return false
	}
	i.pos = 0
	return true
}

func (i *MemoryIterator[T]) Get() (string, T, error) {
	if i.released {
		return "", i.m[""], ErrReleased
	}

	key := i.Key()
	return key, i.m[key], nil
}

func (i *MemoryIterator[T]) GetAll() (map[string]T, error) {
	m := map[string]T{}
	if i.released {
		return m, ErrReleased
	}

	i.reset()
	for i.Next() {
		key, value, _ := i.Get()
		m[key] = value
	}
	return m, nil
}

func (i *MemoryIterator[T]) GetAllKeys() []string {
	keys := []string{}
	if i.released {
		return keys
	}

	i.reset()
	for i.Next() {
		keys = append(keys, i.Key())
	}
	return keys
}

func (i *MemoryIterator[T]) Key() string {
	if i.pos < 0 || i.pos > i.Count() || i.released {
		return ""
	}
	return i.k[i.pos]
}

func (i *MemoryIterator[T]) Last() bool {
	if i.released {
		return false
	}
	i.pos = len(i.k) - 1
	return true
}

func (i *MemoryIterator[T]) Next() bool {
	if i.released {
		return false
	}

	hasNext := i.pos+1 <= i.Count()
	if hasNext {
		i.pos++
	}
	return hasNext
}

func (i *MemoryIterator[T]) Prev() bool {
	if i.released {
		return false
	}

	end := i.pos > 0
	if !end {
		i.pos--
	}
	return end
}

func (i *MemoryIterator[T]) Release() {
	i.k = []string{}
	i.m = map[string]T{}
	i.released = true

	if i.prev != nil {
		i.prev.Release()
	}
}

func (i *MemoryIterator[T]) Sort(f SortFunc[T]) Iterator[T] {
	if i.released {
		return i
	}

	s := &valueSort[T]{
		a: makeSortable(i.m),
		f: f,
	}
	sort.Stable(s)
	k := s.Result()

	return newMemoryIterator(i.m, k, i)
}

func (i *MemoryIterator[T]) SortKeys(f SortFunc[string]) Iterator[T] {
	if i.released {
		return i
	}

	s := &keySort{
		a: i.k,
		f: f,
	}
	sort.Stable(s)
	k := s.Result()

	return newMemoryIterator(i.m, k, i)
}

func (i *MemoryIterator[T]) Value() (T, error) {
	key := i.Key()
	return i.m[key], nil
}

func (i *MemoryIterator[T]) reset() {
	i.pos = -1
}

func newMemoryIterator[T any](m map[string]T, k []string, prev Iterator[T]) *MemoryIterator[T] {
	i := &MemoryIterator[T]{
		k: []string{},
		m: m,

		pos:  -1,
		prev: prev,
	}

	if len(k) > 0 {
		i.k = k
	} else {
		for k := range i.m {
			i.k = append(i.k, k)
		}
	}

	return i
}
