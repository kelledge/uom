package dim

import "github.com/kelledge/uom"

// Energy: M·L²/T² (Joule)
type Energy struct{}

func (Energy) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 2, Time: -2, Mass: 1})
}
func (Energy) Name() string { return "Energy" }

// Power: M·L²/T³ (Watt)
type Power struct{}

func (Power) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 2, Time: -3, Mass: 1})
}
func (Power) Name() string { return "Power" }

// Work: same dimension as Energy (Force·Displacement)
type Work struct{}

func (Work) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 2, Time: -2, Mass: 1})
}
func (Work) Name() string { return "Work" }

// Heat (Thermal Energy): same dimension as Energy
type Heat struct{}

func (Heat) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 2, Time: -2, Mass: 1})
}
func (Heat) Name() string { return "Heat" }

// Enthalpy: same dimension as Energy (state function; convenient semantic tag)
type Enthalpy struct{}

func (Enthalpy) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 2, Time: -2, Mass: 1})
}
func (Enthalpy) Name() string { return "Enthalpy" }

// Thermal conductivity: ability of a material to conduct heat
// W/(m·K) = kg·m/(s³·K)
type ThermalConductivity struct{}

func (ThermalConductivity) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 1, Mass: 1, Time: -3, Temperature: -1})
}
func (ThermalConductivity) Name() string { return "ThermalConductivity" }

// Heat capacity: total energy required to change temperature by 1 K
// J/K = kg·m²/(s²·K)
type HeatCapacity struct{}

func (HeatCapacity) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 2, Mass: 1, Time: -2, Temperature: -1})
}
func (HeatCapacity) Name() string { return "HeatCapacity" }

// Specific heat capacity: per-unit-mass heat capacity
// J/(kg·K) = m²/(s²·K)
type SpecificHeatCapacity struct{}

func (SpecificHeatCapacity) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 2, Time: -2, Temperature: -1})
}
func (SpecificHeatCapacity) Name() string { return "SpecificHeatCapacity" }

// Heat transfer coefficient: rate of heat transfer per area per temperature difference
// W/(m²·K) = kg/(s³·K)
type HeatTransferCoefficient struct{}

func (HeatTransferCoefficient) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Time: -3, Temperature: -1})
}
func (HeatTransferCoefficient) Name() string { return "HeatTransferCoefficient" }

type Pressure struct{}

func (Pressure) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -1, Time: -2})
}
func (Pressure) Name() string { return "Pressure" }
