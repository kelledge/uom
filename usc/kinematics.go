package usc

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/std"
)

// Velocity (linear)
var (
	// Fine-scale motion, actuators, instrumentation
	InchesPerSecond = uom.DeriveUnit[dim.Velocity]("in/s",
		uom.U(Inch),
		uom.U(std.Second).Pow(-1),
	)

	// General-purpose speed in labs, ballistics, and controls
	FootPerSecond = uom.DeriveUnit[dim.Velocity]("ft/s",
		uom.U(Foot),
		uom.U(std.Second).Pow(-1),
	)

	// Sports timing, short-range tracking; sometimes range tables
	YardsPerSecond = uom.DeriveUnit[dim.Velocity]("yd/s",
		uom.U(Yard),
		uom.U(std.Second).Pow(-1),
	)

	// Everyday & vehicle speeds; logging, traffic, and drivetrain calcs
	MilesPerHour = uom.DeriveUnit[dim.Velocity]("mi/h",
		uom.U(Mile),
		uom.U(std.Hour).Pow(-1),
	)

	// Fans, conveyors, lifts; convenient for facility specs
	FootPerMinute = uom.DeriveUnit[dim.Velocity]("ft/min",
		uom.U(Foot),
		uom.U(std.Minute).Pow(-1),
	)

	// Feed rates, small mechanisms, machining
	InchesPerMinute = uom.DeriveUnit[dim.Velocity]("in/min",
		uom.U(Inch),
		uom.U(std.Minute).Pow(-1),
	)
)

// Acceleration (linear)
var (
	// Bench tests, drop tests, small mechanisms
	InchesPerSecondSquared = uom.DeriveUnit[dim.Acceleration]("in/s^2",
		uom.U(Inch),
		uom.U(std.Second).Pow(-2),
	)

	// Canonical USC acceleration (already used to define standard gravity)
	FootPerSecondSquared = uom.DeriveUnit[dim.Acceleration]("ft/s^2",
		uom.U(Foot),
		uom.U(std.Second).Pow(-2),
	)
)
