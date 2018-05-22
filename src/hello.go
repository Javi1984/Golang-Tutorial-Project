package main

import "fmt"

func main() {
	fmt.Println("Hello Go")

	message := "Hello Go World!"
	// this variagble is a pointer to a string
	var greeting *string = &message

	fmt.Println(message, greeting)

	fmt.Println(message, *greeting)

	// what is happening here
	*greeting = "hi"
	fmt.Println(message, *greeting)

}
