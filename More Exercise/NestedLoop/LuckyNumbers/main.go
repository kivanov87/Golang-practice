package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)

	for n1 := 1; n1 <= 9; n1++ {
		for n2 := 1; n2 <= 9; n2++ {
			for n3 := 1; n3 <= 9; n3++ {
				for n4 := 1; n4 <= 9; n4++ {

					if n1+n2 == n3+n4 && N%(n1+n2) == 0 {
						fmt.Printf("%d%d%d%d ", n1, n2, n3, n4)
					}

				}
			}
		}
	}
}
