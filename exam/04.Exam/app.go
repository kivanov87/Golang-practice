package main

import "fmt"

func main() {
	var countOfStudents int
	fmt.Scan(&countOfStudents)

	var sum, fail, lowGrade, midGrade, hiGrade float64

	for i := 0; i < countOfStudents; i++ {
		var grade float64
		fmt.Scan(&grade)

		sum += grade

		if grade < 3.00 {
			fail++
		} else if grade < 4.00 {
			lowGrade++
		} else if grade < 5.00 {
			midGrade++
		} else {
			hiGrade++
		}
	}

	fmt.Printf("Top students: %.2f%%\n", hiGrade/float64(countOfStudents)*100)
	fmt.Printf("Between 4.00 and 4.99: %.2f%%\n", midGrade/float64(countOfStudents)*100)
	fmt.Printf("Between 3.00 and 3.99: %.2f%%\n", lowGrade/float64(countOfStudents)*100)
	fmt.Printf("Fail: %.2f%%\n", fail/float64(countOfStudents)*100)
	fmt.Printf("Average: %.2f\n", sum/float64(countOfStudents))
}
