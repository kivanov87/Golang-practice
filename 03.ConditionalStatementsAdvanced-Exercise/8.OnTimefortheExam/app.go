package main

import (
	"fmt"
)

func main() {
	var examHour, examMinute, arrivalHour, arrivalMinute int
	fmt.Scan(&examHour, &examMinute, &arrivalHour, &arrivalMinute)

	examTime := examHour*60 + examMinute
	arrivalTime := arrivalHour*60 + arrivalMinute

	if arrivalTime > examTime {
		fmt.Println("Late")
		diff := arrivalTime - examTime
		if diff < 60 {
			fmt.Printf("%d minutes after the start\n", diff)
		} else {
			fmt.Printf("%d:%02d hours after the start\n", diff/60, diff%60)
		}
	} else if arrivalTime == examTime || examTime-arrivalTime <= 30 {
		fmt.Println("On time")
		if examTime-arrivalTime > 0 {
			fmt.Printf("%d minutes before the start\n", examTime-arrivalTime)
		}
	} else {
		fmt.Println("Early")
		diff := examTime - arrivalTime
		if diff < 60 {
			fmt.Printf("%d minutes before the start\n", diff)
		} else {
			fmt.Printf("%d:%02d hours before the start\n", diff/60, diff%60)
		}
	}
}
