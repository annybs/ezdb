package ezdb

type MemoryCollection[T any] struct {
	c Collection[T]
	m map[string]T

	open bool
}

func (c *MemoryCollection[T]) Close() error {
	if c.c != nil {
		return c.c.Close()
	}

	c.m = map[string]T{}
	c.open = false

	return nil
}

func (c *MemoryCollection[T]) Delete(key string) error {
	if !c.open {
		return ErrClosed
	}

	if c.c != nil {
		if err := c.c.Delete(key); err != nil {
			return err
		}
	}

	delete(c.m, key)

	return nil
}

func (c *MemoryCollection[T]) Get(key string) (T, error) {
	if !c.open {
		return c.m[""], ErrClosed
	}

	if value, ok := c.m[key]; ok {
		return value, nil
	}

	return c.m[""], ErrNotFound
}

func (c *MemoryCollection[T]) Has(key string) (bool, error) {
	if !c.open {
		return false, ErrClosed
	}

	_, ok := c.m[key]

	return ok, nil
}

func (c *MemoryCollection[T]) Iter() Iterator[T] {
	m := newMemoryIterator[T](c.m, nil)
	if !c.open {
		m.Release()
	}
	return m
}

func (c *MemoryCollection[T]) Open() error {
	if c.c != nil {
		if err := c.c.Open(); err != nil {
			return err
		}
		all, err := c.c.Iter().GetAll()
		if err != nil {
			return err
		}
		c.m = all
	} else {
		c.m = map[string]T{}
	}

	c.open = true

	return nil
}

func (c *MemoryCollection[T]) Put(key string, value T) error {
	if !c.open {
		return ErrClosed
	}

	if err := ValidateKey(key); err != nil {
		return err
	}

	if c.c != nil {
		if err := c.c.Put(key, value); err != nil {
			return err
		}
	}

	c.m[key] = value

	return nil
}

// Memory creates an in-memory collection, which offers fast access without a document marshaler.
//
// If the collection c is non-nil, it will be used as a persistence backend.
//
// If T is a pointer type, the same pointer will be used whenever a document is read from this collection, so you should take care to treat documents as immutable.
func Memory[T any](c Collection[T]) *MemoryCollection[T] {
	return &MemoryCollection[T]{
		c: c,
		m: map[string]T{},
	}
}
