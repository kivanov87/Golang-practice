package main

import (
	"fmt"
)

func main() {
	var flowerType string
	var flowerCount, budget float64

	fmt.Scan(&flowerType)
	fmt.Scan(&flowerCount)
	fmt.Scan(&budget)

	var pricePerFlower float64

	switch flowerType {
	case "Roses":
		pricePerFlower = 5
		if flowerCount > 80 {
			pricePerFlower *= 0.9
		}
	case "Dahlias":
		pricePerFlower = 3.80
		if flowerCount > 90 {
			pricePerFlower *= 0.85
		}
	case "Tulips":
		pricePerFlower = 2.80
		if flowerCount > 80 {
			pricePerFlower *= 0.85
		}
	case "Narcissus":
		pricePerFlower = 3
		if flowerCount < 120 {
			pricePerFlower *= 1.15
		}
	case "Gladiolus":
		pricePerFlower = 2.50
		if flowerCount < 80 {
			pricePerFlower *= 1.20
		}
	}

	totalPrice := flowerCount * pricePerFlower

	if budget >= totalPrice {
		fmt.Printf("Hey, you have a great garden with %.0f %s and %.2f leva left.\n", flowerCount, flowerType, budget-totalPrice)
	} else {
		fmt.Printf("Not enough money, you need %.2f leva more.\n", totalPrice-budget)
	}
}
