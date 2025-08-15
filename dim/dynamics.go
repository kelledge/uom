package dim

import "github.com/kelledge/uom"

// Force: M·L/T² (newton)
type Force struct{}

func (Force) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: 1, Time: -2})
}
func (Force) Name() string { return "Force" }

// Weight: gravitational force on a body (W = m·g) under a specified local g.
// Use this when you want to make dependence on gravity explicit; prefer dim.Force for generic loads.
type Weight struct{}

func (Weight) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 1, Time: -2, Mass: 1})
}
func (Weight) Name() string { return "Weight" }

// Torque (Moment of force): Force·Length → M·L²/T² (same dimension as Energy)
type Torque struct{}

func (Torque) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: 2, Time: -2})
}
func (Torque) Name() string { return "Torque" }

// Momentum (Linear): M·L/T
type Momentum struct{}

func (Momentum) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: 1, Time: -1})
}
func (Momentum) Name() string { return "Momentum" }

// AngularMomentum: M·L²/T
type AngularMomentum struct{}

func (AngularMomentum) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: 2, Time: -1})
}
func (AngularMomentum) Name() string { return "AngularMomentum" }

// Impulse: Force·Time → M·L/T (same dimension as Momentum)
type Impulse struct{}

func (Impulse) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: 1, Time: -1})
}
func (Impulse) Name() string { return "Impulse" }

// MomentOfInertia: M·L²
type MomentOfInertia struct{}

func (MomentOfInertia) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: 2})
}
func (MomentOfInertia) Name() string { return "MomentOfInertia" }
