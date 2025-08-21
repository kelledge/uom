package std

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	Radian = uom.DefineUnit[dim.Angle]("rad", 1)
	Degree = uom.DeriveUnit[dim.Angle]("deg", uom.U(Radian).DivScalar(180))
)
