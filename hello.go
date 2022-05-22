package main

import (
	"fmt"

	"github.com/mariogonzalez99/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Mario")
	fmt.Println(message)
}
