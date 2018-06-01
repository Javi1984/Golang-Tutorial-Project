package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}
type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	//one of the major differences in go below initalizing var in condition like for loop
	//if prefix := "Mr "; isFormal {
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}

}

func GetPrefix(name string) (prefix string) {
	switch name {
	case "Bob":
		prefix = "Mr "
	case "Joe":
		prefix = "Dr "
	case "Mary":
		prefix = "Mrs "
	default:
		prefix = "dude "
	}
	return
}
func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey!" + name
	return
}
func Print(s string) {
	fmt.Print(s)
}
func PrintLine(s string) {
	fmt.Println(s)
}
func PrintCustome(s string, custom string) {
	fmt.Println(s + custom)
}

//function that returns a function
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
