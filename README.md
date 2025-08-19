# Go Units of Measure

[![Go Reference](https://pkg.go.dev/badge/github.com/kelledge/uom.svg)](https://pkg.go.dev/github.com/kelledge/uom)
[![CI](https://github.com/kelledge/uom/actions/workflows/ci.yml/badge.svg)](https://github.com/kelledge/uom/actions/workflows/ci.yml)

**Type-safe, expressive, and efficient handling of physical quantities in Go.**  

**NOTE: Very much still in active development. API refinements are likely.**

---

## Goals

1. **Dimensions as first-class citizens** — Compile-time and Runtime checks prevent mixing incompatible quantities (e.g., adding mass to length).  
2. **Type safety** — The compiler warns you when passing or defining invalid physical quantities.
3. **Minimal boilerplate** — Adding a new unit is quick and painless.
4. **Readable and expressive** — Expressions with physical quantities remain easy to follow, even without operator overloading.  
5. **Reasonable performance** — Dimension safety shouldn’t boil the ocean.

---

## Why This Matters

- **Catches mistakes early** — At compile time where possible, at runtime otherwise.  
- **Produces clear, informative errors** — So you know *exactly* where the mismatch is.  
- **Keeps code readable** — Code should remain simple to write and inspect for correctness.

[The Cost of Getting it Wrong](https://en.wikipedia.org/wiki/Mars_Climate_Orbiter#Cause_of_failure)

---

## Usage and Examples
Here are is a minimal example that outline the usage.

### Step 1: Defining Dimensions
We need a distinct go type and this involes implementing the uom.Dimension interface.

No one likes implementing interfaces, but fortunately you do not need many dimension types to fill out a very complete system of units.

Ships with a very reasonable set of dimensions to work from. Only fringe/esoteric uses should be expected to define their own.
```go
Time struct{}
func (Time) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Time: 1}) }
func (Time) Name() string          { return "Time" }

Length struct{}
func (Length) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Length: 1}) }
func (Length) Name() string          { return "Length" }

Area struct{}
func (Area) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Length: 2}) }
func (Area) Name() string          { return "Area" }

Volume struct{}
func (Volume) Dimension() uom.DimInt { return uom.NewDimInt(uom.DimensionSpec{Length: 3}) }
func (Volume) Name() string          { return "Volume" }
```

### Step 2: Defining Units
```go
var (
    Second = uom.DefineUnit[dim.Time]("s", 1.0)
    Hour   = uom.DefineUnit[dim.Time]("h", 3600.0)

    Meter = uom.DefineUnit[dim.Length]("m", 1.0)
    Inch  = uom.DefineUnit[dim.Length]("in", 0.0254)

    SquareMeter = uom.DefineUnit[dim.Area]("m^2", 1.0)
    CubicMeter = uom.DefineUnit[dim.Volume]("m^3", 1.0)
)
```

### Step 2.5: Deriving Units
```go
var (
    Feet = uom.DeriveUnit[dim.Length]("ft", uom.U(Inch).MulScalar(12))
    Mile = uom.DeriveUnit[dim.Length]("mile", uom.U(Feet).MulScalar(5280))

	InchesPerSecond = uom.DeriveUnit[dim.Velocity]("in/s", uom.U(Inch), uom.U(Second).Per())
    MilesPerHour    = uom.DeriveUnit[dim.Velocity]("mph", uom.U(Mile), uom.U(Hour).Per())

    CubicInch = uom.DeriveUnit[dim.Volume]("in^3", uom.U(Inch).Pow(3))
    Gallon    = uom.DeriveUnit[dim.Volume]("gallon", uom.U(CubicInch).MulScalar(231))
)
```

### Step 3: Using Quantities and Expressions
```go
type Cylinder struct {
    Width uom.Quantity[dim.Length]
    Height uom.Quantity[dim.Length]
}

func (c Cylinder) Radius() uom.Quantity[dim.Length] {
    // d / 2
    return uom.E[dim.Length](c.Width).Div(uom.Scalar(2)).As(Meter)
}

func (c Cylinder) CrossSection() uom.Quantity[dim.Area] {
    // r^2 * pi
    return uom.E[dim.Area](c.Radius()).Pow(2).Mul(uom.Scalar(math.Pi)).As(SquareMeter)
}

func (c Cylinder) Volume() uom.Quantity[dim.Volume] {
    // r^2 * pi * h
    return uom.E[dim.Volume](c.Height).Mul(c.CrossSection()).As(CubicMeter)
}

func (c Cylinder) VolumeInSingleCall() uom.Quantity[dim.Volume] {
    return uom.E[dim.Volume](c.Width).
        Div(uom.Scalar(2)).
        Pow(2).
        Mul(uom.Scalar(math.Pi)).
        Mul(c.Height).
        As(CubicMeter)
}

func main() {
    c := Cylinder{
        Width: uom.Q(6, Inch),
        Height: uom.Q(1, Meter)
    }

    fmt.Printf("Volume: %.2f gallons\n", uom.As(c.Volume(), Gallon))
    fmt.Printf("Volume: %.2f m^3\n", uom.As(c.Volume(), CubicMeter))
    fmt.Printf("Volume: %.2f in^3\n", uom.As(c.Volume(), CubicInch))
}
```

---

### Step 4: Fix Dimension Errors
This library takes the stance that a dimension error is not a valid program. I.E. dimension analysis error is a compile error or runtime panic.

No amount of error handling can recover adding 60 seconds to 5 miles. The author simply made a mistake and will need to be more careful in their implementation.

We're all human and this will happen. The outcome to measure is how quickly the issue can be identified and corrected.

#### Compile Errors
Best case scenario: the compiler tells you exactly where your error is:
```go
func main() {
    c := Cylinder{
        Width: uom.Q(6, Gallon), // Compile-time error; Can not use type uom.Quantity[dim.Volume] as type uom.Quantity[dim.Length]
        Height: uom.Q(1, Meter)
    }
}
```

#### Import Runtime Errors
Next best case scenario: there is a panic a package import time. 
```go
var (
    CubicInch = uom.DeriveUnit[dim.Volume]("in^3", uom.U(Inch).Pow(3), uom.U(Second).Per()) // Import-time panic. Target dimension L^3; Term dimension L^4*T^-1
)
```

#### Evaluation Runtime Errors
This is the worst case scenario: you asked an expression to evaluate something that does not make physical sense and you only find out about this when the specific unit of code is called.

This will panic with a message describing the illegal operation as well as calling stack when the unit was called.
```go
func GetVolume(l, w, h uom.Quantity[dim.Length]) uom.Quantity[dim.Volume] {
    return uom.One[dim.Volume]().Mul(l, w, h, h).As(CubicMeter) // Runtime panic when called. Target dimension L^3; Expression dimension L^4
}
```

## How Dimensions are Handled
Crucially, dimensions are tied to Go types for compile-time enforcment.

At runtime the dimension vectors are encoded to simple int64 values which means comparison and operations generally happen in a single instruction.

See: https://www.cs.utexas.edu/~novak/units95.html and related works.

---

## Next:

  1. Affine Conversions: Accomdates Celsius and Fahrenheit
  2. Serialization: It should be trivial to add MarshalText/UnmarshalText to uom.Quantity[T Dimension]
  3. Aliases: "foot", "feet", "ft"
  4. Registry:

---

## Horizon
Work that is further off, but has promise to be extremely useful.

### Leverage the Existing Work of [QUDT](https://www.qudt.org/pages/HomePage.html)

Write a go:generate tool that converts the library of maintained dimension vectors and units directly to the `dim`/`si`/`usc` packages.

There is already a generally accept body that maintains this information. Use it.


---

## Headwinds in Go
Go's strength is its simplicity. But this also means that any library will have to make some compromises in its implementation.

1. **No operator overloading**  
    Expressions like `distance / time` must call functions instead of leaning on the compiler for evaluation. Expressions can still be quite readable, but will never be as readable as they could be in c++ or rust for example.

2. **Minimal type system**  
    Go’s type system is intentionally lean. In practice this puts upper limits on what can be validated at compile time and forces more runtime checks.

3. **No `constexpr` or useful `const`**  
    Without compile-time evaluation of constants, defining large unit libraries means the many conversion factors are computed at runtime.