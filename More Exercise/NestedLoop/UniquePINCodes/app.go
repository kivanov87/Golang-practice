package main

import "fmt"

func main() {
	var end1, end2, end3 int
	fmt.Scan(&end1, &end2, &end3)

	for i := 2; i <= end1; i += 2 {
		for j := 2; j <= end2; j++ {
			isPrime := true
			for k := 2; k*k <= j; k++ {
				if j%k == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				for k := 2; k <= end3; k += 2 {
					fmt.Printf("%d %d %d\n", i, j, k)
				}
			}
		}
	}
}
