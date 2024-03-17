package main

import (
	"fmt"
)

func main() {
	var start, end, skip rune

	fmt.Scanf("%c\n", &start)

	fmt.Scanf("%c\n", &end)

	fmt.Scanf("%c\n", &skip)

	count := 0
	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			for k := start; k <= end; k++ {
				if i != skip && j != skip && k != skip {
					fmt.Printf("%c%c%c ", i, j, k)
					count++
				}
			}
		}
	}

	fmt.Printf("%d", count)
}
