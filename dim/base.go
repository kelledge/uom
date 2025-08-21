package dim

import "github.com/kelledge/uom"

type (
	Dimensionless struct{}
	Length        struct{}
	Time          struct{}
	Temperature   struct{}
	Mass          struct{}
	Current       struct{}
	Substance     struct{}
	Luminosity    struct{}
	Money         struct{}
)

func (Dimensionless) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{}) }
func (Dimensionless) Name() string          { return "Dimensionless" }

func (Length) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Length: 1}) }
func (Length) Name() string          { return "Length" }

func (Time) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Time: 1}) }
func (Time) Name() string          { return "Time" }

func (Temperature) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Temperature: 1}) }
func (Temperature) Name() string          { return "Temperature" }

func (Mass) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Mass: 1}) }
func (Mass) Name() string          { return "Mass" }

func (Current) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Current: 1}) }
func (Current) Name() string          { return "Current" }

func (Substance) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Substance: 1}) }
func (Substance) Name() string          { return "Substance" }

func (Luminosity) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Luminosity: 1}) }
func (Luminosity) Name() string          { return "Luminosity" }

func (Money) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Money: 1}) }
func (Money) Name() string          { return "Money" }

//

type Volume struct{}

func (Volume) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Length: 3}) }
func (Volume) Name() string          { return "Volume" }

type Area struct{}

func (Area) Dimension() uom.DimInt {
	return uom.NewDimInt(uom.DimensionSpec{Length: 2})
}
func (Area) Name() string { return "Area" }

type Angle struct{}

func (Angle) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{}) }
func (Angle) Name() string          { return "Angle" }
