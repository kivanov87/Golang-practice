package main

import "fmt"

func main() {
	var num, currentN, s1, s2, s3, s4, s5 int
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&currentN)

		if currentN < 200 {
			s1++
		} else if currentN <= 399 {
			s2++
		} else if currentN <= 599 {
			s3++
		} else if currentN <= 799 {
			s4++
		} else if currentN >= 800 {
			s5++
		}
	}

	percentS1 := float32(s1) / float32(num) * 100
	percentS2 := float32(s2) / float32(num) * 100
	percentS3 := float32(s3) / float32(num) * 100
	percentS4 := float32(s4) / float32(num) * 100
	percentS5 := float32(s5) / float32(num) * 100

	fmt.Printf("%.2f%%\n", percentS1)
	fmt.Printf("%.2f%%\n", percentS2)
	fmt.Printf("%.2f%%\n", percentS3)
	fmt.Printf("%.2f%%\n", percentS4)
	fmt.Printf("%.2f%%", percentS5)
}
