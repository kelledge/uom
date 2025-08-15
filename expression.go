package uom

import (
	"fmt"
	"math"
)

// ExprNode is the minimal, typed surface that arithmetic ops consume.
// Both Quantity[T] and Expr implement this.
type ExprNode interface {
	Dimension() DimInt  // dimension of the value/expression
	Canonical() float64 // scalar in canonical units (generally SI)
}

type Expr[T Dimension] struct {
	val float64 // canonical value
	dim DimInt  // resulting dimension
}

func One[T Dimension]() Expr[T] {
	return Expr[T]{
		val: 1.0,
		dim: DefaultDimCodec.MustEncode(DimensionSpec{}),
	}
}

func E[T Dimension](x ...ExprNode) Expr[T] {
	if len(x) > 1 {
		panic("")
	}

	if len(x) == 0 {
		return Expr[T]{val: 0, dim: DimInt(0)}
	}

	return Expr[T]{val: x[0].Canonical(), dim: x[0].Dimension()}
}

func (e Expr[T]) Dimension() DimInt  { return e.dim }
func (e Expr[T]) Canonical() float64 { return e.val }

func (e Expr[T]) Mul(xs ...ExprNode) Expr[T] {
	val := e.Canonical()
	dim := e.Dimension()

	for _, x := range xs {
		if x == nil {
			panic("expression received nil expression node")
		}
		val *= x.Canonical()
		dim += x.Dimension()
	}

	return Expr[T]{val: val, dim: dim}
}

func (e Expr[T]) Div(xs ...ExprNode) Expr[T] {
	val := e.Canonical()
	dim := e.Dimension()

	for _, x := range xs {
		if x == nil {
			panic("expression received nil expression node")
		}
		val /= x.Canonical()
		dim -= x.Dimension()
	}

	return Expr[T]{val: val, dim: dim}
}

func (e Expr[T]) Pow(n int) Expr[T] {
	return Expr[T]{
		val: math.Pow(e.Canonical(), float64(n)),
		dim: DimInt(n) * e.Dimension(),
	}
}

func (e Expr[T]) Inv() Expr[T] {
	return Expr[T]{
		val: 1.0 / e.Canonical(),
		dim: -1 * e.Dimension(),
	}
}

func (e Expr[T]) As(dest Unit[T]) Quantity[T] {
	return As(e, dest)
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
