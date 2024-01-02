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

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, eMail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, eMail, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicket {
			BookTicket(remainingTickets, userTickets, bookings, firstName, lastName, conferenceName, eMail)

			firstNames := getFirstNames(bookings)
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
func greetUsers(confName string, confTickets int, remTickets uint) {
	fmt.Printf("Wellcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remTickets)
	fmt.Println("Get your ticket here to attend")
}
func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, eMail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(eMail, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicket
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var eMail string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your EMAIL: ")
	fmt.Scan(&eMail)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, eMail, userTickets
}
func BookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, conferenceName string, eMail string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, eMail)
	fmt.Printf("%v tickets remaining for %v ticket.\n", remainingTickets, conferenceName)
}
