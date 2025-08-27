package uom

import "fmt"

// CanonicalSet is a mapping from dimension to its designated canonical unit.
//
// A CanonicalSet defines the "preferred" unit for each dimension under a
// particular system of measure (e.g. MKS, CGS, FPS). Canonical units are used
// whenever quantities are normalized or converted via AsCanonical.
type CanonicalSet map[DimInt]UnitTerm

// NewCanonicalSet constructs a CanonicalSet from the provided unit terms.
//
// Each UnitTerm must come from U(Unit[T]), which enforces dimension safety.
// NewCanonicalSet panics if more than one unit is provided for the same
// dimension.
//
// Example:
//
//	var MKS = uom.NewCanonicalSet(
//	    uom.U(si.Meter),
//	    uom.U(std.Second),
//	    uom.U(si.Kilogram),
//	)
func NewCanonicalSet(entries ...UnitTerm) CanonicalSet {
	if len(entries) == 0 {
		panic("cannot create empty CanonicalSet: no entries provided")
	}

	set := make(CanonicalSet, len(entries))

	for i, e := range entries {
		if prev, exists := set[e.dim]; exists {
			panic(fmt.Sprintf(
				"duplicate canonical for dimension %s at index %d: tried %q but %q is already defined",
				e.dim, i, e.name, prev.name,
			))
		}
		set[e.dim] = e
	}

	return set
}
