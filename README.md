# Go Strings

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/0f9759dc5f2e48e7893f5062f66b1436)](https://app.codacy.com/gh/chr15k/go-strings?utm_source=github.com&utm_medium=referral&utm_content=chr15k/go-strings&utm_campaign=Badge_Grade)

Some useful string manipulation functions, extending the standard `strings` Go library.

Only Go standard libs are used. Please feel free to contribute.

## Usage

```go
package main

import (
	"fmt"

	"github.com/chr15k/go-strings/str"
)

func main() {
	mask := str.Mask("chris@example.com", "*", 3, 8)

	fmt.Println(mask) // chr********le.com
}
```

## Documentation

Visit [go.dev](https://pkg.go.dev/github.com/chr15k/go-strings#section-documentation)
