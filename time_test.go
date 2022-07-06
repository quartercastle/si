package si

import (
	"fmt"
	"math"
	"time"
)

func ExampleSecond() {
	fmt.Println(math.Abs(((1*Second)/Nanosecond)-float64(time.Second)) <= 1e-6)
	// Output: true
}

func ExampleMinute() {
	fmt.Println(1*Minute == (1 * time.Minute).Seconds())
	// Output: true
}

func ExampleHour() {
	fmt.Println(1*Hour == (1 * time.Hour).Seconds())
	// Output: true
}

func ExampleDay() {
	fmt.Println(1*Day == (24 * time.Hour).Seconds())
	// Output: true
}

func ExampleWeek() {
	fmt.Println(1*Week == (7 * 24 * time.Hour).Seconds())
	// Output: true
}
