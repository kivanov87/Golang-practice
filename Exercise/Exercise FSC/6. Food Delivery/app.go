package main

import "fmt"

func main() {
	const priceChicken = 10.35
	const priceFish = 12.40
	const priceVegetarian = 8.15
	const delivery = 2.50

	var orderChicken int
	var orderFish int
	var orderVegetarian int

	fmt.Scanln(&orderChicken)
	fmt.Scanln(&orderFish)
	fmt.Scanln(&orderVegetarian)

	var chickenOrder float32 = float32(orderChicken) * priceChicken
	var fishOrder float32 = float32(orderFish) * priceFish
	var vegetarianOrder float32 = float32(orderVegetarian) * priceVegetarian
	priceForOrder := chickenOrder + fishOrder + vegetarianOrder
	dessert := priceForOrder * 0.20

	total := priceForOrder + dessert + delivery

	fmt.Println(total)

}
