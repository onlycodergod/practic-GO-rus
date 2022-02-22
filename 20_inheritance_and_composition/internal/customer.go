package internal

import "errors"

const DEFAULT_DISCOUNT = 500

type Customer struct {
	Overduer
	Name     string
	Age      int
	discount bool
}

type Overduer struct {
	balance int
	debt    int
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Overduer: Overduer{
			balance: balance,
			debt:    debt,
		},
		Name:     name,
		Age:      age,
		discount: discount,
	}
}

func (c *Overduer) WrOffDebt() error {
	if c.debt >= c.balance {
		return errors.New("not possible write off")
	}

	c.balance -= c.debt
	c.debt = 0

	return nil
}
