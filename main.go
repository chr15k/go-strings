package main

import (
	"fmt"

	"github.com/chr15k/go-strings/str"
)

func main() {

	mask := str.Mask("chris@example.com", "*", 3, 8)

	fmt.Println(mask) // chr********le.com
}
