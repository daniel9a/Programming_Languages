package eval

// UnitKind defines the different classes of units that we work with.
// Multiple units can have the same kind, e.g. meters and miles are both
// units of length. Units of different kinds can only interact in restricted
// ways.
type UnitKind uint
const (
  Scalar = iota
  Length
  Time
  Mass
  Temperature
  Volume
  LT
  LV
)

// Unit represents the measurement scale that a quantity sits on.
// For a unit `u`, `scale` defines the size of `1 u` relative to
// some base unit. (If `u` is the base unit, `scale` is just 1.)
// Similarly, `offset` defines the distance that `0 u` is from the 0
// of some base unit. (If `u` is the base unit, `offset` is just 0.)
type Unit struct {
  kind UnitKind
  scale float64
  offset float64
}

// This is our table of units, describing their relative magnitudes
// and offsets from each other. Note that, for instance, `1 km` is defined
// to be 1000 of the base unit, and `1 m` is defined as 1 of the base unit.
// So `1 km` is `1000 m`. Furthermore, zero meters is the same as zero
// kilometers, so the offset of `km` is 0.
//
// Temperature units are relatively unique in that they have distinct offsets.
// 0 degrees Celsius is 273.15 degrees Kelvin, but an increase of one degree
// Celsius is the same as an increase of one degree Kelvin. The relationship
// between Fahrenheit and Kelvin is the most complex, as its degrees are smaller
// and its zero point is higher.
var units = map[string]Unit{
  "scalar": {Scalar, 1, 0},

  // Length
  "m": {Length, 1, 0},
  "km": {Length, 1000, 0},
  "mi": {Length, 1609.344, 0},

  // Time
  "s": {Time, 1, 0},
  "ms": {Time, 0.001, 0},
  "min": {Time, 60, 0},

  // Mass
  "kg": {Mass, 1, 0},
  "g": {Mass, 0.001, 0},
  "lbs": {Mass, 0.45359237, 0},

  // Temperature
  "K": {Temperature, 1, 0},
  "C": {Temperature, 1, 273.15},
  "F": {Temperature, 5.0/9.0, 459.67},

  // Volume
  "ltr": {Volume, 1, 0},
  "gal": {Volume, 0.264172, 0},
    
  //LT
  "m_p_s": {LT, 1.0/1.0, 0},
  "km_p_s": {LT, 1000.0/1.0, 0},
  "mi_p_s": {LT, 1609.344/1.0, 0}, 
    
  "m_p_ms": {LT, 1.0/0.001, 0},
  "m_p_min": {LT, 1.0/60.0, 0},
    
  "km_p_ms": {LT, 1000.0/0.001, 0},
  "km_p_min": {LT, 1000.0/60.0, 0},
    
  "mi_p_ms": {LT, 1609.344/0.001, 0},
  "mi_p_min": {LT, 1609.344/60, 0},
    
  //LV
  "m_p_ltr": {LV, 1, 0},
  "m_p_gal": {LV, 1.0/3.78541, 0},
    
  "km_p_ltr": {LV, 1000.0/1.0, 0},
  "km_p_gal": {LV, 1000.0/3.78541, 0},
    
  "mi_p_ltr": {LV, 1609.344/1.0, 0},
  "mi_p_gal": {LV, 1609.344/3.78541, 0},
    
}


func (l Literal) FlattenUnits() (Expr, string) {
  // Literals (like 5 or 3.1415) are scalar, i.e. unitless.
  return l, "scalar"
}

func (v Var) FlattenUnits() (Expr, string) {
  // Variables are scalar.
  return v, "scalar"
}

func (u unary) FlattenUnits() (Expr, string) {
  // Unary expressions take on the unit of their subexpression.
  x, x_unit := u.x.FlattenUnits()
  return unary{u.op, x}, x_unit
}

func (b binary) FlattenUnits() (Expr, string) {
  switch b.op {
  case '+', '-':
    // Addition and subtraction take on the unit of their left subexpression.
    // We convert the right subexpression to the unit of the left subexpression
    // by applying an extra measure operator to it.
    x, x_unit := b.x.FlattenUnits()
    y, _      := measure{x_unit, b.y}.FlattenUnits()

    return binary{b.op, x, y}, x_unit
  case '*', '/':
    // Multiplication and division are more complex, since any pair of units
    // can combine to become a totally new unit. In this sample code, we avoid
    // the question entirely. In this assignment, the task is to include limited
    // support for some composite units of measurement.
    x, x_unit := b.x.FlattenUnits()
    y, y_unit := b.y.FlattenUnits()

    y_scale := units[y_unit]
    x_scale := units[x_unit]

    if y_scale.kind == x_scale.kind {
      panic("Unexpected operator")
    }

    // If one subexpression is scalar, we take the unit of the other subexpression.
    if units[y_unit].kind == Scalar {
      return binary{b.op, x, y}, x_unit
    } else if units[x_unit].kind == Scalar && b.op == '*' {
      return binary{b.op, x, y}, y_unit
    } else if units[x_unit].kind == Scalar && b.op == '/'{
      panic("Unexpected operator")
    } else { 
      return binary{b.op, x, y}, x_unit+"_p_"+y_unit
    } 
  default:
    panic("Unexpected operator")
  }
}

func (m measure) FlattenUnits() (Expr, string) {
  // Measure operators must be removed and potentially replaced with
  // sufficient arithmetic to convert from the units of the subexpressions
  // to the given unit of measurement.
  x, x_unit := m.x.FlattenUnits()

  m_scale := units[m.unit]
  x_scale := units[x_unit]

  if x_unit == "scalar" {
    // Scalars are unitless, so we can assign a unit of measurement
    // without performing any conversion arithmetic.
    return x, m.unit
  } else if x_scale.kind != m_scale.kind {
    // Only like-kinded units can be inter-converted. Lengths cannot be
    // converted into times without mastering relativity in physics.
    panic("Can't convert between units of different kinds")
  } else if x_scale.scale == m_scale.scale && x_scale.offset == m_scale.offset {
    // If both units are the same, we don't have to do any conversion.
    // In this case, m.unit is just a synonym for x_unit, e.g. "s" and "sec".
    return x, m.unit
  } else if x_scale.scale == m_scale.scale {
    // If the offsets are distinct but the scales are the same, we only need to
    // translate the value to be relative to the new zero point.
    return binary{'+', x, Literal(x_scale.offset - m_scale.offset)}, m.unit
  } else {
    // If the scales are distinct, we have to do a full-fledged conversion.
    return binary{'+',
      binary{'*', x, Literal(x_scale.scale/m_scale.scale)},
      Literal(x_scale.offset*x_scale.scale/m_scale.scale - m_scale.offset),
    }, m.unit
  }
}