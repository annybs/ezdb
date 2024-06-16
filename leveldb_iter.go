package ezdb

import "github.com/syndtr/goleveldb/leveldb/iterator"

type LevelDBIterator[T any] struct {
	i iterator.Iterator
	m DocumentMarshaler[T, []byte]
}

func (i *LevelDBIterator[T]) Count() int {
	n := 0

	if i.First() {
		n++
	}

	for i.Next() {
		n++
	}

	return n
}

func (i *LevelDBIterator[T]) Filter(f FilterFunc[T]) Iterator[T] {
	m := map[string]T{}

	if i.First() {
		key, value, err := i.Get()
		if err == nil {
			m[key] = value
		}
	}

	for i.Next() {
		key, value, err := i.Get()
		if err != nil {
			continue
		}
		m[key] = value
	}

	return newMemoryIterator(m, nil, i)
}

func (i *LevelDBIterator[T]) First() bool {
	return i.i.First()
}

func (i *LevelDBIterator[T]) Get() (string, T, error) {
	value, err := i.Value()
	return i.Key(), value, err
}

func (i *LevelDBIterator[T]) GetAll() (map[string]T, error) {
	values := map[string]T{}

	if i.First() {
		key, value, err := i.Get()
		if err != nil {
			return values, err
		}
		values[key] = value
	}

	for i.Next() {
		key, value, err := i.Get()
		if err != nil {
			return values, err
		}
		values[key] = value
	}

	return values, nil
}

func (i *LevelDBIterator[T]) GetAllKeys() []string {
	keys := []string{}

	if i.First() {
		keys = append(keys, i.Key())
	}

	for i.Next() {
		keys = append(keys, i.Key())
	}

	return keys
}

func (i *LevelDBIterator[T]) Key() string {
	return string(i.i.Key())
}

func (i *LevelDBIterator[T]) Last() bool {
	return i.i.Last()
}

func (i *LevelDBIterator[T]) Next() bool {
	return i.i.Next()
}

func (i *LevelDBIterator[T]) Prev() bool {
	return i.i.Prev()
}

func (i *LevelDBIterator[T]) Release() {
	i.i.Release()
}

func (i *LevelDBIterator[T]) Sort(f SortFunc[T]) Iterator[T] {
	all, _ := i.GetAll()
	m := newMemoryIterator(all, nil, i)
	return m.Sort(f)
}

func (i *LevelDBIterator[T]) SortKeys(f SortFunc[string]) Iterator[T] {
	all, _ := i.GetAll()
	m := newMemoryIterator(all, nil, i)
	return m.SortKeys(f)
}

func (i *LevelDBIterator[T]) Value() (T, error) {
	value := i.m.Factory()
	err := i.m.Unmarshal(i.i.Value(), value)
	return value, err
}
