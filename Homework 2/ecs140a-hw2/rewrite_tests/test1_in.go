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

	result = eval.ParseAndEval("1 + 2", eval.Env{})
	fmt.Printf("%d\n", result)

	result = ParseAndEval("1 + 2", eval.Env{})
	fmt.Printf("%d\n", result)

	x := "1 + 2"
	result = eval.ParseAndEval(x, eval.Env{})
	fmt.Printf("%d\n", result)

	result = eval.ParseAndEval("1 + / / 2", eval.Env{})
	fmt.Printf("%d\n", result)
}
