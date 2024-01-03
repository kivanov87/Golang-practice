package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicket {
			BookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("first name of bookings are: %v\n", firstNames)

			//end program
			if remainingTickets == 0 {
				fmt.Println("We are out of tickets")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Email you entered is invalid")
			}
			if !isValidTicket {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}
func greetUsers() {
	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your EMAIL: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func BookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v ticket.\n", remainingTickets, conferenceName)
}
