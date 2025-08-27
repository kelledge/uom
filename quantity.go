package uom

import (
	"fmt"
	"strconv"
	"strings"
)

// Represents
type Quantity[T Dimension] struct {
	value float64 // scalar in the specified unit
	unit  Unit[T] // the unit
}

func (q Quantity[T]) Value() float64     { return q.value }
func (q Quantity[T]) Dimension() DimInt  { return q.unit.Dimension() }
func (q Quantity[T]) Canonical() float64 { return q.value * q.unit.ConversionFactor() }

func (q Quantity[T]) MarshalText() ([]byte, error) {
	// Use the primary name of the unit for stable output.
	// This could be the "preferred alias" registered with DefaultRegistry.
	name := q.unit.name

	return []byte(fmt.Sprintf("%g %s", q.value, name)), nil
}

func (q *Quantity[T]) UnmarshalText(b []byte) error {
	s := strings.TrimSpace(string(b))
	if s == "" {
		return fmt.Errorf("uom: empty quantity string")
	}

	parts := strings.Fields(s)
	if len(parts) != 2 {
		return fmt.Errorf("uom: invalid quantity format %q", s)
	}

	valStr, token := parts[0], parts[1]

	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return fmt.Errorf("uom: invalid numeric value %q", valStr)
	}

	var dim T
	term, ok := DefaultRegistry.Lookup(dim.Dimension(), token)
	if !ok {
		return fmt.Errorf("uom: unknown unit %q for dimension %s", token, dim.Dimension())
	}

	*q = Quantity[T]{value: val, unit: AsUnit[T](term)}
	return nil
}
