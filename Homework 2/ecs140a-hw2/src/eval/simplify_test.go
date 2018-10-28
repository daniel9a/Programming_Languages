package eval

import (
  "fmt"
  "testing"
)

//!+Simplify
func TestSimplify(t *testing.T) {
  tests := []struct {
    expr string
    env  Env
    want string
  } {
    {"10 / X", Env{"X": 2}, "5"},
    {"(X + X) - Y", Env{"X": 2}, "(4 - Y)"},
    {"(X + X) - Y", Env{"Y": 8}, "((X + X) - 8)"},
    {"5 + 2", Env{}, "7"},
    {"10 - 1 + X - Y", Env{}, "((9 + X) - Y)"},
    {"X + 3 + 5", Env{}, "((X + 3) + 5)"},
    {"1 + 2 + -X", Env{}, "(3 + (-X))"},
    {"10 * X", Env{"X": 2}, "20"},
    {"10 * X", Env{}, "(10 * X)"},
    {"10 / X", Env{}, "(10 / X)"},
  }

  for _, test := range tests {
    expr, err := Parse(test.expr)
    if err != nil {
      t.Error(err) // parse error
      continue
    }

    fmt.Printf("\n%s\n", test.expr)

    // Run the method
    result := expr.Simplify(test.env)

    // Display the result
    got := Format(result)
    fmt.Printf("\t%s, %v => %s\n", Format(expr), test.env, got)

    // Check the result
    if got != test.want {
      t.Errorf("(%s).Simplify() in %v = %q, want %q\n",
        test.expr, test.env, got, test.want)
    }
  }
}
func TestSimplify_Failure(t *testing.T) {
  tests := []struct {
    expr Expr
  } {
    {measure{"m", Literal(10.0)}},
    {binary{'%', Literal(10.0), Literal(1.0)}},
  }

  for _, test := range tests {
    func() {
      defer func() {
        if recover() == nil {
          t.Errorf("(%s).Simplify(Env{}) did not panic, but should\n",
            test.expr)
        }
      }()

      test.expr.Simplify(Env{})
    }()
  }
}
//!-Simplify