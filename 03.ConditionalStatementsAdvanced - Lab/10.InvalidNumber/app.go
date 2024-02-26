package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	var isvalid bool = number >= 100 && number <= 200 || number == 0

	if !isvalid {
		fmt.Println("invalid")
	}
}
