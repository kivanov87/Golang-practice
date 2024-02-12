package main

import "fmt"

func main() {
	const midiPrice = 7.50
	var skumriaPrice, cacaPrice, palamudWeight, safridWeight float64
	var midiWeight int

	fmt.Scanln(&skumriaPrice)
	fmt.Scanln(&cacaPrice)
	fmt.Scanln(&palamudWeight)
	fmt.Scanln(&safridWeight)
	fmt.Scanln(&midiWeight)

	palamudPrice := skumriaPrice * 1.6
	safridPrice := cacaPrice * 1.8

	palamudFinal := palamudWeight * palamudPrice
	safridFinal := safridWeight * safridPrice
	midiFinal := float64(midiWeight) * midiPrice

	totalBill := palamudFinal + safridFinal + midiFinal

	fmt.Printf("%.2f\n", totalBill)
}
