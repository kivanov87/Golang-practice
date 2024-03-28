package main

import "fmt"

func main() {
	var countOfPeople int
	var season string
	var price, total float64

	fmt.Scanln(&countOfPeople)
	fmt.Scanln(&season)

	switch season {
	case "spring":
		if countOfPeople <= 5 {
			price = 50.00
			total = float64(countOfPeople) * price
		} else {
			price = 48.00
			total = float64(countOfPeople) * price
		}
	case "summer":
		if countOfPeople <= 5 {
			price = 48.50 * 0.85
			total = float64(countOfPeople) * price
		} else {
			price = 45.00 * 0.85
			total = float64(countOfPeople) * price
		}
	case "autumn":
		if countOfPeople <= 5 {
			price = 60.00
			total = float64(countOfPeople) * price
		} else {
			price = 49.50
			total = float64(countOfPeople) * price
		}
	case "winter":
		if countOfPeople <= 5 {
			price = 86.00 * 1.08
			total = float64(countOfPeople) * price
		} else {
			price = 85.00 * 1.08
			total = float64(countOfPeople) * price
		}
	}
	fmt.Printf("%.2f leva.", total)
}
