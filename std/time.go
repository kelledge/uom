package std

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	Day         = uom.DefineUnit[dim.Time]("day", 86400)
	Hour        = uom.DefineUnit[dim.Time]("hour", 3600)
	Minute      = uom.DefineUnit[dim.Time]("minute", 60)
	Second      = uom.DefineUnit[dim.Time]("second", 1)
	Millisecond = uom.DefineUnit[dim.Time]("millisecond", 1e-3)
	Microsecond = uom.DefineUnit[dim.Time]("microsecond", 1e-6)
	Nanosecond  = uom.DefineUnit[dim.Time]("nanosecond", 1e-9)
)
