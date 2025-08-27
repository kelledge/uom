package ops

import (
	"fmt"
	"math"

	"github.com/kelledge/uom"
)

// REDEFINE of uom.ExprNode
type ExprNode interface {
	Dimension() uom.DimInt // Dimension returns the physical dimension of the value, encoded as a DimInt.
	Canonical() float64    // Canonical returns the value scaled into canonical (usually SI) units.
}

// Expr represents a composable expression that ultimately evaluates to a
// quantity of dimension T.
//
// An Expr is always created with a target dimension in mind, then built up
// through fluent arithmetic until it resolves to that target.
//
// This mirrors real-world calculations, which typically start from a goal
// (velocity, pressure, power, ...) and work through intermediate steps to
// reach it.
//
// New expressions are created using E()
type Expr[T uom.Dimension] struct {
	val float64    // canonical value
	dim uom.DimInt // resulting dimension
}

// E creates a new expression with target dimension T.
//
// It is the standard way to begin building an expression. Typically you
// lift an existing Quantity or Expr into the builder, then combine it with
// other terms until the result matches the intended dimension.
func E[T uom.Dimension](x ExprNode) Expr[T] {
	return Expr[T]{val: x.Canonical(), dim: x.Dimension()}
}

func (e Expr[T]) Dimension() uom.DimInt { return e.dim }
func (e Expr[T]) Canonical() float64    { return e.val }

// Add adds this expression to one or more others. All operands must share
// the same dimension T.
//
// Panics if called with no arguments.
func (e Expr[T]) Add(xs ...ExprNode) Expr[T] {
	if len(xs) == 0 {
		panic("Add requires at least one operand")
	}

	val := e.Canonical()
	dim := e.Dimension()

	for i, x := range xs {
		if x == nil {
			panic(fmt.Sprintf("Add received nil operand at index %d", i))
		}
		xd := x.Dimension()
		if xd != dim {
			panic(fmt.Sprintf("Add dimension mismatch: expr has %v, operand[%d] has %v", dim, i, xd))
		}
		val += x.Canonical()
	}
	return Expr[T]{val: val, dim: dim}
}

// Sub subtracts one or more expressions from this one. All operands must
// share the same dimension T.
//
// Panics if called with no arguments.
func (e Expr[T]) Sub(xs ...ExprNode) Expr[T] {
	if len(xs) == 0 {
		panic("Sub requires at least one operand")
	}

	val := e.Canonical()
	dim := e.Dimension()

	for i, x := range xs {
		if x == nil {
			panic(fmt.Sprintf("Sub received nil operand at index %d", i))
		}
		xd := x.Dimension()
		if xd != dim {
			panic(fmt.Sprintf("Sub dimension mismatch: expr has %v, operand[%d] has %v", dim, i, xd))
		}
		val -= x.Canonical()
	}
	return Expr[T]{val: val, dim: dim}
}

// Mul multiplies this expression by one or more other operands.
// The resulting dimension is the sum of operand dimensions.
//
// Panics if called with no arguments, or if any operand is nil.
func (e Expr[T]) Mul(xs ...ExprNode) Expr[T] {
	if len(xs) == 0 {
		panic("Mul requires at least one operand")
	}

	val := e.Canonical()
	dim := e.Dimension()

	for i, x := range xs {
		if x == nil {
			panic(fmt.Sprintf("Mul received nil operand at index %d", i))
		}
		val *= x.Canonical()
		dim += x.Dimension()
	}

	return Expr[T]{val: val, dim: dim}
}

// Div divides this expression by one or more other operands.
// The resulting dimension is the difference of operand dimensions.
//
// Panics if called with no arguments, or if any operand is nil.
func (e Expr[T]) Div(xs ...ExprNode) Expr[T] {
	if len(xs) == 0 {
		panic("Div requires at least one operand")
	}

	val := e.Canonical()
	dim := e.Dimension()

	for i, x := range xs {
		if x == nil {
			panic(fmt.Sprintf("Div received nil operand at index %d", i))
		}
		val /= x.Canonical()
		dim -= x.Dimension()
	}

	return Expr[T]{val: val, dim: dim}
}

// Pow raises the expression to an integer power.
//
// The resulting dimension is scaled by the same exponent
// (e.g. (Length)^2 → Area).
//
// Special cases:
//   - n == 0: result is dimensionless with value 1, unless the base is zero, in which case Pow(0) panics.
func (e Expr[T]) Pow(n int) Expr[T] {
	if n == 0 {
		if e.Canonical() == 0 {
			panic("Pow(0): 0^0 is indeterminate")
		}
		return Expr[T]{
			val: 1.0,
			dim: uom.DimInt(0), // dimensionless
		}
	}

	// Negative and value is zero?

	return Expr[T]{
		val: math.Pow(e.Canonical(), float64(n)),
		dim: uom.DimInt(n) * e.Dimension(),
	}
}

// Inv returns the reciprocal of the expression (1 / e).
//
// The resulting dimension is the negation of the original
// (e.g. Length → 1/Length).
//
// Panics if the canonical value is zero.
func (e Expr[T]) Inv() Expr[T] {
	val := e.Canonical()
	if val == 0 {
		panic("Inv: division by zero")
	}
	return Expr[T]{
		val: 1.0 / val,
		dim: -1 * e.Dimension(),
	}
}

// AsCanonical materializes the expression as a Quantity in canonical units
// (generally SI).
//
// This is most often used when a Quantity is required but the specific
// choice of concrete unit does not matter. It lets authors avoid hardcoding
// concrete units like meters or feet in their implementations. For example:
//
//	func Span() uom.Quantity[dim.Length] {
//	    expr := ops.E[dim.Length](someValue).Add(otherValue)
//	    return expr.AsCanonical() // caller gets Length, no unit assumption
//	}
func (e Expr[T]) AsCanonical() uom.Quantity[T] {
	return uom.AsCanonical[T](e)
}

// As evaluates the expression as a Quantity in the given unit.
//
// Use this when you need the result expressed in a specific concrete unit.
// If you are unsure which to use, prefer AsCanonical.
func (e Expr[T]) As(dest uom.Unit[T]) uom.Quantity[T] {
	return uom.As(e, dest)
}
