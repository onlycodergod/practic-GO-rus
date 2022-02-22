package internal

type Discounter interface {
	CalcDiscount() (int, error)
}

type Debtor interface {
	WrOffDebt() error
}
