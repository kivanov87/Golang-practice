package main

import (
	"fmt"
)

func main() {
	var month string
	var nights int
	fmt.Scan(&month, &nights)

	var studioPrice, apartmentPrice float64

	switch month {
	case "May", "October":
		studioPrice = 50
		apartmentPrice = 65
	case "June", "September":
		studioPrice = 75.20
		apartmentPrice = 68.70
	case "July", "August":
		studioPrice = 76
		apartmentPrice = 77
	}

	studioCost := studioPrice * float64(nights)
	apartmentCost := apartmentPrice * float64(nights)

	if month == "May" || month == "October" {
		if nights > 14 {
			studioCost *= 0.7
		} else if nights > 7 {
			studioCost *= 0.95
		}
	} else if month == "June" || month == "September" {
		if nights > 14 {
			studioCost *= 0.8
		}
	}

	if nights > 14 {
		apartmentCost *= 0.9
	}

	fmt.Printf("Apartment: %.2f lv.\n", apartmentCost)
	fmt.Printf("Studio: %.2f lv.\n", studioCost)
}
