package dim

import "github.com/kelledge/uom"

// Density: M/L³ (mass per unit volume)
type Density struct{}

func (Density) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -3})
}
func (Density) Name() string { return "Density" }

// LinearMass: M/L (mass per unit length)
type LinearMass struct{}

func (LinearMass) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -1})
}
func (LinearMass) Name() string { return "LinearMass" }

// ArealMass: M/L² (mass per unit area)
type ArealMass struct{}

func (ArealMass) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -2})
}
func (ArealMass) Name() string { return "ArealMass" }

// SpecificVolume: L³/M (inverse of density)
type SpecificVolume struct{}

func (SpecificVolume) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: -1, Length: 3})
}
func (SpecificVolume) Name() string { return "SpecificVolume" }

// SpecificWeight: Force per unit volume → M/(L²·T²)
type SpecificWeight struct{}

func (SpecificWeight) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -2, Time: -2})
}
func (SpecificWeight) Name() string { return "SpecificWeight" }

// NumberDensity: 1/L³ (count per unit volume; dimensionless numerator)
type NumberDensity struct{}

func (NumberDensity) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: -3})
}
func (NumberDensity) Name() string { return "NumberDensity" }

// MassConcentration: M/L³ (same as density but semantic distinction)
type MassConcentration struct{}

func (MassConcentration) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -3})
}
func (MassConcentration) Name() string { return "MassConcentration" }

// AmountConcentration: N/L³ (moles per unit volume)
type AmountConcentration struct{}

func (AmountConcentration) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Substance: 1, Length: -3})
}
func (AmountConcentration) Name() string { return "AmountConcentration" }

// VolumetricFlowRate: L³/T
type VolumetricFlowRate struct{}

func (VolumetricFlowRate) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 3, Time: -1})
}
func (VolumetricFlowRate) Name() string { return "VolumetricFlowRate" }

// MassFlowRate: M/T
type MassFlowRate struct{}

func (MassFlowRate) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Time: -1})
}
func (MassFlowRate) Name() string { return "MassFlowRate" }

// Compressibility: 1/Pressure → L·T²/M
type Compressibility struct{}

func (Compressibility) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 1, Time: 2, Mass: -1})
}
func (Compressibility) Name() string { return "Compressibility" }
