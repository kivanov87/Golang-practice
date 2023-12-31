package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}
	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
	for {
		var firstName string
		var lastName string
		var eMail string
		var userTickets uint

		//ask user for name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your EMAIL: ")
		fmt.Scan(&eMail)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) <= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(eMail, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicket {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, eMail)
			fmt.Printf("%v tickets remaining for %v ticket.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

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

			fmt.Println("your input data is invalid try again")
		}
	}
}
