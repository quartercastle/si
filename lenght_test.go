package si

import "fmt"

func ExampleMillimeter() {
	fmt.Println(10*Millimeter == 1*Centimeter)
	// Output: true
}

func ExampleCentimeter() {
	fmt.Println(100*Centimeter == 1*Meter)
	// Output: true
}

func ExampleKilometer() {
	fmt.Println(1000*Meter == 1*Kilometer)
	// Output: true
}
