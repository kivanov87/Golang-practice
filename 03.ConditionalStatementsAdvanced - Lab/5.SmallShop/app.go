package main

import "fmt"

func main() {

	var city, product string
	var quantity float64

	fmt.Scanln(&product)
	fmt.Scanln(&city)
	fmt.Scanln(&quantity)
	var coffee, water, beer, sweets, peanuts, price float64
	
	
	switch city {
	case "Sofia":
		switch product:{ 
		case "coffee":
			price= 0.50
		case "water": 
			price= 0.80
		case "beer" :
			price= 1.20
		case "sweets": 
			price= 1.45
		case "peanuts": 
			price= 1.60
		}		
	case "Plovdiv":
		switch product:{ 
		case "coffee":
			price= 0.50
		case "water": 
			price= 0.80
		case "beer" :
			price= 1.20
		case "sweets": 
			price= 1.45
		case "peanuts": 
			price= 1.60
		}
		
	case "Varna":
		switch product:{ 
		case "coffee":
			price= 0.50
		case "water": 
			price= 0.80
		case "beer" :
			price= 1.20
		case "sweets": 
			price= 1.45
		case "peanuts": 
			price= 1.60
		}
	}	
	
}
