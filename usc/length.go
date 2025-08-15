package usc

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	Inch = uom.DefineUnit[dim.Length]("in", 0.0254)
	Foot = uom.DefineUnit[dim.Length]("ft", 0.3048)
	Yard = uom.DefineUnit[dim.Length]("yd", 0.9144)
	Mile = uom.DefineUnit[dim.Length]("mi", 1609.344)
)
