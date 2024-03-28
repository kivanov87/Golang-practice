package main

import "fmt"

func main() {
	var priceOfShirt, goalSum float64

	fmt.Scanln(&priceOfShirt)
	fmt.Scanln(&goalSum)

	priceOfShort := float64(priceOfShirt) * 0.75
	priceOfSocks := float64(priceOfShort) * 0.20
	priceOfButtons := (priceOfShirt + priceOfShort) * 2
	var totalSum float64 = priceOfShirt + priceOfShort + priceOfSocks + priceOfButtons

	totalSumDiscount := totalSum * 0.85

	if totalSumDiscount >= goalSum {
		fmt.Println("Yes, he will earn the world-cup replica ball!")
		fmt.Printf("His sum is %.2f lv.", totalSumDiscount)
	} else {
		neededMoney := goalSum - totalSumDiscount

		fmt.Println("No, he will not earn the world-cup replica ball.")
		fmt.Printf("He needs %.2f lv. more.", neededMoney)
	}
}
