# EZ DB

This package provides simple interfaces for working with basic key-value storage in your Go application.

**EZ DB is not a database unto itself.** If you want more control or features, just use the appropriate database software and connector for your needs.

## Basic usage

The primary interface in EZ DB is `Collection[T]` which reflects a single key-value store. This is analogous to tables in RDBMS, collections in NoSQL databases etc.

Collections use a generic type `T` to specify the document type. You can use this to enforce a document schema. This example creates a collection which only accepts `Student` documents:

```go
package main

import "github.com/annybs/ezdb"

type Student struct {
	Name string
	Age int
}

var db = ezdb.Memory[Student](nil)

func main() {
	db.Open()
	db.Put("annie", Student{Name: "Annie", Age: "32"})
	db.Close()
}
```

In other cases, such as media stores, you may prefer not to specify a document type. This example allows arbitrary bytes to be written:

```go
package main

import "github.com/annybs/ezdb"

var db = ezdb.Memory[[]byte](nil)

func main() {
	db.Open()
	db.Put("data", []byte("arbitrary bytes"))
	db.Close()
}
```

## Marshaling data

Some database backends require marshaling and unmarshaling data. The `DocumentMarshaler[T1, T2]` interface allows you to use whatever marshaler suits your needs or the requirements of your chosen database.

The following marshalers are included in EZ DB:

- `Bytes` allows you to write `[]byte` directly to a database that requires `[]byte`
- `JSON[T]` marshals your data `T` to `[]byte` using [encoding/json](https://pkg.go.dev/encoding/json)

## Supported databases

The following databases are included in EZ DB:

- `LevelDB[T]` is [fast key-value storage](https://github.com/google/leveldb) on disk
- `Memory[T]` is essentially a wrapper for `map[string]T`. It can be provided another Collection to use as a persistence backend
