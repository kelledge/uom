package usc

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/si"
)

var (
	PoundsForce = uom.DeriveUnit[dim.Force]("lbf",
		uom.U(PoundMass),
		uom.U(StandardGravity),
	)

	OuncesForce = uom.DeriveUnit[dim.Force]("ozf",
		uom.U(PoundsForce),
		uom.Const(1.0/16.0),
	)

	KipsForce = uom.DeriveUnit[dim.Force]("kip",
		uom.U(si.Kilo),
		uom.U(PoundsForce),
	)
)

var (
	FootPound = uom.DeriveUnit[dim.Torque]("ft·lb",
		uom.U(PoundsForce),
		uom.U(Foot),
	)

	InchPound = uom.DeriveUnit[dim.Torque]("in·lb",
		uom.U(PoundsForce),
		uom.U(Inch),
	)

	FootOunce = uom.DeriveUnit[dim.Torque]("ft·oz",
		uom.U(OuncesForce),
		uom.U(Foot),
	)

	InchOunce = uom.DeriveUnit[dim.Torque]("in·oz",
		uom.U(OuncesForce),
		uom.U(Inch),
	)
)

var (
	// SlugMass: mass unit where 1 lbf accelerates 1 slug at 1 ft/s^2
	SlugMass = uom.DeriveUnit[dim.Mass]("slug",
		uom.U(PoundsForce),
		uom.U(si.Second).Pow(2),
		uom.U(Foot).Per(),
	)
)

var (
	PoundsWeight = uom.DeriveUnit[dim.Weight]("lbf",
		uom.U(PoundMass),
		uom.U(StandardGravity),
	)
)
