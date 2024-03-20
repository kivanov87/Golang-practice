package main

import (
	"fmt"
	"math"
)

func main() {
	var num, currentN, sum int
	var maxN int32 = math.MinInt32

	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&currentN)

		if int32(currentN) > maxN {
			maxN = int32(currentN)
		}

		sum += currentN
	}

	sum -= int(maxN)

	if sum == int(maxN) {
		fmt.Println("Yes")
		fmt.Printf("Sum = %d", sum)
	} else {
		diff := math.Abs(float64(sum) - float64(maxN))

		fmt.Println("No")
		fmt.Printf("Diff = %.0f", diff)
	}
}
