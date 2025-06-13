package main

import (
	"fmt"
	"log"

	"greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Gladys"))
	fmt.Println(greetings.HelloRand("Gladys"))

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message. One normal and one random.
	message, err := greetings.Hello("")
	messageRand, errRand := greetings.HelloRand("Gladys")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	if errRand != nil {
		log.Fatal(errRand)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	fmt.Println(messageRand)
}
