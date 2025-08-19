// Package uom provides a type-safe system for working with physical units of
// measurement in Go.
//
// The goal of uom is to make physical calculations in Go safe,
// expressive, and efficient for real world software.
//
// Core concepts
//
//   - Dimension    - a type-level tag (from package dim) representing a physical dimension
//   - Unit[T]      - a named unit of measure with scale factor, dimension, and kind
//   - Quantity[T]  - a value paired with its unit
//   - UnitTerm     - building block for derived units and registry entries
//   - Expr         - a fluent builder for dimensionally-safe expressions with quantities
//
// Units are grouped into subpackages:
//
//   - std  - system-agnostic anchors (seconds, radians, scalars, percent, ...)
//   - si   - International System (meter, kilogram, pascal, ...)
//   - usc  - U.S. customary (foot, gallon, pound-force, ...)
//
// # Algebra
//
// Quantities of compatible dimensions can be added and subtracted. Multiplying
// and dividing quantities produces new derived quantities with appropriate
// dimensions. Affine units (e.g., temperatures) are planned, but not currently
// handled.
//
// # Example
//
//  // Setup the units
//  // In practice these would be imported
//  inch := uom.DefineUnit[dim.Length]("inch", 0.0254)
//  foot := uom.DeriveUnit[dim.Length]("foot", uom.U(inch).MulScalar(12))
//
//  footSquare := uom.DeriveUnit[dim.Area]("ft^2", uom.U(foot).Pow(2))
//
//  lbf := uom.DefineUnit[dim.Force]("lbf", 4.4482)
//  psi := uom.DeriveUnit[dim.Pressure]("psi", uom.U(lbf), uom.U(inch).Pow(-2))
//
//  // Calculate the force
//	buildingSide := footSquare.Of(1_000)
//  pressureReading := psi.Of(2.9)
//  totalForce := uom.Expr(buildingSide).Mul(pressureReading).As(lbf)
//
//  fmt.Printf("force on building side: %.2f lbf\n", totalForce.Value())

package uom
