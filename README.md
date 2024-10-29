# Go Strings

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/1a26c1b2bcfc48c5920d33838015a69e)](https://app.codacy.com/gh/chr15k/go-strings/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)&nbsp;[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/1a26c1b2bcfc48c5920d33838015a69e)](https://app.codacy.com/gh/chr15k/go-strings/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)

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
