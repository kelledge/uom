package uom

import "fmt"

// Unit: a base or directly defined unit of a single dimension.
// Example: Meter, Second, Newton.
type Unit[T Dimension] struct {
	name   string
	factor float64 // conversion factor to canonical unit for T
	dim    T
}

func (u Unit[T]) Name() string              { return u.name }
func (u Unit[T]) ConversionFactor() float64 { return u.factor }
func (u Unit[T]) Dimension() DimInt         { return u.dim.Dimension() }

// DefineUnit: constructor for SimpleUnit.
func DefineUnit[T Dimension](name string, factor float64) Unit[T] {
	var dim T
	return Unit[T]{name: name, factor: factor, dim: dim}
}

// DeriveUnit: multiplies factors/dimensions from the provided Units,
// panics if the resulting dimension does not match T{}.DimInt().
func DeriveUnit[T Dimension](name string, terms ...UnitTerm) Unit[T] {
	var t T

	dim := DimInt(0)
	factor := 1.0

	for _, t := range terms {
		dim += t.Dimension()
		factor *= t.ConversionFactor()
	}

	if dim != t.Dimension() {
		panic(fmt.Sprintf("dimension error: failed to derive unit: target dimensions: %q, term dimensions: %q", t.Dimension(), dim))
	}

	return Unit[T]{name: name, factor: factor, dim: t}
}
