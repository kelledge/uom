package uom

import "fmt"

var DefaultRegistry *Registry

type Registry struct {
	byDim     map[DimInt][]UnitTerm // dimension -> units in that dimension
	canonical CanonicalSet
}

func NewRegistry() *Registry {
	return &Registry{
		byDim:     make(map[DimInt][]UnitTerm),
		canonical: make(CanonicalSet),
	}
}

// RegisterUnit inserts a unit term for the given dimension.
// Will eventually index it by any aliases.
func (r *Registry) RegisterUnit(dim DimInt, term UnitTerm, aliases ...string) {
	r.byDim[dim] = append(r.byDim[dim], term)
}

// Lookup finds a unit in the registry for the given dimension and alias.
// Currently does a linear search. Just sorting out APIs at the moment and
// not terribly concerned with performance yet.
func (r *Registry) Lookup(dim DimInt, alias string) (UnitTerm, bool) {
	terms, ok := r.byDim[dim]
	if !ok {
		return UnitTerm{}, false
	}

	for _, term := range terms {
		if term.name == alias {
			return term, true
		}
	}

	return UnitTerm{}, false
}

// UseCanonicalSet specifies the set of concrete units to use for specific dimensions.
// The design is currently in flux and this likely to be the target of refactoring, as
// it smells.
func (r *Registry) UseCanonicalSet(canonical CanonicalSet) {
	r.canonical = canonical
}

func DefineUnitIn[T Dimension](reg *Registry, name string, factor float64) Unit[T] {
	var dim T

	u := Unit[T]{name: name, factor: factor, dim: dim}
	reg.RegisterUnit(dim.Dimension(), U(u))

	return u
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

	u := Unit[T]{name: name, factor: factor, dim: t}
	reg.RegisterUnit(dim, U(u))

	return u
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
