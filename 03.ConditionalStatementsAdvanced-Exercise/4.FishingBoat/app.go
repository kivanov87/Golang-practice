package main

import (
	"fmt"
)

func main() {
	var budget float64
	var season string
	var fishermen int

	fmt.Scan(&budget)
	fmt.Scan(&season)
	fmt.Scan(&fishermen)

	var boatRent float64

	switch season {
	case "Spring":
		boatRent = 3000
	case "Summer", "Autumn":
		boatRent = 4200
	case "Winter":
		boatRent = 2600
	}

	if fishermen <= 6 {
		boatRent *= 0.9
	} else if fishermen <= 11 {
		boatRent *= 0.85
	} else {
		boatRent *= 0.75
	}

	if fishermen%2 == 0 && season != "Autumn" {
		boatRent *= 0.95
	}

	if budget >= boatRent {
		fmt.Printf("Yes! You have %.2f leva left.\n", budget-boatRent)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva.\n", boatRent-budget)
	}
}
