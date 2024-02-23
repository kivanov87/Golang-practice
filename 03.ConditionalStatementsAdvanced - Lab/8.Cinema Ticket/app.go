package main

import "fmt"

func main() {

	var day string
	fmt.Scanln(&day)

	switch day {
	case "Monday", "Tuesday", "Friday":
		price := 12
		fmt.Println(price)
	case "Wednesday", "Thursday":
		price := 14
		fmt.Println(price)
	case "Saturday", "Sunday":
		price := 16
		fmt.Println(price)
	}
}
