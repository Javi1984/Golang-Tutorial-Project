package main

import "fmt"

//using user defined type renaming of string type to our own type
type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	fmt.Println(salutation.name)
	fmt.Println(salutation.greeting)
}

func main() {
	//var s = Salutation{"Joe", "Hello!"}
	//var s = Salutation{greeting: "Hello!", name: "Joe"}
	var s = Salutation{"Bob", "Hello"}
	//call function called greet
	Greet(s)
}
