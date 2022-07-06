package si

import (
	"fmt"
	"math"
)

func ExampleRadian() {
	fmt.Println(2*math.Pi == 6.283185307179586*Radian)
	// Output: true
}

func ExampleDegree() {
	fmt.Println(math.Pi == 180*Degree)
	// Output: true
}

func ExampleArcminute() {
	fmt.Println(21600*Arcminute == 360*Degree)
	// Output: true
}

func ExampleArcsecond() {
	fmt.Println(1.296e+6*Arcsecond == 360*Degree)
	// Output: true
}

func ExampleMilliarcsecond() {
	fmt.Println(math.Abs(((1.296e+6/Milli)*Milliarcsecond)-(360*Degree)) <= 1e-15)
	// Output: true
}

func ExampleMicroarcsecond() {
	fmt.Println(math.Abs(((1.296e+6/Micro)*Microarcsecond)-(360*Degree)) <= 1e-15)
	// Output: true
}
