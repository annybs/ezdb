package ezdb

import (
	"errors"
	"testing"
)

// Basic struct for testing.
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var invalidStudents = map[string]*Student{
	"": {},
}

var nonexistentStudentKey = "nonexistent"

// Basic marshaler for testing.
var studentMarshaler = JSON(func() *Student {
	return &Student{}
})

// Sample data.
var students = map[string]*Student{
	"annie": {Name: "Annie", Age: 32},
	"ben":   {Name: "Ben", Age: 50},
	"clive": {Name: "Clive", Age: 21},
}

// Sample data (marshaled).
var studentsMarshaled = map[string][]byte{
	"annie": []byte("{\"name\":\"Annie\",\"age\":32}"),
	"ben":   []byte("{\"name\":\"Ben\",\"age\":50}"),
	"clive": []byte("{\"name\":\"Clive\",\"age\":21}"),
}

type CollectionTest struct {
	C Collection[*Student]
	T *testing.T

	F map[string]func() error
}

func (c *CollectionTest) open() error {
	if err := c.C.Open(); err != nil {
		c.T.Errorf("(open) failed to open collection: %v", err)
		return err
	}

	return nil
}

func (c *CollectionTest) put() error {
	// Test collection can store all students
	for key, value := range students {
		if err := c.C.Put(key, value); err != nil {
			c.T.Errorf("(put) failed to put student '%s': %v", key, err)
			return err
		}
		c.T.Logf("(put) put student '%s'", key)
	}

	// Test collection does not accept invalid keys
	for key, value := range invalidStudents {
		if err := c.C.Put(key, value); err == nil {
			c.T.Errorf("(put) should not have put invalid student '%s'", key)
			return err
		}
		c.T.Logf("(put) skipped invalid student '%s'", key)
	}

	return nil
}

func (c *CollectionTest) has() error {
	// Test collection has all students
	for key := range students {
		has, err := c.C.Has(key)
		if err != nil {
			c.T.Errorf("(has) failed to test whether collection has student '%s': %v", key, err)
			return err
		} else if !has {
			c.T.Errorf("(has) expected collection to have student '%s'", key)
			return err
		}
		c.T.Logf("(has) found student '%s'", key)
	}

	// Test collection does claim to have a student that doesn't exist
	has, err := c.C.Has(nonexistentStudentKey)
	if err != nil {
		c.T.Errorf("(has) failed to test whether collection has nonexistent student: %v", err)
	} else if has {
		c.T.Error("(has) expected collection not to have nonexistent student")
	} else {
		c.T.Logf("(has) collection does not have nonexistent student")
	}

	return nil
}

func (c *CollectionTest) get() error {
	// Test collection can retrieve all students
	for key, expected := range students {
		actual, err := c.C.Get(key)
		if err != nil {
			c.T.Errorf("(get) failed to get student '%s': %v", key, err)
			continue
		} else if actual.Name != expected.Name {
			c.T.Errorf("(get) student '%s' has wrong name (expected '%s', got '%s')", key, expected.Name, actual.Name)
		} else if actual.Age != expected.Age {
			c.T.Errorf("(get) student '%s' has wrong age (expected '%s', got '%s')", key, expected.Name, actual.Name)
		} else {
			c.T.Logf("(get) correctly got student '%s'", key)
		}
	}

	// Test collection does not retrieve a nonexistent student
	_, err := c.C.Get(nonexistentStudentKey)
	if err == nil {
		c.T.Error("(get) expected collection to return an error for nonexistent student")
	} else {
		c.T.Log("(get) collection did not get a nonexistent student")
	}

	return nil
}

func (c *CollectionTest) delete() error {
	if err := c.C.Delete("annie"); err != nil {
		c.T.Errorf("(delete) failed to delete student '%s': %v", "annie", err)
		return err
	}

	// Confirm student has been deleted
	has, err := c.C.Has("annie")
	if err != nil {
		c.T.Errorf("(delete) failed to test whether collection has deleted student 'annie': %v", err)
		return err
	} else if has {
		c.T.Error("(delete) expected collection not to have deleted student 'annie'")
		return err
	} else {
		c.T.Log("(delete) collection did not get the deleted student 'annie'")
	}

	// Reinsert deleted student
	if err := c.C.Put("annie", students["annie"]); err != nil {
		c.T.Errorf("(delete) failed to reinsert student 'annie': %v", err)
		return err
	} else {
		c.T.Log("(delete) reinserted student 'annie'")
	}

	return nil
}

func (c *CollectionTest) iterCount() error {
	iter := c.C.Iter()
	defer iter.Release()

	expected := len(students)
	actual := iter.Count()
	if expected != actual {
		c.T.Errorf("(iterCount) incorrect count of students (expected %d, got %d)", expected, actual)
		return errors.New("incorrect count")
	} else {
		c.T.Logf("(iterCount) correct count of students (expected %d, got %d)", expected, actual)
	}

	return nil
}

func (c *CollectionTest) close() error {
	if c.F["close"] != nil {
		if err := c.F["close"](); err != nil {
			c.T.Errorf("(close) failed to close collection: %v", err)
			return err
		}
	} else if err := c.C.Close(); err != nil {
		c.T.Errorf("(close) failed to close collection: %v", err)
		return err
	}

	c.T.Log("(close) closed database")

	return nil
}

func (c *CollectionTest) Run() {
	tests := []func() error{
		c.open,
		c.put,
		c.has,
		c.get,
		c.delete,
		c.iterCount,
		c.close,
	}

	for _, test := range tests {
		if err := test(); err != nil {
			return
		}
	}
}
