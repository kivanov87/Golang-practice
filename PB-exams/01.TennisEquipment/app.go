package main

import (
	"fmt"
	"math"
)

func main() {

	var tennisRocketPrice float64
	var tennisRocketCount int
	var shoesCount int

	fmt.Scanln(&tennisRocketPrice)
	fmt.Scanln(&tennisRocketCount)
	fmt.Scanln(&shoesCount)

	shoesPrice := tennisRocketPrice / 6

	rocketTotal := float64(tennisRocketCount) * float64(tennisRocketPrice)

	shoesTotal := float64(shoesCount) * float64(shoesPrice)

	total := rocketTotal + shoesTotal

	equipment := float64(total) * 0.2

	finalPrice := float64(total) + float64(equipment)

	djokovicPrice := math.Floor(float64(finalPrice) / 8.0)
	sponsorsPrice := math.Ceil(7.0 * float64(finalPrice) / 8.0)

	fmt.Printf("Price to be paid by Djokovic %v\n", djokovicPrice)
	fmt.Printf("Price to be paid by sponsors %v\n", sponsorsPrice)

}
