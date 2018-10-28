// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted for ECS 140A, Spring 2018 at UC Davis

package eval

import (
  "bytes"
  "fmt"
)

// Format formats an expression as a string.
// It does not attempt to remove unnecessary parens.
func Format(e Expr) string {
  var buf bytes.Buffer
  write(&buf, e)
  return buf.String()
}

func write(buf *bytes.Buffer, e Expr) {
  switch e := e.(type) {
  case Literal:
    fmt.Fprintf(buf, "%g", e)

  case Var:
    fmt.Fprintf(buf, "%s", e)

  case unary:
    fmt.Fprintf(buf, "(%c", e.op)
    write(buf, e.x)
    buf.WriteByte(')')

  case binary:
    buf.WriteByte('(')
    write(buf, e.x)
    fmt.Fprintf(buf, " %c ", e.op)
    write(buf, e.y)
    buf.WriteByte(')')

  case measure:
    fmt.Fprintf(buf, "%s(", e.unit)
    write(buf, e.x)
    buf.WriteByte(')')

  default:
    panic(fmt.Sprintf("unknown Expr: %T", e))
  }
}
