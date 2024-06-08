package ezdb

import (
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBCollection[T any] struct {
	path string

	db *leveldb.DB
	m  DocumentMarshaler[T, []byte]
}

func (c *LevelDBCollection[T]) Close() error {
	if c.db != nil {
		if err := c.db.Close(); err != nil {
			return err
		}

		c.db = nil
	}

	return nil
}

func (c *LevelDBCollection[T]) Delete(key string) error {
	return c.db.Delete([]byte(key), nil)
}

// Destroy the database completely, removing it from disk.
func (c *LevelDBCollection[T]) Destroy() error {
	if err := c.Close(); err != nil {
		return err
	}

	return os.RemoveAll(c.path)
}

func (c *LevelDBCollection[T]) Get(key string) (T, error) {
	dest := c.m.Factory()

	src, err := c.db.Get([]byte(key), nil)
	if err != nil {
		return dest, err
	}

	err = c.m.Unmarshal(src, dest)

	return dest, err
}

func (c *LevelDBCollection[T]) Has(key string) (bool, error) {
	return c.db.Has([]byte(key), nil)
}

func (c *LevelDBCollection[T]) Iter() Iterator[T] {
	i := &LevelDBIterator[T]{
		i: c.db.NewIterator(nil, nil),
		m: c.m,
	}

	return i
}

func (c *LevelDBCollection[T]) Open() error {
	if c.db == nil {
		db, err := leveldb.OpenFile(c.path, nil)
		if err != nil {
			return err
		}

		c.db = db
	}

	return nil
}

func (c *LevelDBCollection[T]) Put(key string, src T) error {
	if err := ValidateKey(key); err != nil {
		return err
	}

	dest, err := c.m.Marshal(src)
	if err != nil {
		return err
	}

	return c.db.Put([]byte(key), dest, nil)
}

// LevelDB creates a new collection using LevelDB storage.
func LevelDB[T any](path string, m DocumentMarshaler[T, []byte]) *LevelDBCollection[T] {
	c := &LevelDBCollection[T]{
		path: path,

		m: m,
	}
	return c
}
