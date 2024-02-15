package main

import "fmt"

func main() {
	var usd float32 = 1.79549
	var bgn float32

	fmt.Scan(&bgn)
	var converted float32 = bgn * usd
	fmt.Println(converted)

}
