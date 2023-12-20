package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

const DEFAULT_DISCOUNT = 5000
const DEFAULT_PRICE = 10000

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	fmt.Printf("%+v\n", cust)

	finalPrice, err := internal.CalcPrice(*cust, DEFAULT_PRICE)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The final price of a product worth %d is: %d\n", DEFAULT_PRICE, finalPrice)
}
