package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	startTransactoinDynamic(cust)

	fmt.Printf("%+v\n", cust)

	price := 2000
	finalPrice, _ := internal.CalcPrice(cust, price)
	fmt.Println("finalPrice: ", finalPrice)
}

func startTransactoinDynamic(p interface{}) error {
	debtor, ok := p.(internal.Debtor)
	if !ok {
		return errors.New("incorrect type")
	}
	return debtor.WrOffDebt()
}
