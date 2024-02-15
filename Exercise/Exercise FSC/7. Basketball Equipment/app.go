package main

import "fmt"

func main() {
	var tax int
	fmt.Scanln(&tax)
	var sneackersPrice float32 = float32(tax) * 0.6
	clothes := sneackersPrice * 0.8
	basketBall := clothes / 4
	accessory := basketBall / 5

	total := float32(tax) + sneackersPrice + clothes + basketBall + accessory

	fmt.Println(total)

}
