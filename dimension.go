package uom

import (
	"fmt"
	"strings"
)

type Dimension interface {
	// Dimension returns the encoded dimensional identity.
	Dimension() DimInt

	// Name is a human-readable name for the dimension (e.g. "Length", "Force").
	Name() string
}

// DimensionSpec is a human-readable specification of a unit's dimension exponents.
//
// This struct exists solely as a clear and self-documenting way for callers
// to express the dimensional exponents they want. Each field corresponds to
// one of the eight base quantities in the chosen base measurement system
// (almost always SI, but other coherent systems are possible).
//
// This is *not* used internally for computation or storage; it is converted
// into an encoded form (DimInt) via NewDimInt(DimensionSpec) for fast operations.
//
// Limitations:
//   - Only integer exponents are supported. Fractional exponents such as
//     square roots (e.g., sqrt(Hz)) or cube roots cannot be represented in
//     this system. While such forms are common in some sensor specifications
//     and empirical formulas, they must be handled outside this dimensional
//     encoding.
//
// Example:
//
//	Pressure = mass¹ · length⁻¹ · time⁻²
//	d := DimensionSpec{Mass: 1, Length: -1, Time: -2}
//	di := NewDimInt(d)
type DimensionSpec struct {
	Length      int // exponent of the base length quantity
	Time        int // exponent of the base time quantity
	Temperature int // exponent of the base temperature quantity
	Mass        int // exponent of the base mass quantity
	Current     int // exponent of the base electric current quantity
	Substance   int // exponent of the base amount-of-substance quantity
	Luminosity  int // exponent of the base luminous intensity quantity
	Money       int // exponent of the base monetary quantity
}

// String returns a compact textual representation of the dimension,
// showing only the nonzero exponents in the form "L^1 T^-2".
func (d DimensionSpec) String() string {
	parts := []string{}
	if d.Length != 0 {
		parts = append(parts, fmt.Sprintf("L^%d", d.Length))
	}
	if d.Time != 0 {
		parts = append(parts, fmt.Sprintf("T^%d", d.Time))
	}
	if d.Temperature != 0 {
		parts = append(parts, fmt.Sprintf("Θ^%d", d.Temperature)) // Θ for temperature
	}
	if d.Mass != 0 {
		parts = append(parts, fmt.Sprintf("M^%d", d.Mass))
	}
	if d.Current != 0 {
		parts = append(parts, fmt.Sprintf("I^%d", d.Current))
	}
	if d.Substance != 0 {
		parts = append(parts, fmt.Sprintf("N^%d", d.Substance))
	}
	if d.Luminosity != 0 {
		parts = append(parts, fmt.Sprintf("J^%d", d.Luminosity)) // J for luminous intensity
	}
	if d.Money != 0 {
		parts = append(parts, fmt.Sprintf("$^%d", d.Money))
	}
	if len(parts) == 0 {
		return "scalar"
	}
	return strings.Join(parts, "·")
}
