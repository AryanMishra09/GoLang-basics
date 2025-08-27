package main

import (
	"basics/helper"
	"fmt"
	"strconv"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTicket = 50

var remainingTickets uint = 50

var bookings = make([]map[string]string, 0) //slice

func main() {

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)
	greetUser()

	/*
		Go is statically typed language
		You need to tell Go compiler, the data type when declaring the variable
		Type inference: But, Go can infer the type when you assign a value
	*/

	for {
		//ask user for name:
		userName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			fmt.Printf("User %v with email %v booked %v tickets\n", userName, email, userTickets)

			bookTicket(userName, email, userTickets)

			firstNames := getFirstNames()
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

		city := "London"
		fmt.Printf("City selected: %v\n", city)
		switch city {
		case "New York":
		//new york code
		case "Singapore":
		//singapore code
		case "London", "Berlin":
		//Code
		case "Delhi":
		//Code
		default:
			fmt.Println("No valid city selected")
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still remaining\n", conferenceTicket, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, b := range bookings {
		names := strings.Fields(b["userName"]) // -> Split the string with space as seperator, and returns silce with the split elements, Ex: "Aryan Mishra" -> ["Aryan", "Mishra"]
		firstName := names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func getUserInput() (string, string, uint) {
	var userName string
	var email string
	var userTickets uint
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&userName)
	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)
	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return userName, email, userTickets
}

func bookTicket(userName string, email string, userTickets uint) {

	//Create a map for user:
	userData := make(map[string]string)
	userData["userName"] = userName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thankyou %v for booking %v tickets. You wil get confirmation mail at %v\n", userName, userTickets, email)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	fmt.Printf("List of bookings: %v ", bookings)
}
