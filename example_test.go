package si_test

import (
	"fmt"
	"math"

	"github.com/quartercastle/si"
)

func Example() {
	const (
		mg  = si.Milligram
		MWh = si.MegawattHour
		c   = si.SpeedOfLight
	)

	energyIsEqualToMass := func(energy si.Energy, mass si.Mass) bool {
		return energy == mass*math.Pow(c, 2)
	}

	fmt.Println(
		energyIsEqualToMass(24.965421631578266*MWh, 1*mg),
	)
	// Output: true
}
