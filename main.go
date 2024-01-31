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
	// arrays
	bookings := [50]string{}
	bookings[0] = firstName + "" + lastName
	fmt.Printf("Thank you, %v %v for purchasing %v tickets. They've been sent to %v", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v remaining tickets for %v", remainingTickets, conferenceName)
	fmt.Printf("Whole bookings array: %v\n", bookings)
	fmt.Printf("The first value in the array: %v\n", bookings[0])
	fmt.Printf("The array tyoe: %T\n", bookings)
	fmt.Printf("The array length: %v\n", len(bookings))

	// slices
	slice := bookings[0:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
}
