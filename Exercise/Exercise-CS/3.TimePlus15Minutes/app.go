package main

import "fmt"

func main() {
	var timeHour, timeMinutes int

	fmt.Scanln(&timeHour)
	fmt.Scanln(&timeMinutes)

	timeMinutes += 15

	var hour int = timeHour + timeMinutes/60
	var minutes int = timeMinutes % 60

	if hour == 24 {
		hour = 0
	}

	fmt.Printf("%d:%.2d", hour, minutes)

}
