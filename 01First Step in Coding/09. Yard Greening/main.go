package main

import "fmt"

func main() {
	var priceSquare float32 = 7.61
	var square float32

	fmt.Scan(&square)

	var price float32 = square * priceSquare
	var discount float32 = 0.18 * price
	var total float32 = price - discount
	fmt.Printf("The final price is: %v lv.\n", total)
	fmt.Printf("The discount is: %v lv.", discount)

}
