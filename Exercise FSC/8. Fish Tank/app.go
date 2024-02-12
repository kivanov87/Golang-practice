package main

import "fmt"

func main() {
	var lenght, width, height int
	var percent float32

	fmt.Scanln(&lenght)
	fmt.Scanln(&width)
	fmt.Scanln(&height)
	fmt.Scanln(&percent)

	volume := lenght * width * height

	var convertVolume float32 = float32(volume) * 0.001
	var usedSpace float32 = percent / 100
	var requiredLitres float32 = convertVolume * (1 - usedSpace)

	fmt.Println(requiredLitres)

}
