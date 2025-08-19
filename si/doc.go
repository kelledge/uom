// Package si defines International System of Units (SI) base and derived units.
// These form the canonical system for scientific and engineering work.
//
// Canonical units provided here include:
//
//   - Length: Meter (m)
//   - Mass:   Kilogram (kg)
//   - Time:   Second (s) [from std]
//   - Current: Ampere (A)
//   - Temperature: Kelvin (K), with Celsius (°C) for absolute and delta forms
//   - Amount: Mole (mol)
//   - Luminous intensity: Candela (cd)
//
// Derived units such as Newton, Pascal, Joule, Watt, and derived prefixes
// (milli-, kilo-, mega-, etc.) are also defined here.
//
// Example usage:
//
//	dist := si.Meter.Of(250)
//	mass := si.Kilogram.Of(80)
//	force := si.Newton.Of(500)
//
// The si package is the primary target for most scientific and engineering
// calculations, and interoperates cleanly with std and usc units.
package si
