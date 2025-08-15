package dim

import "github.com/kelledge/uom"

// Displacement: vector length (same dimension as Length, distinct semantic tag)
type Displacement struct{}

func (Displacement) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Length: 1}) }
func (Displacement) Name() string          { return "Displacement" }

type Velocity struct{}

func (Velocity) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 1, Time: -1})
}
func (Velocity) Name() string { return "Velocity" }

type Acceleration struct{}

func (Acceleration) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 1, Time: -2})
}
func (Acceleration) Name() string { return "Acceleration" }

// Jerk: third derivative of displacement (L/T^3)
type Jerk struct{}

func (Jerk) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 1, Time: -3})
}
func (Jerk) Name() string { return "Jerk" }

// Snap (a.k.a. Jounce): fourth derivative of displacement (L/T^4)
type Snap struct{}

func (Snap) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 1, Time: -4})
}
func (Snap) Name() string { return "Snap" }

// AngularVelocity: radians per second (radian is dimensionless → 1/T)
type AngularVelocity struct{}

func (AngularVelocity) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Time: -1})
}
func (AngularVelocity) Name() string { return "AngularVelocity" }

// AngularAcceleration: radians per second squared (1/T^2)
type AngularAcceleration struct{}

func (AngularAcceleration) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Time: -2})
}
func (AngularAcceleration) Name() string { return "AngularAcceleration" }
