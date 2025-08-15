package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	StandardGravity = uom.DeriveUnit[dim.Acceleration]("g", uom.U(MeterPerSecondSquared), uom.Const(9.80665))
)
