package main

import "fmt"

func main() {

	const puzzlePrice float64 = 2.60
	const doolPrice float64 = 3
	const bearPrice float64 = 4.10
	const minionPrice float64 = 8.20
	const truckPrice float64 = 2

	var excPrice float64
	var countPuzzle, countDools, countBears, countMinions, countTrucks int
	fmt.Scanln(&excPrice)
	fmt.Scanln(&countPuzzle)
	fmt.Scanln(&countDools)
	fmt.Scanln(&countBears)
	fmt.Scanln(&countMinions)
	fmt.Scanln(&countTrucks)

	countOfStocks := countPuzzle + countDools + countBears + countMinions + countTrucks

	var fullPrice float64 = float64(countPuzzle)*float64(puzzlePrice) + float64(countDools)*float64(doolPrice) + float64(countBears)*bearPrice + float64(countMinions)*minionPrice + float64(countTrucks)*truckPrice
	var discount float64
	if countOfStocks >= 50 {
		discount = fullPrice * 0.25
		fullPrice = fullPrice - discount
	}
	rent := fullPrice * 0.1
	revenue := fullPrice - rent //rent is here

	if revenue > excPrice {
		moneyLeft := revenue - excPrice
		fmt.Printf("Yes! %.2f lv left.", moneyLeft)

	} else if revenue <= excPrice {
		moneyNeeded := excPrice - revenue
		fmt.Printf("Not enough money! %.2f lv needed.", moneyNeeded)
	}

}
