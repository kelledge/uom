package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

// Volumes commonly used across fluids, lab work, and large-scale geophysics
var (
	// Canonical SI volume; baseline for storage tanks, flow integrals, and conversions
	CubicMeter = uom.DefineUnit[dim.Volume]("m^3", 1)

	// SI-accepted unit favored in fluids/lab specs; ubiquitous in data sheets and dosing
	Liter = uom.DefineUnit[dim.Volume]("L", 1e-3)

	// Bench-scale measurements; reagents, small containers, pipetting
	Milliliter = uom.DefineUnit[dim.Volume]("mL", 1e-6)

	// Microlab and analytical work; microfluidics, chromatography, assays
	Microliter = uom.DefineUnit[dim.Volume]("µL", 1e-9)

	// Engineering drawings and engine displacement; also packaging volumes
	CubicCentimeter = uom.DefineUnit[dim.Volume]("cm^3", 1e-6)

	// Small parts and precision manufacturing; solder paste, adhesives, beads
	CubicMillimeter = uom.DefineUnit[dim.Volume]("mm^3", 1e-9)

	// Hydrology and geophysics; reservoir and basin volumes
	CubicKilometer = uom.DefineUnit[dim.Volume]("km^3", 1e9)

	// Often seen in chemistry texts and metrology; numerically equal to 1 L
	CubicDecimeter = uom.DefineUnit[dim.Volume]("dm^3", 1e-3)
)
