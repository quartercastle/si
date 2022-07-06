package si

const (
	Kelvin      Temperature = 1e0
	Yoctokelvin             = Yocto * Kelvin
	Zeptokelvin             = Zepto * Kelvin
	Attokelvin              = Atto * Kelvin
	Femtokelvin             = Femto * Kelvin
	Picokelvin              = Pico * Kelvin
	Nanokelvin              = Nano * Kelvin
	Microkelvin             = Micro * Kelvin
	Millikelvin             = Milli * Kelvin
	Centikelvin             = Centi * Kelvin
	Decikelvin              = Deci * Kelvin
	Decakelvin              = Deca * Kelvin
	Hectokelvin             = Hecto * Kelvin
	Kilokelvin              = Kilo * Kelvin
	Megakelvin              = Mega * Kelvin
	Gigakelvin              = Giga * Kelvin
	Terakelvin              = Tera * Kelvin
	Petakelvin              = Peta * Kelvin
	Exakelvin               = Exa * Kelvin
	Zettakelvin             = Zetta * Kelvin
	Yottakelvin             = Yotta * Kelvin
)

func FromCelsius(t float64) Temperature {
	return t + 273.15
}

func FromDelisle(t float64) Temperature {
	return 373.15 - (t * 2 / 3)
}

func FromFahrenheit(t float64) Temperature {
	return (t + 459.67) * 5 / 9
}

func FromNewton(t float64) Temperature {
	return t*100/33 + 273.15
}

func FromRankine(t float64) Temperature {
	return (t-491.67)*5/9 + 273.15
}

func FromReaumur(t float64) Temperature {
	return t*5/4 + 273.15
}

func FromRomer(t float64) Temperature {
	return (t-7.5)*40/21 + 273.15
}

func ToCelsius(t Temperature) float64 {
	return t - 273.15
}

func ToDelisle(t Temperature) float64 {
	return (373.15 - t) * 3 / 2
}

func ToFahrenheit(t Temperature) float64 {
	return (t * 9 / 5) - 459.67
}

func ToNewton(t Temperature) float64 {
	return (t - 273.15) * 33 / 100
}

func ToRankine(t Temperature) float64 {
	return (t-273.15)*9/5 + 491.67
}

func ToReaumur(t Temperature) float64 {
	return (t - 273.15) * 4 / 5
}

func ToRomer(t Temperature) float64 {
	return (t-273.15)*21/40 + 7.5
}
