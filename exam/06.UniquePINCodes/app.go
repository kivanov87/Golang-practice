package main

import "fmt"

func main() {
	var firstLimit, secondLimit, thirdLimit int
	fmt.Scan(&firstLimit)
	fmt.Scan(&secondLimit)
	fmt.Scan(&thirdLimit)

	for i := 2; i <= firstLimit; i += 2 {
		for j := 2; j <= secondLimit; j++ {
			if j == 2 || j == 3 || j == 5 || j == 7 {
				for k := 2; k <= thirdLimit; k += 2 {
					fmt.Printf("%d %d %d\n", i, j, k)
				}
			}
		}
	}
}
