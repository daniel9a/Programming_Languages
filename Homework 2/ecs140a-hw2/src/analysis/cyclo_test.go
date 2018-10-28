package analysis

import (
  // "fmt"
  "testing"
)

//!+CyclomaticComplexity
func TestCyclomaticComplexity(t *testing.T) {
  var test_code = `
    package main

    import (
      "fmt"
      "eval"
    )

    func f() {
      return 42
    }

    func g(x int) {
      if x < 0 {
        return -1;
      } else if x > 0 {
        return 1;
      } else {
        return 0;
      }
    }

    func h() {
      switch 5 {
      case 0:
        // pass
      case 5:
        fmt.Println("It's five!")
      default:
        fmt.Println("It isn't five...")
      }
    }
  `

  tests := []struct {
    name string
    cyclo uint
  }{
    {"f", 1},
    {"g", 3},
    {"h", 3},
  }

  cyclos := CyclomaticComplexity(test_code)

  for _, test := range tests {
    if cyclos[test.name] != test.cyclo {
      t.Errorf("CyclomaticComplexity()[%v] = %d, want %d\n",
        test.name, cyclos[test.name], test.cyclo)
    }
  }
}
//!-CyclomaticComplexity
