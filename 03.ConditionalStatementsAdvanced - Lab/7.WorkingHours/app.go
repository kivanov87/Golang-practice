package main

import "fmt"

func main() {
	var hour int
	var day string

	switch day {
	case "Monday":
		if hour >= 10 && hour <= 18 {
			fmt.Println("open")
		} else {
			fmt.Println("closed")
		}
	case "Tuesday":
		if hour >= 10 && hour <= 18 {
			fmt.Println("open")
		} else {
			fmt.Println("closed")
		}
	case "Wednesday":
		if hour >= 10 && hour <= 18 {
			fmt.Println("open")
		} else {
			fmt.Println("closed")
		}
	case "Thursday":
		if hour >= 10 && hour <= 18 {
			fmt.Println("open")
		} else {
			fmt.Println("closed")
		}
	case "Friday":
		if hour >= 10 && hour <= 18 {
			fmt.Println("open")
		} else {
			fmt.Println("closed")
		}
	case "Saturday":
		if hour >= 10 && hour <= 18 {
			fmt.Println("open")
		} else {
			fmt.Println("closed")
		}
	case "Sunday":
		if hour >= 10 && hour <= 18 {
			fmt.Println("open")
		} else {
			fmt.Println("closed")
		}
	}

}
