package uom

import (
	"fmt"
	"math"
)

// UnitTerm retains the runtime dimension vector and conversion factor information but drops
// the connection to type safe dimensions.
//
// This is useful when deriving units programmatically via the DeriveUnit constructor.
// Go (thankfully) does not support variadic type parameters in generics
type UnitTerm struct {
	name   string
	factor float64
	dim    DimInt
}

// U: lift a typed unit into a term (one-time).
func U[T Dimension](u Unit[T]) UnitTerm {
	return UnitTerm{
		name:   u.Name(),
		factor: u.ConversionFactor(),
		dim:    u.Dimension(),
	}
}

func Const(v float64) UnitTerm {
	return UnitTerm{
		name:   "const",
		factor: v,
		dim:    NewDimInt(DimensionSpec{}),
	}
}

func (u *UnitTerm) Name() string              { return u.name }
func (u *UnitTerm) ConversionFactor() float64 { return u.factor }
func (u *UnitTerm) Dimension() DimInt         { return u.dim }

func (t UnitTerm) Pow(exp int) UnitTerm {
	return UnitTerm{
		name:   fmt.Sprintf("%s^%d", t.name, exp),
		factor: math.Pow(t.factor, float64(exp)),
		dim:    DimInt(exp) * t.dim,
	}
}

func (t UnitTerm) Per() UnitTerm {
	return UnitTerm{
		name:   "1/" + t.name,
		factor: 1 / t.factor,
		dim:    -1 * t.dim,
	}
}

// MulScalar multiplies the UnitTerm's factor by a scalar (dimensionless).
func (t UnitTerm) MulScalar(s float64) UnitTerm {
	return UnitTerm{
		name:   fmt.Sprintf("%g·%s", s, t.name),
		factor: t.factor * s,
		dim:    t.dim, // dimension unchanged
	}
}

// DivScalar divides the UnitTerm's factor by a scalar (dimensionless).
func (t UnitTerm) DivScalar(s float64) UnitTerm {
	return UnitTerm{
		name:   fmt.Sprintf("%s/%g", t.name, s),
		factor: t.factor / s,
		dim:    t.dim, // dimension unchanged
	}
}
