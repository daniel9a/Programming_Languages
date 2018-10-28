package eval

import (
  "fmt"
  "testing"
)

//!+Depth
func TestDepth(t *testing.T) {
  tests := []struct {
    expr string
    want string
  }{
    {"10", "1"},
    {"km(10)", "2"},
    {"km(km(10))", "3"},
    {"m(km(10))", "3"},
    {"1 + 2 + 3 + -X", "4"},
    {"m(1) + km(m(km(2)))", "5"},
  }

  for _, test := range tests {
    expr, err := Parse(test.expr)
    if err != nil {
      t.Error(err) // parse error
      continue
    }

    fmt.Printf("\n%s\n", test.expr)

    // Run the method
    result := expr.Depth()

    // Display the result
    got := fmt.Sprintf("%d", result)
    fmt.Printf("\t%s => %s\n", expr, got)

    // Check the result
    if got != test.want {
      t.Errorf("(%s).Depth() = %q, want %q\n", test.expr, got, test.want)
    }
  }
}
//!-Depth