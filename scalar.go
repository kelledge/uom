package uom

type Unitless struct{}

func (Unitless) Dimension() DimInt { return NewDimInt(DimensionSpec{}) }
func (Unitless) Name() string      { return "Unitless" }

// Optional canonical unit “1”
var ScalarUnit = DefineUnit[Unitless]("scalar", 1)

// Constructor
func Scalar(x float64) Quantity[Unitless] {
	return Quantity[Unitless]{value: x, unit: ScalarUnit}
}
