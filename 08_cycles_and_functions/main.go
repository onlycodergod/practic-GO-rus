package main

import "fmt"

func main() {
	a := []string{"one", "two", "three"}
	x := "two"
	fmt.Println(contains(x, a...))

	fmt.Println(getMax(80, 1, 2, 3, 4, 5, -99, 0))
}

// Find x in a
func contains(x string, a ...string) (flag bool) {
	for _, v := range a {
		if v == x {
			flag = true
		}
	}
	return
}

// Find max in nums
func getMax(nums ...int) (max int) {
	max = nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return
}
