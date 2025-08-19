// Package usc defines U.S. customary units commonly used in engineering,
// petroleum, and industrial contexts. These complement SI units by supporting
// measurements in feet, gallons, pounds, and related derived units.
//
// Canonical units provided here include:
//
//   - Length: Foot (ft), Inch (in)
//   - Volume: Gallon (gal), Barrel (bbl)
//   - Mass:   PoundMass (lbm)
//   - Force:  PoundForce (lbf)
//   - Pressure: psi (pound per square inch)
//
// Example usage:
//
//	length := usc.Foot.Of(10)
//	vol    := usc.Gallon.Of(5)
//	weight := usc.PoundForce.Of(2000)
//
// The usc package is intended for applications where U.S. customary measures
// are standard, while still interoperating with si and std quantities for
// consistent calculations and conversions.
package usc
