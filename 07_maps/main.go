package main

import "fmt"

func main() {
	items := map[string]map[string][]string{}

	// Add some values
	items["Ivan"] = map[string][]string{
		"Books":     []string{"Cisco", "Patterns", "Algorithms"},
		"Magazines": []string{"Play Boy"},
	}

	items["Semen"] = map[string][]string{
		"Books":     []string{"Chemistry"},
		"Magazines": []string{"Young naturalist"},
	}

	// Show number of users
	fmt.Println("Number of users:", len(items))

	// Show how many books and how many magazines user has
	for u := range items {
		fmt.Printf("%v: %v books, %v magazines\n", u, len(items[u]["Books"]), len(items[u]["Magazines"]))
	}
}

