package main

import "fmt"

func main() {
	var typeProjection string
	var rows, columns int
	var price float64
	fmt.Scanln(&typeProjection)
	fmt.Scanln(&rows)
	fmt.Scanln(&columns)

	const premier = 12.00
	const normal = 7.50
	const discount = 5.00
	seats := rows * columns
	switch typeProjection {
	case "Premiere":
		price = float64(seats) * float64(premier)
	case "Normal":
		price = float64(seats) * float64(normal)
	case "Discount":
		price = float64(seats) * float64(discount)
	}

	fmt.Printf("%.2f leva", price)
}
