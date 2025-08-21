package uom

// Represents
type Quantity[T Dimension] struct {
	value float64 // scalar in the specified unit
	unit  Unit[T] // the unit; type-checked to match T at compile time
}

func (q Quantity[T]) Value() float64     { return q.value }
func (q Quantity[T]) Dimension() DimInt  { return q.unit.Dimension() }
func (q Quantity[T]) Canonical() float64 { return q.value * q.unit.ConversionFactor() }
