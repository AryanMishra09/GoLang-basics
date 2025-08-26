package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still remaining\n", conferenceTicket, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	/*
		Go is statically typed language
		You need to tell Go compiler, the data type when declaring the variable
		Type inference: But, Go can infer the type when you assign a value
	*/

	var userName string
	var email string
	var userTickets uint
	//ask user for name:
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&userName)
	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)
	fmt.Printf("Enter numberof tickets: ")
	fmt.Scan(&userTickets)
	fmt.Printf("User %v with email %v booked %v tickets\n", userName, email, userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thankyou %v for booking %v tickets. You wil get confirmation mail at %v\n", userName, userTickets, email)
	fmt.Printf("Remaining tickets: %v", remainingTickets)
}
