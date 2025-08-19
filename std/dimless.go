package std

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	Each    = uom.DefineUnit[dim.Dimensionless]("each", 1)
	Percent = uom.DefineUnit[dim.Dimensionless]("percent", 1e-2)
	PPM     = uom.DefineUnit[dim.Dimensionless]("ppm", 1e-6)
)
