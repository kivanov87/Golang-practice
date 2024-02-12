package main

import (
	"fmt"
)

func main() {
	var w float64
	var h float64

	fmt.Print("Въведете дължината на учебната зала (w): ")
	fmt.Scan(&w)
	fmt.Print("Въведете широчината на учебната зала (h): ")
	fmt.Scan(&h)

	var wcm = w * 100
	var hcm = h * 100

	// Calculate available space for desks
	corridorWidth := 100.0
	deskWidth := 70.0
	deskLength := 120.0
	//entranceDoorWidth := 160.0
	//podiumWidth := 160.0
	podiumLength := 120.0

	usableWidth := wcm - corridorWidth
	usableLength := hcm - podiumLength

	desksPerRow := int(usableWidth / deskWidth)
	rows := int(usableLength / deskLength)

	totalDesks := desksPerRow * rows
	deskClear := totalDesks - 3

	fmt.Printf("Брой места в учебната зала: %d\n", deskClear)
}
