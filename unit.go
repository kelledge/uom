package uom

// Unit: a base or directly defined unit of a single dimension.
// Example: Meter, Second, Newton.
type Unit[T Dimension] struct {
	name   string
	factor float64 // conversion factor to canonical unit for T
	dim    T
}

func (u Unit[T]) Name() string                 { return u.name }
func (u Unit[T]) ConversionFactor() float64    { return u.factor }
func (u Unit[T]) Dimension() DimInt            { return u.dim.Dimension() }
func (u Unit[T]) Of(value float64) Quantity[T] { return Quantity[T]{value: value, unit: u} }
