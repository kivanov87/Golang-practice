package main

import (
	"fmt"
)

func main() {
	const priceNylon float32 = 1.50
	const pricePaint float32 = 14.50
	const pricealcohol float32 = 5
	const priceBags float32 = 0.40

	var nylonNeeded int
	var paintNeeded int
	var alcoholNeeded int
	var WorkHoursNeeded int

	fmt.Scanln(&nylonNeeded)
	fmt.Scanln(&paintNeeded)
	fmt.Scanln(&alcoholNeeded)
	fmt.Scanln(&WorkHoursNeeded)

	var totalNylon float32 = (float32(nylonNeeded) + 2) * float32(priceNylon)
	var totalPaint float32 = (float32(paintNeeded) * 1.1) * pricePaint
	var totalAlcohol float32 = float32(alcoholNeeded) * pricealcohol
	var totalPrice float32 = totalAlcohol + totalNylon + totalPaint + priceBags
	var workHour float32 = (totalPrice * 0.3) * float32(WorkHoursNeeded)
	var fullPrice float32 = totalPrice + workHour

	fmt.Println(fullPrice)

}
