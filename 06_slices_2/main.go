package main

import "fmt"

func main() {
	workDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	weekends := []string{"Saturday", "Sunday"}

	// Create new slice with all days
	weekDays := append(workDays, weekends...)

	// Display
	fmt.Println(weekDays)
}
