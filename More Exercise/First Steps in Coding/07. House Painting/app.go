package main

import "fmt"

func main() {
	var x, y, h int
	const greenPaint float64 = 3.4
	const redPaint float64 = 4.3
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Scanln(&h)

	var sideWalls float64 = float64(x) * float64(y)
	window := 1.5 * 1.5
	var sideWallsTotal float64 = (2 * float64(sideWalls)) - (2 * float64(window))
	backWall := x * x
	door := 1.2 * 2
	frontBack := (2 * float64(backWall)) - door
	totalWalls := sideWallsTotal + frontBack
	//green paint
	gPaintNeed := totalWalls / greenPaint

	var roof float64 = 2 * (float64(x) * float64(y))
	var frontBackRoof float64 = 2 * (float64(x) * float64(h) / 2)
	fBroof := 2 * frontBackRoof
	totalroof := roof + fBroof
	rPaintNeed := totalroof / redPaint

	fmt.Println(roof)
	fmt.Println(frontBackRoof)
	fmt.Println(fBroof)
	fmt.Println(totalroof)

	fmt.Printf("%.2f\n", gPaintNeed)
	fmt.Printf("%.2f\n", rPaintNeed)
}
