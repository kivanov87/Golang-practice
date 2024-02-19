package main

import (
	"fmt"
	"math"
)

func main() {
	var seriesName string
	var episodeDuration, breakDuration int
	fmt.Scan(&seriesName, &episodeDuration, &breakDuration)

	lunchTime := float64(breakDuration) / 8
	restTime := float64(breakDuration) / 4
	freeTime := float64(breakDuration) - lunchTime - restTime

	if freeTime >= float64(episodeDuration) {
		leftTime := math.Ceil(freeTime - float64(episodeDuration))
		fmt.Printf("You have enough time to watch %s and left with %.0f minutes free time.\n", seriesName, leftTime)
	} else {
		needTime := math.Ceil(float64(episodeDuration) - freeTime)
		fmt.Printf("You don't have enough time to watch %s, you need %.0f more minutes.\n", seriesName, needTime)
	}
}
