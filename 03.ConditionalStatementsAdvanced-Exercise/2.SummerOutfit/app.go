package main

import "fmt"

func main() {
	var degreese int
	var timeofDay, outfit, shoes string
	fmt.Scanln(&degreese)
	fmt.Scanln(&timeofDay)

	switch timeofDay {
	case "Morning":
		if degreese >= 10 && degreese <= 18 {
			outfit = "Sweatshirt"
			shoes = "Sneakers"
		} else if degreese > 18 && degreese <= 24 {
			outfit = "Shirt"
			shoes = "Moccasins"
		} else if degreese >= 25 {
			outfit = "T-Shirt"
			shoes = "Sandals"
		}
	case "Afternoon":
		if degreese >= 10 && degreese <= 18 {
			outfit = "Shirt"
			shoes = "Moccasins"
		} else if degreese > 18 && degreese <= 24 {
			outfit = "T-Shirt"
			shoes = "Sandals"
		} else if degreese >= 25 {
			outfit = "Swim Suit"
			shoes = "Barefoot"
		}
	case "Evening":
		if degreese >= 10 && degreese <= 18 {
			outfit = "Sweatshirt"
			shoes = "Sneakers"
		} else if degreese > 18 && degreese <= 24 {
			outfit = "Shirt"
			shoes = "Moccasins"
		} else if degreese >= 25 {
			outfit = "Shirt"
			shoes = "Moccasins"
		}
	}
	fmt.Printf("It's %v degrees, get your %v and %v.", degreese, outfit, shoes)
}
