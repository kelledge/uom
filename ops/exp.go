package ops

import (
	"math"

	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/std"
)

// Exp returns e^x for a dimensionless x.
func Exp(x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Exp(x.Value()))
}

// Ln returns the natural log of x for a dimensionless x.
func Ln(x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Log(x.Value()))
}

// Log10 returns log base 10 of x for a dimensionless x.
func Log10(x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Log10(x.Value()))
}

// Log returns log base b of x for dimensionless x and b.
func Log(b, x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Log(x.Value()) / math.Log(b.Value()))
}
