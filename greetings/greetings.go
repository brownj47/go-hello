package greetings

import (
	"errors"
	"fmt"
)

//hello returns a greeting for the named person

func Hello(name string) (string, error) {
	//if user enters no name, throw error
	if name == "" {
		return "", errors.New("empty name")
	}

	//return a greeting embedding the name in a message
	message := fmt.Sprintf("Hi %v, Welcome!", name)
	return message, nil
}
