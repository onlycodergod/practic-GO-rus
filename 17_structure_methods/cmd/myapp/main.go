package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	cust.WrOffDebt()

	fmt.Printf("%+v\n", cust)

	price := 2000
	finalPrice, _ := internal.CalcPrice(cust, price)
	fmt.Println("finalPrice: ", finalPrice)
}
