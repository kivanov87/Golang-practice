package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	if number != 0 && number >= -100 && number <= 100 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
