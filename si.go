package si

// Prefixes
const (
	Yocto = 1e-24
	Zepto = 1e-21
	Atto  = 1e-18
	Femto = 1e-15
	Pico  = 1e-12
	Nano  = 1e-9
	Micro = 1e-6
	Milli = 1e-3
	Centi = 1e-2
	Deci  = 1e-1
	Deca  = 1e1
	Hecto = 1e2
	Kilo  = 1e3
	Mega  = 1e6
	Giga  = 1e9
	Tera  = 1e12
	Peta  = 1e15
	Exa   = 1e18
	Zetta = 1e21
	Yotta = 1e24
)

// SI
type Length = float64
type Distance = Length
type Time = float64
type Duration = Time
type Mass = float64
type Weight = Mass
type AmountOfSubstance = float64
type ElectricCurrent = float64
type LuminousIntensity = float64
type Temperature = float64

// Derived SI
type Acceleration = float64
type Speed = float64
type Velocity = Speed
type Frequency = float64
type Force = float64
type Pressure = float64
type Energy = float64
type Power = float64
type ElectricCharge = float64
type ElectricPotential = float64
type ElectricResistance = float64
type Angle = float64
type Area = float64
type Volume = float64
type Data = float64
type Datarate = float64
type Datasize = float64
type Luminance = float64
type Illuminance = float64
type MassDensity = float64
