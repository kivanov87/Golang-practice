package main

import "fmt"

func main() {
	var number int

	fmt.Scanln(&number)

	var bScores float64

	if number > 1000 {
		bScores := float64(number) * 0.1
	} else if number <= 100 {
		bScores := float64(number) + 5
	} else if number > 100 {
		bScores := float64(number) * 0.2
	}

	fmt.Println(bScores)

}
