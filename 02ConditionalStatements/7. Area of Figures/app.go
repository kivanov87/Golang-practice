package main

import (
	"fmt"
	"math"
)

func main() {

	var form string
	fmt.Scanln(&form)

	if form == "square" {
		var size float64
		fmt.Scanln(&size)
		var sizesquare float64 = size * size
		fmt.Printf("%0.3f\n", sizesquare)
	} else if form == "rectangle" {
		var sideA, sideB float64
		fmt.Scanln(&sideA)
		fmt.Scanln(&sideB)
		size := sideA * sideB
		fmt.Printf("%0.3f\n", size)
	} else if form == "circle" {
		var radius float64
		fmt.Scanln(&radius)
		size := math.Pi * (radius * radius)
		fmt.Printf("%.3f\n", size)
	} else if form == "triangle" {
		var a float64
		var h float64
		fmt.Scanln(&a)
		fmt.Scanln(&h)
		size := (a * h) / 2

		fmt.Printf("%0.3f\n", size)
	}
}
