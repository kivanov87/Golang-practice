package main

import "fmt"

func main() {
	var typeOfFlowers string
	var countOfFlowers int
	var budget int
	var price, diff float64
	const Roses float64 = 5
	const dahlia = 3.80
	const tulip = 2.80
	const narcis float64 = 3
	const gladiola = 2.50

	fmt.Scan(&typeOfFlowers, &countOfFlowers, &budget)
	switch typeOfFlowers {
	case "Roses": ////////////////////////////
		if countOfFlowers > 80 {
			price = (float64(countOfFlowers) * Roses) * 0.9
		} else {
			price = float64(countOfFlowers) * Roses
		}
	case "Dahlias": ////////////////////
		if countOfFlowers > 90 {
			price = (float64(countOfFlowers) * dahlia) * 0.85
		} else {
			price = float64(countOfFlowers) * dahlia
		}
	case "Tulips":
		if countOfFlowers > 80 {
			price = (float64(countOfFlowers) * tulip) * 0.85
		} else {
			price = float64(countOfFlowers) * tulip
		}
	case "Narcissus":
		if countOfFlowers < 120 {
			price = (float64(countOfFlowers) * narcis) * 1.15
		} else {
			price = float64(countOfFlowers) * narcis
		}
	case "Gladiolus":
		if countOfFlowers < 80 {
			price = (float64(countOfFlowers) * gladiola) * 1.20
		} else {
			price = float64(countOfFlowers) * gladiola
		}
	}
	if float64(budget) > price {
		diff = float64(budget) - price
		fmt.Printf("Hey, you have a great garden with %v %v and %.2f leva left.", countOfFlowers, typeOfFlowers, diff)
	} else {
		diff = price - float64(budget)
		fmt.Printf("Not enough money, you need %.2f leva more.", diff)
	}
}
