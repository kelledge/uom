package dim

import "github.com/kelledge/uom"

// Frequency: 1/T (hertz)
type Frequency struct{}

func (Frequency) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Time: -1})
}
func (Frequency) Name() string { return "Frequency" }

// AngularFrequency: radians per second (same dims as Frequency; different semantics)
type AngularFrequency struct{}

func (AngularFrequency) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Time: -1})
}
func (AngularFrequency) Name() string { return "AngularFrequency" }

// Period: T (inverse of frequency)
type Period struct{}

func (Period) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Time: 1})
}
func (Period) Name() string { return "Period" }

// Wavelength: L (distance per cycle)
type Wavelength struct{}

func (Wavelength) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 1})
}
func (Wavelength) Name() string { return "Wavelength" }

// WaveNumber: 1/L (cycles per unit length)
type WaveNumber struct{}

func (WaveNumber) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: -1})
}
func (WaveNumber) Name() string { return "WaveNumber" }
