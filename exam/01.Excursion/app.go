package main

import "fmt"

func main() {
	const pricePerNight int = 20
	const transeportCard float64 = 1.60
	const priceOfTickets int = 6

	var countOfPeople, countOfNight, countOfCards, countOfTickets int

	fmt.Scanln(&countOfPeople)
	fmt.Scanln(&countOfNight)
	fmt.Scanln(&countOfCards)
	fmt.Scanln(&countOfTickets)

	costPerNight := countOfNight * pricePerNight
	var costForTransport float64 = float64(countOfCards) * float64(transeportCard)
	costForMuseum := countOfTickets * priceOfTickets

	var costPerPerson float64 = float64(costPerNight) + costForTransport + float64(costForMuseum)

	wholeGroup := float64(countOfPeople) * float64(costPerPerson)

	total := float64(wholeGroup) * 1.25

	fmt.Printf("%.2f", total)
}
