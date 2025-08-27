package uom

import (
	"fmt"
)

// TODO: Resolve duplication with ops.ExprNode
type ExprNode interface {
	Dimension() DimInt  // dimension of the value/expression
	Canonical() float64 // scalar in canonical units (generally SI)
}

// AsUnit returns a UnitTerm to a typed unit at runtime
func AsUnit[T Dimension](t UnitTerm) Unit[T] {
	var dim T

	expected := dim.Dimension()
	if t.dim != expected {
		panic(fmt.Sprintf("dimension mismatch: got %s, expected %s", t.dim, expected))
	}

	return Unit[T]{
		name:   t.name,
		factor: t.factor,
		dim:    dim,
	}
}

func As[T Dimension](n ExprNode, dest Unit[T]) Quantity[T] {
	if n.Dimension() != dest.Dimension() {
		panic(fmt.Sprintf("dimension mismatch: expression %q, destination %q", n.Dimension(), dest.Dimension()))
	}

	// q_s * f_s == q_canonical == q_d * f_d
	// (q_s * f_s) / f_d == q_d
	value := n.Canonical() / dest.ConversionFactor()

	q := Quantity[T]{
		value: value,
		unit:  dest,
	}

	return q
}

// AsCanonical evaluates the given expression or quantity and converts it
// into the canonical unit for its dimension, as defined by the active
// canonical set in DefaultRegistry.
//
// Panics if no canonical unit is defined for the dimension.
func AsCanonical[T Dimension](n ExprNode) Quantity[T] {
	dim := n.Dimension()

	unit, ok := DefaultRegistry.canonical[dim]
	if !ok {
		panic(fmt.Sprintf("no canonical unit defined for dimension %s", dim))
	}

	return As(n, AsUnit[T](unit))
}
