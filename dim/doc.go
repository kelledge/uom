// Package dim defines dimension tags used by the units-of-measurement system.
// Each tag encodes the exponents of a physical dimension (e.g., length^1 ·
// time^-2) and serves as the type parameter for Unit[T] and Quantity[T].
//
// The purpose of this package is to provide a type-safe foundation for unit
// algebra. Dimensions are declared as empty types that implement the Dimension
// interface, and each one maps to a unique DimInt encoding.
//
// Canonical dimensions include:
//
// - Length
// - Time
// - Temperature
// - Mass
// - Current
// - Substance
// - Luminosity
// - Money
//
// Additional derived dimensions (e.g., Force, Pressure, Energy, Power) are
// also defined, grouped into topical files such as kinematics.go, dynamics.go,
// mechanics.go, and materials.go.
//
// Example usage:
//
//	var dist uom.Quantity[dim.Length] = si.Meter.Of(10)
//	var time uom.Quantity[dim.Time]   = si.Second.Of(2)
//	var vel  uom.Quantity[dim.Velocity] = dist.Div(time)
//
// The dim package itself contains no units or values. It exists purely to
// provide the dimension types that all units and quantities are parameterized
// over.
package dim
