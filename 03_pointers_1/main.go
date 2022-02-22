package main

import "fmt"

func main() {
	var A *int
	var B int = 42

	A = &B
	fmt.Println(*A) // 42

	*A = 24
	fmt.Println(B) // 24
}
