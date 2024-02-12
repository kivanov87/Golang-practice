package main

import "fmt"

func main() {
	var a, h float64

	fmt.Scanln(&a)
	fmt.Scanln(&h)

	area := a * h / 2
	fmt.Printf("%.2f\n", area)
}
