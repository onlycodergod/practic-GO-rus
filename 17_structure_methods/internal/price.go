package internal

func CalcPrice(customer *Customer, price int) (int, error) {
	discount, err := customer.CalcDiscount()
	if err != nil {
		return 0, err
	}
	finalPrice := price - discount
	return finalPrice, err
}
