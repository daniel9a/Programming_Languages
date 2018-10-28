// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted for ECS 140A, Spring 2018 at UC Davis

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
  "fmt"
)

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) Eval(env Env) float64 {
  return env[v]
}

func (l Literal) Eval(_ Env) float64 {
  return float64(l)
}

//!-Eval1

//!+Eval2

func (u unary) Eval(env Env) float64 {
  switch u.op {
  case '+':
    return +u.x.Eval(env)
  case '-':
    return -u.x.Eval(env)
  }
  panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
  switch b.op {
  case '+':
    return b.x.Eval(env) + b.y.Eval(env)
  case '-':
    return b.x.Eval(env) - b.y.Eval(env)
  case '*':
    return b.x.Eval(env) * b.y.Eval(env)
  case '/':
    return b.x.Eval(env) / b.y.Eval(env)
  }
  panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (m measure) Eval(env Env) float64 {
  panic(fmt.Sprintf("this Eval cannot evaluate units: %s", m.unit))
}

//!-Eval2
