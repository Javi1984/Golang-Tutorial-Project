package main

import "fmt"

//using user defined type renaming of string type to our own type
type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

//specify all as astring
func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}
func main() {
	//var s = Salutation{"Joe", "Hello!"}
	//var s = Salutation{greeting: "Hello!", name: "Joe"}
	var s = Salutation{"Bob", "Hello"}
	//call function called greet
	Greet(s)
}
