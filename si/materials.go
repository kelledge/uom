package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

// Mass density (ρ): mass per unit volume
var (
	KilogramPerCubicMeter = uom.DefineUnit[dim.Density]("kg/m^3", 1)

	GramPerCubicCentimeter = uom.DeriveUnit[dim.Density]("g/cm^3",
		uom.U(Gram),
		uom.U(CubicCentimeter).Per(),
	)

	GramPerMilliliter = uom.DeriveUnit[dim.Density]("g/mL",
		uom.U(Gram),
		uom.U(Milliliter).Per(),
	)
)

// Linear mass density (μ): mass per unit length
var (
	KilogramPerMeter = uom.DefineUnit[dim.LinearMass]("kg/m", 1)

	GramPerMeter = uom.DeriveUnit[dim.LinearMass]("g/m",
		uom.U(Gram),
		uom.U(Meter).Per(),
	)

	GramPerMillimeter = uom.DeriveUnit[dim.LinearMass]("g/mm",
		uom.U(Gram),
		uom.U(Millimeter).Per(),
	)
)

// Areal mass density (σ): mass per unit area
var (
	KilogramPerSquareMeter = uom.DefineUnit[dim.ArealMass]("kg/m^2", 1)

	GramPerSquareMeter = uom.DeriveUnit[dim.ArealMass]("g/m^2",
		uom.U(Gram),
		uom.U(Meter).Pow(-2),
	)
)
