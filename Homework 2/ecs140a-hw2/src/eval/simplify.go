package eval

func (v Var) Simplify(env Env) Expr {
  x, ok := env[v]
  if ok {
      return Literal(x)
  } else {
    return v
  }
}

func (f Literal) Simplify(env Env) Expr {
  return f
}

func (u unary) Simplify(env Env) Expr {
  x := u.x.Simplify(env);
  return unary{u.op, x};
}

func (b binary) Simplify(env Env) Expr {
  x := b.x.Simplify(env)
  y := b.y.Simplify(env)
 
  switch b.op {
  case '+':
    xt, ok1 := x.(Literal)
    yt, ok2 := y.(Literal)
    if ok1 && ok2 {
      return xt + yt
    } else {
      return binary{b.op, x, y}
    } 
   
  case '-':
    xt, ok1 := x.(Literal)
    yt, ok2 := y.(Literal)
    if ok1 && ok2 {
      return xt - yt
    } else {
      return binary{b.op, x, y}
    } 
   
  case '*':
    xt, ok1 := x.(Literal)
    yt, ok2 := y.(Literal)
    if ok1 && ok2 {
      return xt * yt
    } else {
      return binary{b.op, x, y}
    } 
   
  case '/':
    xt, ok1 := x.(Literal)
    yt, ok2 := y.(Literal)
    if ok1 && ok2 {
      return xt / yt
    } else {
      return binary{b.op, x, y}
    }
  default:
    panic("Unexpected operator")
  }

}

func (m measure) Simplify(env Env) Expr {
  // don't need to implement this
  panic("cannot simplify measure expression")
}


