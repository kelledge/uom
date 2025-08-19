package uom

import "fmt"

var DefaultRegistry *Registry

type Registry struct {
	// byDim     map[DimInt][]UnitTerm
	// byToken   map[string]tokenRef
	// canonical map[DimInt]int
	// frozen    bool
}

// type tokenRef struct {
// 	dim   DimInt
// 	index int
// }

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) RegisterUnit(u UnitTerm, token ...string) {}

//
//
//

func DefineUnitIn[T Dimension](reg *Registry, name string, factor float64) Unit[T] {
	var dim T
	return Unit[T]{name: name, factor: factor, dim: dim}
}

func DeriveUnitIn[T Dimension](reg *Registry, name string, terms ...UnitTerm) Unit[T] {
	var t T

	dim := DimInt(0)
	factor := 1.0

	for _, term := range terms {
		dim += term.Dimension()
		factor *= term.ConversionFactor()
	}

	if dim != t.Dimension() {
		panic(fmt.Sprintf("dimension error: failed to derive unit: target dimensions: %q, term dimensions: %q", t.Dimension(), dim))
	}

	return Unit[T]{name: name, factor: factor, dim: t}
}

func DefineUnit[T Dimension](name string, factor float64) Unit[T] {
	return DefineUnitIn[T](DefaultRegistry, name, factor)
}

// DeriveUnit: multiplies factors/dimensions from the provided Units,
// panics if the resulting dimension does not match T{}.DimInt().
func DeriveUnit[T Dimension](name string, terms ...UnitTerm) Unit[T] {
	return DeriveUnitIn[T](DefaultRegistry, name, terms...)
}

func init() {
	DefaultRegistry = NewRegistry()
}
