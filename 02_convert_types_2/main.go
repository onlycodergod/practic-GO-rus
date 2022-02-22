package main

import (
	"fmt"
	"math"
)

func main() {
	type AmericanVelocity float64
	type EuropeanVelocity float64

	// Convert
	speed1 := 120.4                                                  // m/s
	speed1Conv := AmericanVelocity(math.Round(speed1*3.6*100) / 100) // km/h

	speed2 := float64(130)                                             // m/s
	speed2Conv := EuropeanVelocity(math.Round(speed2*2.237*100) / 100) // mile/h

	// Check value and type
	fmt.Println(speed1Conv)
	fmt.Printf("type: %T\n", speed1Conv)
	fmt.Println(speed2Conv)
	fmt.Printf("type: %T\n", speed2Conv)
}

