package eval


func (v Var) Depth() uint {
  //panic("TODO: implement this!")
  return + 1;
}

func (f Literal) Depth() uint {
  //panic("TODO: implement this!")
  return + 1;
}

func (u unary) Depth() uint {
  //panic("TODO: implement this!")
  return u.x.Depth() + 1
}

func (b binary) Depth() uint {
  //panic("TODO: implement this!")
  if b.y.Depth() > b.x.Depth() {
      return b.y.Depth() + 1
  } else {
      return b.x.Depth() + 1
  }
}

func (m measure) Depth() uint {
  //panic("TODO: implement this!")
  return m.x.Depth() + 1;
}

