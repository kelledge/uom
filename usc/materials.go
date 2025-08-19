package usc

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

// Mass density (ρ): common specs for metals, plastics, fluids
var (
	PoundMassPerCubicFoot = uom.DeriveUnit[dim.Density]("lbm/ft^3",
		uom.U(PoundMass),
		uom.U(Foot).Pow(-3),
	)

	PoundMassPerCubicInch = uom.DeriveUnit[dim.Density]("lbm/in^3",
		uom.U(PoundMass),
		uom.U(Inch).Pow(-3),
	)

	PoundMassPerCubicYard = uom.DeriveUnit[dim.Density]("lbm/yd^3",
		uom.U(PoundMass),
		uom.U(Yard).Pow(-3),
	)

	SlugPerCubicFoot = uom.DeriveUnit[dim.Density]("slug/ft^3",
		uom.U(SlugMass),
		uom.U(Foot).Pow(-3),
	)
)

// Linear mass density (μ): cables, fibers, wire, textiles
var (
	PoundMassPerFoot = uom.DeriveUnit[dim.LinearMass]("lbm/ft",
		uom.U(PoundMass),
		uom.U(Foot).Per(),
	)

	PoundMassPerInch = uom.DeriveUnit[dim.LinearMass]("lbm/in",
		uom.U(PoundMass),
		uom.U(Inch).Per(),
	)
)

var (
	PSI = uom.DeriveUnit[dim.Pressure]("psi",
		uom.U(PoundsForce),
		uom.U(InchSquared).Per(),
	)
)
