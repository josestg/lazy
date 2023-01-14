# Lazy

A lazy value loader, it can be used to load expensive computation only when it is needed 
and cache the result for future use.

## Installation

```bash
go get github.com/josestg/lazy
```

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/josestg/lazy"
)

func main() {
	// Create a lazy value loader.
	l := lazy.New(func() (string, error) {
		time.Sleep(1 * time.Second)
		return "Hello, World!", nil
	})

	// Get the value for the first time will be slow.
	fmt.Println(l.Value())  // Hello, World!
	fmt.Println(l.Loaded()) // true

	// Get the value for the next will return the cached value.
	fmt.Println(l.Value()) // Hello, World!
	fmt.Println(l.Value()) // Hello, World!
}
```