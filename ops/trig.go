package ops

import (
	"math"

	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/std"
)

// Trigonometric functions integrate with uom by enforcing dimensional rules:
//   • sin/cos/tan/sec/csc/cot: accept Angle, return dimensionless (Scalar).
//   • asin/acos/atan: accept dimensionless, return Angle (radians).
//   • sinh/cosh/tanh: accept dimensionless, return dimensionless.

func Sin(theta uom.Quantity[dim.Angle]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Sin(uom.As(theta, std.Radian).Value()))
}

func Cos(theta uom.Quantity[dim.Angle]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Cos(uom.As(theta, std.Radian).Value()))
}

func Tan(theta uom.Quantity[dim.Angle]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Tan(uom.As(theta, std.Radian).Value()))
}

func Sec(theta uom.Quantity[dim.Angle]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(1 / math.Cos(uom.As(theta, std.Radian).Value()))
}

func Csc(theta uom.Quantity[dim.Angle]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(1 / math.Sin(uom.As(theta, std.Radian).Value()))
}

func Cot(theta uom.Quantity[dim.Angle]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(1 / math.Tan(uom.As(theta, std.Radian).Value()))
}

// // Inverse trig functions return angles (canonical: radians).
func Asin(x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Angle] {
	return std.Radian.Of(math.Asin(x.Value()))
}

func Acos(x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Angle] {
	return std.Radian.Of(math.Acos(x.Value()))
}

func Atan(x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Angle] {
	return std.Radian.Of(math.Atan(x.Value()))
}

// Hyperbolic trig functions, dimensionless in/out.
func Sinh(x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Sinh(x.Value()))
}

func Cosh(x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Cosh(x.Value()))
}

func Tanh(x uom.Quantity[dim.Dimensionless]) uom.Quantity[dim.Dimensionless] {
	return std.Scalar.Of(math.Tanh(x.Value()))
}
