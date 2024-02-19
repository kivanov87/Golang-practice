package main

import (
	"fmt"
	"math"
)

func main() {
	var record, distance, timePerMeter float64
	fmt.Scan(&record, &distance, &timePerMeter)

	totalTime := distance * timePerMeter
	delay := math.Floor(distance/15) * 12.5
	totalTime += delay

	if totalTime < record {
		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.\n", totalTime)
	} else {
		fmt.Printf("No, he failed! He was %.2f seconds slower.\n", totalTime-record)
	}
}
