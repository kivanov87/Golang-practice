package main

import (
	"fmt"
)

func main() {
	var budget float64
	var nVideoCards, nProcessors, nRam int
	fmt.Scan(&budget, &nVideoCards, &nProcessors, &nRam)

	videoCardPrice := 250.0
	processorPrice := 0.35 * videoCardPrice * float64(nVideoCards)
	ramPrice := 0.10 * videoCardPrice * float64(nVideoCards)

	totalPrice := float64(nVideoCards)*videoCardPrice + float64(nProcessors)*processorPrice + float64(nRam)*ramPrice

	if nVideoCards > nProcessors {
		totalPrice *= 0.85
	}

	diff := budget - totalPrice

	if diff >= 0 {
		fmt.Printf("You have %.2f leva left!\n", diff)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva more!\n", -diff)
	}
}
