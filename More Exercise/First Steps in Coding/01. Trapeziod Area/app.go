package main

import "fmt"

func main() {
	var b1, b2, h float64

	fmt.Scanln(&b1)
	fmt.Scanln(&b2)
	fmt.Scanln(&h)

	var face float64 = (b1 + b2) * h / 2

	fmt.Printf("%.2f\n", face)
}
