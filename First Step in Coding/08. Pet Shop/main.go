package main

import "fmt"

func main() {
	var dogFood int
	var catFood int

	fmt.Scan(&dogFood)
	fmt.Scan(&catFood)

	var dogPrice float32 = 2.5
	catPrice := 4

	var finalDog float32 = float32(dogFood) * dogPrice
	finalCat := catFood * catPrice

	var bill float32 = float32(finalCat) + finalDog

	fmt.Printf("%v lv.", bill)
}
