package main

import "fmt"

func main() {
	var deposit float32
	var interest float32
	var period int

	fmt.Println("Въведете сума за депозит")
	fmt.Scanln(&deposit)
	fmt.Println("Въведете срок на депозита")
	fmt.Scanln(&period)
	fmt.Println("Въведете лихвен процент")
	fmt.Scanln(&interest)

	profit := deposit * interest / 100
	var profForMonth float32 = profit / 12
	var total float32 = deposit + (float32(period) * profForMonth)
	fmt.Println(total)
}
