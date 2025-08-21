package uom

import (
	"fmt"
)

// TODO: Resolve duplication with ops.ExprNode
type ExprNode interface {
	Dimension() DimInt  // dimension of the value/expression
	Canonical() float64 // scalar in canonical units (generally SI)
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

func AsCanonical[T Dimension](n ExprNode) Quantity[T] {
	panic("unimplemented")
}
