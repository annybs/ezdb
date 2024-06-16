package ezdb

// Collection is a key-value store for documents of any type.
type Collection[T any] interface {
	Open() error  // Open the collection.
	Close() error // Close the collection.

	Delete(key string) error              // Delete a document by key.
	Get(key string) (value T, err error)  // Get a document by key.
	Has(key string) (has bool, err error) // Check whether a document exists by key.
	Put(key string, value T) error        // Put a document into the collection.

	Iter() Iterator[T] // Get an iterator for this collection.
}

// DocumentMarshaler facilitates conversion between two types - a document and its storage representation, depending on the implementation of the Collection.
type DocumentMarshaler[T1 any, T2 any] interface {
	Factory() T1                     // Create a new, empty document.
	Marshal(src T1) (T2, error)      // Marshal a document to bytes.
	Unmarshal(src T2, dest T1) error // Unmarshal bytes into a document.
}

// FilterFunc processes a document as part of a filter operation.
// This function returns true if the document passes all checks defined in the filter.
type FilterFunc[T any] func(key string, value T) bool

// Iterator provides functionality to explore a collection.
//
// Be mindful that the order of documents is not assured by any Collection implementation.
// Use the Sort or SortKeys function before iterating over documents to ensure deterministic sort.
type Iterator[T any] interface {
	First() bool // Move the iterator to the first document. Returns false if there is no first document.
	Last() bool  // Move the iterator to the last document. Returns false if there is no last document.
	Next() bool  // Move the iterator to the next document. Returns false if there is no next document.
	Prev() bool  // Move the iterator to the previous document. Returns false if there is no previous document.

	Release() // Release the iterator and any associated resources, including those of previous iterators.

	Count() int // Count the number of documents in the iterator.

	Get() (key string, value T, err error) // Get the key and value of the current document.
	Key() string                           // Get the key of the current document.
	Value() (T, error)                     // Get the value of the current document.

	GetAll() (map[string]T, error) // Get all documents as a key-value map.
	GetAllKeys() []string          // Get all document keys.

	Filter(f FilterFunc[T]) Iterator[T]      // Create a new iterator with a subset of documents. The previous iterator will not be affected.
	Sort(f SortFunc[T]) Iterator[T]          // Create a new iterator with sorted documents. The previous iterator will not be affected.
	SortKeys(f SortFunc[string]) Iterator[T] // Create a new iterator with documents sorted by key. The previous iterator will not be affected.
}

// SortFunc compares two documents as part of a sort operation.
// This function returns false if a is less than b.
type SortFunc[T any] func(a T, b T) bool
