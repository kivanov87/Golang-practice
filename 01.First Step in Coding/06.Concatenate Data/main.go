package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var age int
	var town string

	fmt.Scanln(&firstName)
	fmt.Scanln(&lastName)
	fmt.Scanln(&age)
	fmt.Scanln(&town)

	fmt.Printf("You are %v %v, a %v-years old person from %v.", firstName, lastName, age, town)
}
