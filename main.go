package main

import "fmt"

func main() {

	var conferenceName = "Go Conferenc"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available", conferenceTickets, remainingTickets)

}
