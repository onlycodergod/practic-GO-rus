package internal

type Discounter interface {
	CalcDiscount() (int, error)
}
