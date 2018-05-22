package main

import "fmt"

//using user defined type renaming of string type to our own type
type Salutation string

func main() {
	fmt.Println("Hello Go")
	var message Salutation = "Hello World!"
	fmt.Println(message)
}
