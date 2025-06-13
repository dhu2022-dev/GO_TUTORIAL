package main

import (
	"fmt"
	"log"

	"greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message. One normal and one random.
	message, err := greetings.Hello("Bob")
	messageRand, errRand := greetings.HelloRand("Gladys")
	messages, errMany := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	if errRand != nil {
		log.Fatal(errRand)
	}
	if errMany != nil {
		log.Fatal(errMany)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	fmt.Println(messageRand)
	fmt.Println(messages)
	for name, message := range messages {
		fmt.Printf("%v: %v\n", name, message)
	}
}
