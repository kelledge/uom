package uom

import "fmt"

var DefaultDimCodec = MustDimCodec(DimSizes{64, 64, 64, 64, 64, 64, 64, 64})

// DimInt represents a compact, reversible encoding of a DimensionVector using a
// configurable mixed-radix number system.
//
// Each element of a DimensionVector is constrained to a fixed range defined by a
// corresponding DimSizes value (i.e., the number of allowed exponent values per dimension).
// These ranges are centered around zero and mapped to non-negative offsets using
// precomputed biases.
//
// The DimInt encoding is computed by offsetting each dimension's exponent to a
// non-negative value, multiplying by its corresponding radix weight (DimVals),
// and summing across all dimensions. An additional scalar bias (DimBias) is added
// to ensure the entire range of valid DimensionVectors maps into the positive integers.
//
// This encoding is deterministic, lossless, and invertible as long as all dimension
// exponents remain within their configured bounds.
type DimInt int64

func NewDimInt(spec DimensionSpec) DimInt {
	return DefaultDimCodec.MustEncode(spec)
}

func (d DimInt) String() string {
	spec := DefaultDimCodec.Decode(d)
	return spec.String()
}

// DimSizes defines the per-axis digit sizes (radices) for the mixed-radix encoding.
// Length is fixed at 8 to match the number of base quantities.
//
// Example literal:
//
//	sizes := DimSizes{64, 64, 64, 64, 64, 64, 64, 64}
type DimSizes [8]int

// Validate ensures each radix is even and >= 2 so zero is centered.
func (d DimSizes) Validate() error {
	for i, s := range d {
		if s < 2 || s%2 != 0 {
			return fmt.Errorf("dimsizes[%d]=%d must be even and >= 2", i, s)
		}
	}
	return nil
}

// Offsets returns the per-axis zero-centering offsets (size/2).
func (d DimSizes) Offsets() [8]int {
	var off [8]int
	for i, s := range d {
		off[i] = s / 2
	}
	return off
}

// DimVals returns the positional weights for this mixed-radix system.
// v[0] = 1; v[i] = v[i-1] * sizes[i-1].
func (d DimSizes) DimVals() [8]int64 {
	var v [8]int64
	v[0] = 1
	for i := 1; i < 8; i++ {
		v[i] = v[i-1] * int64(d[i-1])
	}
	return v
}

// Bias returns the constant that centers the zero exponent vector at a single code.
func (d DimSizes) Bias() int64 {
	off := d.Offsets()
	vals := d.DimVals()
	var b int64
	for i := 0; i < 8; i++ {
		b += int64(off[i]) * vals[i]
	}
	return b
}

// DimCodec binds a DimSizes configuration with its derived weights and bias.
// It provides safe Encode/Decode and bias-aware arithmetic helpers.
type DimCodec struct {
	Sizes DimSizes // per-axis radices (even, >= 2)
	Vals  [8]int64 // positional weights; Vals[0]=1, Vals[i]=Vals[i-1]*Sizes[i-1]
	Bias  int64    // centering constant: sum_i (offset[i] * Vals[i])
}

// NewDimCodec validates sizes and precomputes vals/bias.
func NewDimCodec(sizes DimSizes) (DimCodec, error) {
	if err := sizes.Validate(); err != nil {
		return DimCodec{}, err
	}
	return DimCodec{
		Sizes: sizes,
		Vals:  sizes.DimVals(),
		Bias:  sizes.Bias(),
	}, nil
}

// MustDimCodec panics on invalid sizes (useful during package init).
func MustDimCodec(sizes DimSizes) DimCodec {
	c, err := NewDimCodec(sizes)
	if err != nil {
		panic(err)
	}
	return c
}

// Encode converts a human-readable spec to DimInt (checked).
func (c DimCodec) Encode(spec DimensionSpec) (DimInt, error) {
	if err := c.checkRange(spec); err != nil {
		return 0, err
	}
	sum := int64(spec.Length)*c.Vals[0] +
		int64(spec.Time)*c.Vals[1] +
		int64(spec.Temperature)*c.Vals[2] +
		int64(spec.Mass)*c.Vals[3] +
		int64(spec.Current)*c.Vals[4] +
		int64(spec.Substance)*c.Vals[5] +
		int64(spec.Luminosity)*c.Vals[6] +
		int64(spec.Money)*c.Vals[7]
	return DimInt(sum), nil
}

// MustEncode panics if the spec is out of range (programmer error).
func (c DimCodec) MustEncode(spec DimensionSpec) DimInt {
	di, err := c.Encode(spec)
	if err != nil {
		panic(err)
	}
	return di
}

// Decode converts a DimInt back to a human-readable spec.
func (c DimCodec) Decode(di DimInt) DimensionSpec {
	n := int64(di) + c.Bias
	off := c.Sizes.Offsets()

	var e [8]int
	for i := 7; i >= 0; i-- {
		r := n / c.Vals[i]
		n = n % c.Vals[i]
		e[i] = int(r) - off[i]
	}
	return DimensionSpec{
		Length:      e[0],
		Time:        e[1],
		Temperature: e[2],
		Mass:        e[3],
		Current:     e[4],
		Substance:   e[5],
		Luminosity:  e[6],
		Money:       e[7],
	}
}

// checkRange ensures each exponent ∈ [-offset, offset-1].
func (c DimCodec) checkRange(s DimensionSpec) error {
	off := c.Sizes.Offsets()
	exp := [8]int{
		s.Length, s.Time, s.Temperature, s.Mass,
		s.Current, s.Substance, s.Luminosity, s.Money,
	}
	for i := 0; i < 8; i++ {
		min := -off[i]
		max := off[i] - 1
		if exp[i] < min || exp[i] > max {
			return fmt.Errorf("exponent out of range at axis %d: %d not in [%d,%d]", i, exp[i], min, max)
		}
	}
	return nil
}
