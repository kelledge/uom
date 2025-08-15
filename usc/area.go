package usc

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	FootQuared  = uom.DeriveUnit[dim.Area]("ft^2", uom.U(Foot).Pow(2))
	InchSquared = uom.DeriveUnit[dim.Area]("in^2", uom.U(Inch).Pow(2))
)
