package ezdb

import (
	"bytes"
	"testing"
)

func TestJSONFactory(t *testing.T) {
	t.Logf("creating empty student")
	var value any = studentMarshaler.Factory()
	if _, ok := value.(*Student); !ok {
		t.Errorf("factory did not create correct value type (expected '*Student', got '%T')", value)
	}
}

func TestJSONMarshal(t *testing.T) {
	for key, value := range students {
		t.Logf("marshaling student '%s'", key)
		b, err := studentMarshaler.Marshal(value)
		if err != nil {
			t.Errorf("failed to marshal student '%s' (%q)", key, err)
		} else if !bytes.Equal(b, studentsMarshaled[key]) {
			t.Errorf("student '%s' incorrectly marshaled (expected '%s', got '%s')", key, studentsMarshaled[key], b)
		}
	}
}

func TestJSONUnmarshal(t *testing.T) {
	for key, b := range studentsMarshaled {
		t.Logf("unmarshaling student '%s'", key)
		value := studentMarshaler.Factory()
		if err := studentMarshaler.Unmarshal(b, value); err != nil {
			t.Errorf("failed to unmarshal student \"%s\" (%q)", key, err)
		} else {
			if value.Name != students[key].Name {
				t.Errorf("student '%s' name incorrectly unmarshaled (expected '%s', got '%s')", key, students[key].Name, value.Name)
			}
			if value.Age != students[key].Age {
				t.Errorf("student '%s' age incorrectly unmarshaled (expected '%d', got '%d')", key, students[key].Age, value.Age)
			}
		}
	}
}
