package main

import (
	"fmt"
	"gopackages-layout/internal/task_01"
	newcolor "gopackages-layout/pkg/color"
	"gopackages-layout/pkg/wordz"

	//Добавляем пакет wordz через точку
	"github.com/huandu/xstrings"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")

	fmt.Println(wordz.Hello)    //Вызов переменной из пакета wordz
	fmt.Println(wordz.Random()) //Вызов функции из пакета wordz

	fmt.Println("-----------------------------")
	fmt.Println(task_01.City())
	fmt.Printf("Try to guess this city: %s\n", xstrings.Shuffle(task_01.City()))
	fmt.Println(task_01.Digit())
}
