package main

import "fmt"

func main() {
	var age, toyPrice, toyCounter, lilyMoney int
	var washingMachinePrice float64
	var birthdayMoney = 10

	fmt.Scanln(&age)
	fmt.Scanln(&washingMachinePrice)
	fmt.Scanln(&toyPrice)

	for i := 1; i <= age; i++ {

		if i%2 != 0 { //нечетно число
			toyCounter++
		} else { //четно число
			lilyMoney += birthdayMoney
			lilyMoney--

			birthdayMoney += 10
		}
	}

	lilyMoney += toyCounter * toyPrice

	if float64(lilyMoney) >= washingMachinePrice {
		diff := float64(lilyMoney) - washingMachinePrice

		fmt.Printf("Yes! %.2f", diff)
	} else {
		neededMoney := washingMachinePrice - float64(lilyMoney)

		fmt.Printf("No! %.2f", neededMoney)
	}

}
