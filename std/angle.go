package std

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	Radian = uom.DefineUnit[dim.Dimensionless]("rad", 1)
	Degree = uom.DeriveUnit[dim.Dimensionless]("deg", uom.U(Radian).DivScalar(180))
)
