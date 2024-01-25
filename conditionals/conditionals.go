package main

import "fmt"

func main() {

	maxMessageLen := 20

	var message string
	fmt.Println("Enter your message...")
	fmt.Scanln(&message)

	fmt.Println("Trying to send a message of length", len(message))

	if len(message) > maxMessageLen {
		fmt.Println("Message exceeds characters; not sent")
	} else {
		fmt.Println("Message sent")
	}
}
