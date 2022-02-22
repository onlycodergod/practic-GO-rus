package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "104"
	num := 35

	// Convert
	strToInt, _ := strconv.Atoi(str)
	numToStr := strconv.Itoa(num)

	// Check value and type
	fmt.Print(strToInt)
	fmt.Printf(", type: %T\n", strToInt)
	fmt.Print(numToStr)
	fmt.Printf(", type: %T\n", numToStr)
}
