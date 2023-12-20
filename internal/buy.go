package internal

func CalcPrice(customer Customer, price int) (int, error) {
	if newDiscount, err := customer.CalcDiscount(); err != nil {
		return 0, err
	} else {
		return price - newDiscount, nil
	}
}
