package main

import (
	"fmt"

	"github.com/malsolo/go-module/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")

	fmt.Print(message)
}
