package main

import (
	"fmt"
)

func main() {
	var men, women, tables int
	fmt.Scan(&men, &women, &tables)

	counter := 0
	for m := 1; m <= men; m++ {
		for w := 1; w <= women; w++ {
			if counter == tables {
				return
			}
			fmt.Printf("(%d <-> %d) ", m, w)
			counter++
		}
	}
	fmt.Println()
}
