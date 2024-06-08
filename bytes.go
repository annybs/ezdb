package ezdb

// BytesMarshaler is a DocumentMarshaler that simply passes along bytes.
type BytesMarshaler struct{}

func (m *BytesMarshaler) Factory() []byte {
	return []byte{}
}

func (m *BytesMarshaler) Marshal(src []byte) ([]byte, error) {
	return src, nil
}

func (m *BytesMarshaler) Unmarshal(src []byte, dest []byte) error {
	dest = src
	return nil
}

// Bytes creates a DocumentMarshaler that simply passes along bytes.
func Bytes() *BytesMarshaler {
	return &BytesMarshaler{}
}
