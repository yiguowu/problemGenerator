package main

import (
	"fmt"
	"problemGenerator/lib"
)

func main() {
	params := lib.ParseParameter()
	gen := new(lib.Generator)
	gen.Init(*params.Operator, *params.Range)
	for i := 0; i < *params.Count; i++ {
		fmt.Println(gen.Generate(*params.Operand))
		fmt.Println()
	}
}
