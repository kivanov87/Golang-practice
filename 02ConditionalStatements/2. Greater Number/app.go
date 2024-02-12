package main

import "fmt"

func main() {
	var fNumber, sNumber int
	fmt.Scanln(&fNumber)
	fmt.Scanln(&sNumber)
	if fNumber < sNumber {
		fmt.Print(sNumber)
	} else {
		fmt.Println(fNumber)
	}
}
