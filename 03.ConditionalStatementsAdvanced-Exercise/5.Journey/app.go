package main

import (
	"fmt"
)

func main() {
	var budget float64
	var season string

	fmt.Scan(&budget)
	fmt.Scan(&season)

	var destination, accommodation string
	var cost float64

	if budget <= 100 {
		destination = "Bulgaria"
		if season == "summer" {
			accommodation = "Camp"
			cost = budget * 0.3
		} else {
			accommodation = "Hotel"
			cost = budget * 0.7
		}
	} else if budget <= 1000 {
		destination = "Balkans"
		if season == "summer" {
			accommodation = "Camp"
			cost = budget * 0.4
		} else {
			accommodation = "Hotel"
			cost = budget * 0.8
		}
	} else {
		destination = "Europe"
		accommodation = "Hotel"
		cost = budget * 0.9
	}

	fmt.Printf("Somewhere in %s\n", destination)
	fmt.Printf("%s - %.2f\n", accommodation, cost)
}
