package main

import "fmt"

func main() {
	var sea, mountain int
	fmt.Scan(&sea)
	fmt.Scan(&mountain)

	var profit int
	var input string

	for {
		fmt.Scan(&input)

		if input == "Stop" {
			break
		}

		if input == "sea" && sea > 0 {
			profit += 680
			sea--
		} else if input == "mountain" && mountain > 0 {
			profit += 499
			mountain--
		}

		if sea == 0 && mountain == 0 {
			fmt.Println("Good job! Everything is sold.")
			break
		}
	}

	fmt.Printf("Profit: %d leva.\n", profit)
}
