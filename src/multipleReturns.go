package main

import "fmt"

//using user defined type renaming of string type to our own type
type Salutation struct {
	name     string
	greeting string
}

//when returning multiple varaibles you have to use both
//if you don't use both then there is an error
func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println("Message: " + message)
	fmt.Println("alternate: " + alternate)
}

//specify all as astring
func CreateMessage(name, greeting string) (string, string) {
	return greeting + " " + name, "Hey!" + name
}
func main() {
	//var s = Salutation{"Joe", "Hello!"}
	//var s = Salutation{greeting: "Hello!", name: "Joe"}
	var s = Salutation{"Bob", "Hello"}
	//call function called greet
	Greet(s)
}
