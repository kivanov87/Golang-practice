package main

import (
	"fmt"
)

func main() {
	var N1, N2 int
	var operator string

	fmt.Scan(&N1)
	fmt.Scan(&N2)
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result := N1 + N2
		if result%2 == 0 {
			fmt.Printf("%d + %d = %d - even\n", N1, N2, result)
		} else {
			fmt.Printf("%d + %d = %d - odd\n", N1, N2, result)
		}
	case "-":
		result := N1 - N2
		if result%2 == 0 {
			fmt.Printf("%d - %d = %d - even\n", N1, N2, result)
		} else {
			fmt.Printf("%d - %d = %d - odd\n", N1, N2, result)
		}
	case "*":
		result := N1 * N2
		if result%2 == 0 {
			fmt.Printf("%d * %d = %d - even\n", N1, N2, result)
		} else {
			fmt.Printf("%d * %d = %d - odd\n", N1, N2, result)
		}
	case "/":
		if N2 == 0 {
			fmt.Printf("Cannot divide %d by zero\n", N1)
		} else {
			result := float64(N1) / float64(N2)
			fmt.Printf("%d / %d = %.2f\n", N1, N2, result)
		}
	case "%":
		if N2 == 0 {
			fmt.Printf("Cannot divide %d by zero\n", N1)
		} else {
			result := N1 % N2
			fmt.Printf("%d %% %d = %d\n", N1, N2, result)
		}
	}
}
