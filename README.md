# safemap

Thread-safe maps using go generics

1. Compile and install go `master` from source https://golang.org/doc/install/source
2. `/my/go/install/go run -gcflags=-G=3 test.go`

test.go:

```go
package main

import (
    "fmt"

    "github.com/jacquayj/safemap"
)

test := safemap.NewMap(map[string]int{
    "test":    1234,
    "another": 456,
})

num, ok := test.Get("test")
fmt.Print(num, ok)

test.Set("another-key", 789)

num2, ok := test.Get("another-key")
fmt.Print(num2, ok)


ok = test.Delete("test")
fmt.Print(ok)

```
