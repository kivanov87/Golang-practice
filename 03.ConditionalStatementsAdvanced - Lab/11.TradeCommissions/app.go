package main

import "fmt"

func main() {
	var city string
	var sales, commision float64

	fmt.Scanln(&city)
	fmt.Scanln(&sales)

	switch city {
	case "Sofia":
		if sales < 0 {
			fmt.Println("error")
		} else if sales >= 0 && sales <= 500 {
			commision = 5
		} else if sales > 500 && sales <= 1000 {
			commision = 7
		} else if sales > 1000 && sales <= 10000 {
			commision = 8
		} else if sales > 10000 {
			commision = 12
		}
	case "Varna":
		if sales < 0 {
			fmt.Println("error")
		} else if sales >= 0 && sales <= 500 {
			commision = 4.5
		} else if sales > 500 && sales <= 1000 {
			commision = 7.5
		} else if sales > 1000 && sales <= 10000 {
			commision = 10
		} else if sales > 10000 {
			commision = 13
		}
	case "Plovdiv":
		if sales < 0 {
			fmt.Println("error")
		} else if sales >= 0 && sales <= 500 {
			commision = 5.5
		} else if sales > 500 && sales <= 1000 {
			commision = 8
		} else if sales > 1000 && sales <= 10000 {
			commision = 12
		} else if sales > 10000 {
			commision = 14.5
		}
	default:
		fmt.Println("error")
	}
	if commision > 0 {
		output := sales * commision / 100
		fmt.Printf("%.2f", output)
	}
}
