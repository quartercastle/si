package si

func DataPretty(value Data) (float64, string) {
	allowed := map[float64]int{
		Peta: 0,
		Tera: 0,
		Giga: 0,
		Mega: 0,
		Kilo: 0,
		Byte: 0,
	}
	for i, prefix := range prefixes {
		if _, ok := allowed[factors[i]]; !ok {
			continue
		}
		if value/Byte < 100 {
			return value / Byte, "B"
		}
		if value/factors[i]/Byte < 1000 {
			return value / factors[i] / Byte, prefix + "B"
		}
	}
	return 0, ""
}

func contains(factors []float64, factor float64) bool {
	for _, f := range factors {
		if f == factor {
			return true
		}
	}
	return false
}

type Base func() (float64, string)

func MeterBase() (float64, string) {
	return Meter, "m"
}

func ByteBase() (float64, string) {
	return Byte, "B"
}

func Readable(value Value, fn Base, allowed ...float64) (float64, string) {
	var (
		result     float64 = 0
		base, unit         = fn()
	)

	for i, prefix := range prefixes {
		factor := factors[i]

		if !contains(allowed, factor) {
			continue
		}

		if value/factor < 1000 {
			return value / factor / base, prefix + unit
		}

		result = value / factor / base
	}

	return result, unit
}
