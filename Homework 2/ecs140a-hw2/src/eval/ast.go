// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted for ECS 140A, Spring 2018 at UC Davis

package eval

// An Expr is an arithmetic expression, potentially with units.
type Expr interface {
  // Eval returns the value of this Expr under the environment env.
  Eval(env Env) float64

  // Check reports errors in this Expr and adds its Vars to the set.
  Check(vars map[Var]bool) error

  // Simplify simplifies an expression given a partial (or empty) environment
  Simplify(env Env) Expr

  // FlattenUnits replaces measure expressions with the appropriate arithmetic
  // required to perform unit conversions, and also returns the unit of the
  // expression.
  FlattenUnits() (Expr, string)

  // Depth determines the maximum depth of the given expression tree.
  Depth() uint
}


//!+ast

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type Literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
  op rune // one of '+', '-'
  x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
  op   rune // one of '+', '-', '*', '/'
  x, y Expr
}

// A measure represents a unit conversion expression, e.g., km(10) or s(ms(4))
type measure struct {
  unit string
  x    Expr
}

//!-ast
