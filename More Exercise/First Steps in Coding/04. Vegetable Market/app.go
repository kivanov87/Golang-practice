package main

import "fmt"

func main() {
	const eur = 1.94

	var priceVeg, priceFruits float64
	var vegQuant, fruitsQuant int

	fmt.Scanln(&priceVeg)
	fmt.Scanln(&priceFruits)
	fmt.Scanln(&vegQuant)
	fmt.Scanln(&fruitsQuant)

	var vegTotal float64 = float64(vegQuant) * priceVeg
	var fruitTotal float64 = float64(fruitsQuant) * priceFruits
	var finalPrice float64 = vegTotal + fruitTotal

	var eurPrice float64 = finalPrice / eur

	fmt.Printf("%.2f\n", eurPrice)

}
