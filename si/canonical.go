package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/std"
)

// MKS is the canonical set for the "meter–kilogram–second" system.
// It includes SI base units plus selected derived units.
var MKS = uom.NewCanonicalSet(
	// base units
	uom.U(Meter),      // length
	uom.U(Kilogram),   // mass
	uom.U(std.Second), // time
	uom.U(Ampere),     // electric current
	uom.U(Kelvin),     // thermodynamic temperature
	uom.U(Mole),       // amount of substance
	uom.U(Candela),    // luminous intensity
	uom.U(std.Radian), // angle

	// derived units
	uom.U(Newton), // force
	uom.U(Joule),  // energy
	uom.U(Watt),   // power
	uom.U(Pascal), // pressure

	// more
)
