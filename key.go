package ezdb

// ValidateKey validates whether a key is valid for putting data into a collection.
func ValidateKey(key string) error {
	if key == "" {
		return ErrInvalidKey
	}

	return nil
}
