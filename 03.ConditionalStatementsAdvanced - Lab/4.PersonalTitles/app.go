package main

import "fmt"

func main() {
	var age float64
	var gender string

	fmt.Scanln(&age)
	fmt.Scanln(&gender)
	switch gender {
	case "m":
		if age >= 16 {
			fmt.Println("Mr.")
		} else if age < 16 {
			fmt.Println("Master")
		}

	case "f":
		if age >= 16 {
			fmt.Println("Ms.")
		} else if age < 16 {
			fmt.Println("Miss")
		}
	}
}
