package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	fmt.Println("message: " + message)
	fmt.Println("alternate: " + alternate)
}

//vardic functions
//make name string because we redefine greeting as varidic fucntion
//Greeting is going to be a slice
//Important to know how many variadic variables are passed in

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	alternate = "Hey!" + name
	return
}
func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
