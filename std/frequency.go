package std

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	Herz = uom.DefineUnit[dim.Frequency]("Hz", 1)
	RPM  = uom.DefineUnit[dim.Frequency]("SPM", 1.0/60.0)
)
