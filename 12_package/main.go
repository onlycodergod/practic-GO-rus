package main

import (
	"fmt"
	. "hw/module03/task_01"
	"hw/module03/wordz"

	"github.com/huandu/xstrings"
)

func main() {
	fmt.Println("Hello world")

	fmt.Println(wordz.Hello)
	fmt.Println(wordz.Random())

	fmt.Println("-----------------------------")
	fmt.Println(City())
	fmt.Printf("Try to guess this city: %s\n", xstrings.Shuffle(City()))
	fmt.Println(Digit())
}
