package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	// Standard SI unit of energy; used for heat, work, and internal energy in thermodynamic systems
	Joule = uom.DefineUnit[dim.Energy]("J", 1)

	// SI unit of power; measures rate of energy transfer (e.g., heating or cooling power)
	Watt = uom.DefineUnit[dim.Power]("W", 1)

	// Pressure is central to state equations and phase-change analysis in thermodynamics
	Pascal = uom.DefineUnit[dim.Pressure]("Pa", 1)

	// Thermal conductivity; quantifies ability of a material to conduct heat
	WattPerMeterK = uom.DefineUnit[dim.ThermalConductivity]("W/(m·K)", 1)

	// Heat capacity; total energy required to change an object's temperature by 1 K
	JoulePerKelvin = uom.DefineUnit[dim.HeatCapacity]("J/K", 1)

	// Specific heat capacity; heat capacity per unit mass, key in calorimetry and energy balance
	JoulePerKilogramK = uom.DefineUnit[dim.SpecificHeatCapacity]("J/(kg·K)", 1)

	// Heat transfer coefficient; overall heat transfer rate across a surface per unit area and temperature difference
	WPerM2K = uom.DefineUnit[dim.HeatTransferCoefficient]("W/(m²·K)", 1)
)
