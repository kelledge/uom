package usc

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	// Standard acceleration due to gravity
	// Used for weight-force conversions (lbm ↔ lbf)
	// By definition: g₀ = 9.80665 m/s² ≈ 32.17405 ft/s²
	StandardGravity = uom.DeriveUnit[dim.Acceleration]("gn",
		uom.U(FootPerSecondSquared),
		uom.Const(32.1740485564304),
	)
)
