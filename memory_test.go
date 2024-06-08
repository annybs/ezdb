package ezdb

import "testing"

func TestMemory(t *testing.T) {
	c := Memory[*Student](nil)

	fixture := &CollectionTest{
		C: c,
		T: t,
	}

	fixture.Run()
}
