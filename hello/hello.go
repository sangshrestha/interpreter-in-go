package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request a greeting message.
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		errMessage := fmt.Sprintf("Err: %v", err)
		log.Fatal(errMessage)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(messages)
}
