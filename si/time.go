package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

// Common prefixed time units
var (
	Millisecond = uom.DeriveUnit[dim.Time]("ms", uom.U(Milli), uom.U(Second))
	Microsecond = uom.DeriveUnit[dim.Time]("µs", uom.U(Micro), uom.U(Second))
	Nanosecond  = uom.DeriveUnit[dim.Time]("ns", uom.U(Nano), uom.U(Second))
)

var (
	Minute = uom.DefineUnit[dim.Time]("min", 60.0)
	Hour   = uom.DefineUnit[dim.Time]("h", 3600.0)
	Day    = uom.DefineUnit[dim.Time]("d", 86400.0)
)

var (
	Herz = uom.DefineUnit[dim.Frequency]("Hz", 1)
	RPM  = uom.DefineUnit[dim.Frequency]("SPM", 1.0/60.0)
)
