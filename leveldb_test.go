package ezdb

import "testing"

func TestLevelDB(t *testing.T) {
	path := ".leveldb/leveldb_test"
	c := LevelDB[*Student](path, studentMarshaler, nil)

	fixture := &CollectionTest{
		C: c,
		T: t,
		F: map[string]func() error{},
	}

	fixture.F["close"] = func() error {
		if err := c.Close(); err != nil {
			return err
		}
		if err := c.Destroy(); err != nil {
			return err
		}
		t.Logf("(leveldb) deleted data at %s", path)
		return nil
	}

	fixture.Run()
}
