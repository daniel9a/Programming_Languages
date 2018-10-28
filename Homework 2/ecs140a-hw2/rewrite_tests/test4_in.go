
package main

import (
	"eval"
	"fmt"
)

func ParseAndEval(x string, y eval.Env) float64 {
	return 42.0
}

func main() {
	var result float64

	// Note the syntax error -- parsing this file should cause a panic().
	result = eval.ParseAndEval("1 + 2", eval.Env{)
	fmt.Printf("%d\n", result)
}
