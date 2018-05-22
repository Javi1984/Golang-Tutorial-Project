package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}
type Printer func(string)

//Pass in a funciton for greet to use called do
//Passing in a paramiter of type funciton:
//Adds value of what is going on
func Greet(salutation Salutation, do Printer) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	do(message)
	do(alternate)
}

//Making things a bit simpler and reusable by defining a type that represents a function
//that takes a string

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	alternate = "Hey!" + name
	return
}
func Print(s string) {
	fmt.Print(s)
}
func PrintLine(s string) {
	fmt.Println(s)
}
func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s, PrintLine)
}
