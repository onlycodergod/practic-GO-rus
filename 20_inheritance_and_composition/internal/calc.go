package internal

import "errors"

func CalcPrice(discounter Discounter, price int) (int, error) {
	discount, err := discounter.CalcDiscount()
	if err != nil {
		return 0, err
	}
	finalPrice := price - discount
	return finalPrice, err
}

func (c *Customer) CalcDiscount() (int, error) {
	if !c.discount {
		return 0, errors.New("discount not available")
	}
	result := DEFAULT_DISCOUNT - c.debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}
