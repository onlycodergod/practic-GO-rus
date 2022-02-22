package main

import (
	"fmt"
	"time"
)

func main() {
	x := time.Now().Format("02.01.2006 15:04")
	fmt.Println("Welcome!")
	fmt.Println(x)
}
