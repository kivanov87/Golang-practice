package main

import (
	"fmt"
	"math"
)

func main() {
	var radiance float32

	fmt.Scanln(&radiance)
	degreеse := radiance * 180 / math.Pi
	fmt.Println(degreеse)
}
