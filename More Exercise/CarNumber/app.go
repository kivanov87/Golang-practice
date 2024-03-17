package main

import (
	"fmt"
)

func main() {
	var start, end int
	fmt.Scan(&start, &end)

	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			for k := start; k <= end; k++ {
				for l := start; l <= end; l++ {
					if (i%2 == 0 && l%2 != 0 || i%2 != 0 && l%2 == 0) && i > l && (j+k)%2 == 0 {
						fmt.Printf("%d%d%d%d ", i, j, k, l)
					}
				}
			}
		}
	}
	fmt.Println()
}
