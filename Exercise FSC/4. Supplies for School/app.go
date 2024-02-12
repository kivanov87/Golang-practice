package main

import (
	"fmt"
)

func main() {
	const priceForPen = 5.80
	const priceForMarkers = 7.20
	const priceForAlchohol = 1.20

	var pen int
	var markers int
	var alchohol int
	var percentdiscount int

	fmt.Scanln(&pen)
	fmt.Scanln(&markers)
	fmt.Scanln(&alchohol)
	fmt.Scanln(&percentdiscount)

	var totalPens float64 = float64(pen) * priceForPen
	var totalMarkers float64 = float64(markers) * priceForMarkers
	var totalAlcohol float64 = float64(alchohol) * priceForAlchohol

	var price float64 = totalPens + totalAlcohol + totalMarkers
	var discount float64 = price * float64(percentdiscount) / 100
	var finalPrice float64 = price - float64(discount)
	fmt.Println(finalPrice)
}
