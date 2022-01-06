package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//Set properties of the predefined Logger:
	// log entry prefix and a flag to disable printing
	//the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("John")
	//If an error was returned, print it to the console and exit program
	if err != nil {
		log.Fatal(err)
	}

	//If no error was returned, print returned message to console
	fmt.Println(message)
}
