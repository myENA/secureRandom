# secureRandom

## Summary

Simple semi-secure random string generation for golang.


## Usage

The package provides a single function:
```go
func New(n int) (string, error)
```

## Example

```go
package main

import (
	"fmt"
	srand "github.com/myENA/secureRandom"
)

func main() {
	s, err := srand.New(16)
	fmt.Println(s, err)
}
```

Please see the code and test case for more information.