package main

import (
	"fmt"
	"math"
)

func main() {
	// Calculate the radius of a circle
	len := 35
	r1 := math.Round(float64(len)/(2*math.Pi)*100) / 100
	fmt.Println(r1) // 5.57

	// Calculate the square of a circle
	var R *float64
	R = &r1
	S := math.Round(math.Pi*(math.Pow(*R, 2))*100) / 100
	fmt.Println(S) // 97.47
}

