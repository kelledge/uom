package dim

import "github.com/kelledge/uom"

// Stress: force per unit area → M·L^-1·T^-2 (same dimension as Pressure)
type Stress struct{}

func (Stress) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -1, Time: -2})
}
func (Stress) Name() string { return "Stress" }

// Strain: dimensionless (ratio of lengths)
type Strain struct{}

func (Strain) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{}) // dimensionless
}
func (Strain) Name() string { return "Strain" }

// YoungsModulus (Elastic modulus E): ratio Stress/Strain → same as Stress
type YoungsModulus struct{}

func (YoungsModulus) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -1, Time: -2})
}
func (YoungsModulus) Name() string { return "YoungsModulus" }

// ShearModulus (Modulus G): same dimensions as Stress
type ShearModulus struct{}

func (ShearModulus) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -1, Time: -2})
}
func (ShearModulus) Name() string { return "ShearModulus" }

// BulkModulus (K): same dimensions as Stress
type BulkModulus struct{}

func (BulkModulus) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: 1, Length: -1, Time: -2})
}
func (BulkModulus) Name() string { return "BulkModulus" }

// ElasticCompliance (S): inverse of modulus → 1/Stress → L·T^2/M
type ElasticCompliance struct{}

func (ElasticCompliance) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Mass: -1, Length: 1, Time: 2})
}
func (ElasticCompliance) Name() string { return "ElasticCompliance" }

// PoissonsRatio (ν): dimensionless
type PoissonsRatio struct{}

func (PoissonsRatio) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{}) // dimensionless
}
func (PoissonsRatio) Name() string { return "PoissonsRatio" }
