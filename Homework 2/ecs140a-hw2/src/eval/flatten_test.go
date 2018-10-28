package eval

import (
  "fmt"
  "testing"
  "math"
)

//!+FlattenUnits
func TestFlattenUnits(t *testing.T) {
  float_eps := 0.0001

  tests := []struct {
    expr string
    want_mag float64
    want_unit string
  } {
    {"10", 10, "scalar"},
    {"km(10)", 10, "km"},
    {"km(km(10))", 10, "km"},
    {"m(km(10))", 10000, "m"},
    {"F(C(0))", 32, "F"},
    {"C(K(0))", -273.15, "C"},
    {"s(4) + min(5)", 304, "s"},
    {"F(1) + K(25)", -413.67, "F"},
    {"K(25) + F(1)", 280.9278, "K"},
    {"mi(1) * 5", 5, "mi"},
    {"14 * min(2)", 28, "min"},
    {"10 * m(5) / 2", 25, "m"},
    {"mi_p_s(km(60) / s(1))", 37.2823, "mi_p_s"},
    {"m_p_gal(km(20) / ltr(1))", 75708.2, "m_p_gal"},

    // all combinations
    {"m_p_s(m(1) / s(1))",         1.0/1.0,   "m_p_s"},
    {"m_p_s(m(1) / ms(1))",        1.0/0.001, "m_p_s"},
    {"m_p_s(m(1) / min(1))",       1.0/60.0,  "m_p_s"},
    {"m_p_s(km(1) / s(1))",     1000.0/1.0,   "m_p_s"},
    {"m_p_s(km(1) / ms(1))",    1000.0/0.001, "m_p_s"},
    {"m_p_s(km(1) / min(1))",   1000.0/60.0,  "m_p_s"},
    {"m_p_s(mi(1) / s(1))",   1609.344/1.0,   "m_p_s"},
    {"m_p_s(mi(1) / ms(1))",  1609.344/0.001, "m_p_s"},
    {"m_p_s(mi(1) / min(1))", 1609.344/60.0,  "m_p_s"},

    {"m_p_ltr(m(1) / ltr(1))",      1.0/1.0,     "m_p_ltr"},
    {"m_p_ltr(m(1) / gal(1))",      1.0/3.78541, "m_p_ltr"},
    {"m_p_ltr(km(1) / ltr(1))",   1000.0/1.0,     "m_p_ltr"},
    {"m_p_ltr(km(1) / gal(1))",   1000.0/3.78541, "m_p_ltr"},
    {"m_p_ltr(mi(1) / ltr(1))", 1609.344/1.0,     "m_p_ltr"},
    {"m_p_ltr(mi(1) / gal(1))", 1609.344/3.78541, "m_p_ltr"},
    
    {"mi(1) * -5", -5, "mi"},
    {"m_p_s(m(-1) / s(1))",         -1.0/1.0,   "m_p_s"},
    {"km(X)", 42, "km"},
  }

  for _, test := range tests {
    expr, err := Parse(test.expr)
    if err != nil {
      t.Error(err) // parse error
      continue
    }

    fmt.Printf("\n%s\n", test.expr)

    // Run the method
    expr, unit := expr.FlattenUnits()

    // Display the result
    got := expr.Eval(Env{"X": 42})
    fmt.Printf("\t%s => %g [%s]\n", Format(expr), got, unit)

    // Check the result
    if math.Abs(test.want_mag - got) > float_eps || unit != test.want_unit {
      t.Errorf("(%s).FlattenUnits() = %g [%q], want %g [%q]\n",
        test.expr, got, unit, test.want_mag, test.want_unit)
    }
  }
}

func TestFlattenUnits_Failure(t *testing.T) {
  tests := []struct {
    expr Expr
  } {
    {binary{'+', measure{"m", Literal(10.0)}, measure{"F", Literal(1.0)}}},
    {binary{'/', measure{"m", Literal(10.0)}, measure{"m", Literal(1.0)}}},
    {binary{'%', measure{"m", Literal(10.0)}, measure{"m", Literal(1.0)}}},
    {binary{'*', measure{"m", Literal(10.0)}, measure{"m", Literal(1.0)}}},
    {binary{'/', Literal(10.0), measure{"m", Literal(1.0)}}},
  }

  for _, test := range tests {
    func() {
      defer func() {
        if recover() == nil {
          t.Errorf("(%s).FlattenUnits() did not panic, but should\n",
            test.expr)
        }
      }()

      test.expr.FlattenUnits()
    }()
  }
}
//!-FlattenUnits