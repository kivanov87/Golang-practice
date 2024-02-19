package main

import "fmt"

func main() {
	var budget, clothesPrice float64
	var extras int
	fmt.Scan(&budget)
	fmt.Scan(&extras)
	fmt.Scan(&clothesPrice)

	decor := budget * 0.1
	var clothes float64
	if extras > 150 {
		clothes = 0.9 * clothesPrice * float64(extras)
	} else {
		clothes = clothesPrice * float64(extras)
	}
	total := decor + clothes
	if total > budget {
		fmt.Println("Not enough money!")
		fmt.Printf("Wingard needs %.2f leva more.\n", total-budget)
	} else {
		fmt.Println("Action!")
		fmt.Printf("Wingard starts filming with %.2f leva left.\n", budget-total)
	}

}
