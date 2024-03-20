package main

import (
	"fmt"
)

func main() {
	var tournaments, startPoints int
	fmt.Scan(&tournaments)
	fmt.Scan(&startPoints)

	points := startPoints
	wins := 0

	for i := 0; i < tournaments; i++ {
		var stage string
		fmt.Scan(&stage)

		switch stage {
		case "W":
			points += 2000
			wins++
		case "F":
			points += 1200
		case "SF":
			points += 720
		}
	}

	fmt.Printf("Final points: %d\n", points)
	fmt.Printf("Average points: %d\n", (points-startPoints)/tournaments)
	fmt.Printf("%.2f%%\n", float64(wins)/float64(tournaments)*100)
}
