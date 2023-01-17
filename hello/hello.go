package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{
		"gladys",
		"Samantha",
		"Darryn",
	}
	messages, err := greetings.Hellos(names)

	//req greeting message

	// Get a greeting message and print it.
	// message, err := greetings.Hello("Justus")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
