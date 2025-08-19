// Package std defines standard, system-agnostic units shared between SI and
// US customary systems. These are the universal anchors for common dimensions
// such as time, angle, and dimensionless scalars.
//
// Canonical units provided here include:
//
//   - Time: Second (s), with convenience units Millisecond, Minute, Hour
//   - Angle: Radian (rad), with Degree (°) as a derived unit
//   - Dimensionless: One (scalar), plus helpers like Percent (%) and PPM
//
// The design intent is that std hosts only those units that are universally
// recognized and not specific to any one measurement system. System-specific
// packages such as si or usc then build on these shared anchors.
//
// Example usage:
//
//	gain  := std.One.Of(0.75)        // dimensionless scalar
//	dt    := std.Millisecond.Of(12)  // 12 ms
//	angle := std.Degree.Of(30)       // 30 degrees
//
// These units are safe to import in both SI and USC contexts without creating
// duplication or ambiguity.
package std
