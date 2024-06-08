package ezdb

import "encoding/json"

// JSONMarshaler is a DocumentMarshaler that converts documents to JSON data.
type JSONMarshaler[T any] struct {
	factory func() T
}

func (m *JSONMarshaler[T]) Factory() T {
	return m.factory()
}

func (m *JSONMarshaler[T]) Marshal(src T) ([]byte, error) {
	return json.Marshal(src)
}

func (m *JSONMarshaler[T]) Unmarshal(src []byte, dest T) error {
	return json.Unmarshal(src, dest)
}

// JSON creates a DocumentMarshaler that converts documents to JSON data.
func JSON[T any](factory func() T) *JSONMarshaler[T] {
	return &JSONMarshaler[T]{factory: factory}
}
