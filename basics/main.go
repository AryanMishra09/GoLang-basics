package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50
	bookings := []string{} //slice

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still remaining\n", conferenceTicket, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	/*
		Go is statically typed language
		You need to tell Go compiler, the data type when declaring the variable
		Type inference: But, Go can infer the type when you assign a value
	*/

	for {
		var userName string
		var email string
		var userTickets uint
		//ask user for name:
		fmt.Printf("Enter your first name: ")
		fmt.Scan(&userName)
		fmt.Printf("Enter your email: ")
		fmt.Scan(&email)
		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(userName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			fmt.Printf("User %v with email %v booked %v tickets\n", userName, email, userTickets)
			bookings = append(bookings, userName)
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("Thankyou %v for booking %v tickets. You wil get confirmation mail at %v\n", userName, userTickets, email)
			fmt.Printf("Remaining tickets: %v\n", remainingTickets)

			firstNames := []string{}
			for _, b := range bookings {
				names := strings.Fields(b) // -> Split the string with space as seperator, and returns silce with the split elements, Ex: "Aryan Mishra" -> ["Aryan", "Mishra"]
				firstName := names[0]
				firstNames = append(firstNames, firstName)
			}
			fmt.Printf("This is all the firstNames of bookins are: %v\n", firstNames)
			if remainingTickets <= 0 {
				//end the program
				fmt.Println("Our confernece is booked out. Come back next year.")
				continue
			}
		}
		if !isValidEmail {
			fmt.Println("Please enter valid email")
		}
		if !isValidName {
			fmt.Println("Please enter valid name")
		}
		if !isValidTicketNumber {
			fmt.Printf("We only have %v tickets remaining, so you can not book %v tickets\n", remainingTickets, userTickets)
			fmt.Println("Please try again")
		}

	}
}
