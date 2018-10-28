package analysis

import (
  "go/ast"
  "go/parser"
  "go/token"
)

func cyclomatic(node ast.Stmt) uint {
  switch v := node.(type){
  case *ast.BlockStmt:
    var y uint = 1
    for i := 0; i < len(v.List); i++ {
      x := cyclomatic(v.List[i])
      y *= x
    }
    return y
  case *ast.CaseClause:
    var y uint = 1
    for i := 0; i < len(v.Body); i++ {
      x := cyclomatic(v.Body[i])
      y *= x
    }
    return y
  case *ast.SwitchStmt:
    var y uint = 0
    var check bool = false
    for i := 0; i < len(v.Body.List); i++ {
      x := cyclomatic(v.Body.List[i])
      y += x
    } 
    if v.Body.List == nil {
      check = true
    }
    if check {
      y += 1
    }
    return y
  case *ast.TypeSwitchStmt:
    var y uint = 0
    var check bool = false
    for i := 0; i < len(v.Body.List); i++ {
      x := cyclomatic(v.Body.List[i])
      y += x
    } 
    if v.Body.List == nil {
      check = true
    }
    if check {
      y += 1
    } 
    return y
  case *ast.IfStmt:
    return cyclomatic(v.Else) + cyclomatic(v.Body)
  case *ast.RangeStmt:
    return 1 + cyclomatic(v.Body)
  default:
    return 1
  }  
}

func CyclomaticComplexity(src string) map[string]uint {
  fset := token.NewFileSet()
  f, err := parser.ParseFile(fset, "src.go", src, 0)
  if err != nil {
    panic(err)
  }

  m := make(map[string]uint)
  for _, decl := range f.Decls {
    switch fn := decl.(type) {
    case *ast.FuncDecl:
      m[fn.Name.Name] = cyclomatic(fn.Body)
    }
  }

  return m
}
