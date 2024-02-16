package main

import "fmt"

func main() {
	var time1, time2, time3 int
	fmt.Scanln(&time1)
	fmt.Scanln(&time2)
	fmt.Scanln(&time3)

	sumTime := time1 + time2 + time3

	var minutes int = sumTime / 60
	seconds := sumTime % 60

	if seconds < 10 {
		fmt.Printf("%d:%.2d", minutes, seconds)
	} else {
		fmt.Printf("%d:%d", minutes, seconds)
	}
}
