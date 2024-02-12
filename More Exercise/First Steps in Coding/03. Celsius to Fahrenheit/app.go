package main

import "fmt"

func main() {

	var celsius float64

	fmt.Scanln(&celsius)

	var farenheit float64 = (celsius * 9 / 5) + 32

	fmt.Printf("%.2f\n", farenheit)
}
