package si

import "fmt"

func ExampleKilogram() {
	fmt.Println(1*Kilogram == 1)
	// Output: true
}

func ExampleGram() {
	fmt.Println(1000*Gram == 1*Kilogram)
	// Output: true
}

func ExampleMilligram() {
	fmt.Println(1000*Milligram == 1*Gram)
	// Output: true
}
