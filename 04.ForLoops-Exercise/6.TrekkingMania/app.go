package main

import (
	"fmt"
)

func main() {
	var groups int
	fmt.Scan(&groups)

	var totalClimbers int
	musala, montBlanc, kilimanjaro, k2, everest := 0, 0, 0, 0, 0

	for i := 0; i < groups; i++ {
		var climbers int
		fmt.Scan(&climbers)
		totalClimbers += climbers

		switch {
		case climbers <= 5:
			musala += climbers
		case climbers <= 12:
			montBlanc += climbers
		case climbers <= 25:
			kilimanjaro += climbers
		case climbers <= 40:
			k2 += climbers
		default:
			everest += climbers
		}
	}

	fmt.Printf("%.2f%%\n", float64(musala)/float64(totalClimbers)*100)
	fmt.Printf("%.2f%%\n", float64(montBlanc)/float64(totalClimbers)*100)
	fmt.Printf("%.2f%%\n", float64(kilimanjaro)/float64(totalClimbers)*100)
	fmt.Printf("%.2f%%\n", float64(k2)/float64(totalClimbers)*100)
	fmt.Printf("%.2f%%\n", float64(everest)/float64(totalClimbers)*100)
}
