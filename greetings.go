package greetings

import "fmt"

//hello returns a greeting for the named person

func Hello(name string) string {
	//return a greeting embedding teh name in a message
	message := fmt.Sprintf("Hi %v, Welcome!", name)
	return message
}
