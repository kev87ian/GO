package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to the %v booking application! \n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets, and", remainingTickets, "are still available.")

	// user input
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name;")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you need")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you, %v %v for purchasing %v tickets. They've been sent to %v", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v remaining tickets for %v", remainingTickets, conferenceName)

	// arrays and slices

}
