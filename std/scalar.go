package std

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	Scalar = uom.DefineUnit[dim.Dimensionless]("scalar", 1.0)
)
