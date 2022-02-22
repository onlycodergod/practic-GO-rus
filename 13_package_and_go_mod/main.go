package main

import (
	"fmt"
	newcolor "module03/color"
	"module03/task_01"
	"module03/wordz"

	"github.com/huandu/xstrings"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")

	fmt.Println(wordz.Hello)
	fmt.Println(wordz.Random())

	fmt.Println("-----------------------------")
	fmt.Println(task_01.City())
	fmt.Printf("Try to guess this city: %s\n", xstrings.Shuffle(task_01.City()))
	fmt.Println(task_01.Digit())
}
