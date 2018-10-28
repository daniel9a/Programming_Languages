package main

import (
	"eval"
	"fmt"
)

func main() {
	var result float64

	result = eval.ParseAndEval("1 + 2")
	fmt.Printf("%d\n", result)
}
