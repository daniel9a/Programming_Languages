// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted for ECS 140A, Spring 2018 at UC Davis

package eval

import (
  "fmt"
  "strings"
)

//!+Check

func (v Var) Check(vars map[Var]bool) error {
  vars[v] = true
  return nil
}

func (Literal) Check(vars map[Var]bool) error {
  return nil
}

func (u unary) Check(vars map[Var]bool) error {
  if !strings.ContainsRune("+-", u.op) {
    return fmt.Errorf("unexpected unary op %q", u.op)
  }
  return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
  if !strings.ContainsRune("+-*/", b.op) {
    return fmt.Errorf("unexpected binary op %q", b.op)
  }
  if err := b.x.Check(vars); err != nil {
    return err
  }
  return b.y.Check(vars)
}

func (m measure) Check(vars map[Var]bool) error {
  return m.x.Check(vars)
}

//!-Check
