package main

import "fmt"

//using user defined type renaming of string type to our own type
type Salutation struct {
	name     string
	greeting string
}

//anything inside params are consts
//the type of PI is not necessarily defined
//It will bedefined on where you use it
//IOTA it increments constants
const (
	PI       = 3.14
	Language = "Go"
	A        = iota
	B        = iota
	C        = iota
)
const (
	D = iota
	E = iota
	F = iota
)
const (
	G = iota
	H
	I
)

func main() {
	//var s = Salutation{"Joe", "Hello!"}
	//var s = Salutation{greeting: "Hello!", name: "Joe"}
	var s = Salutation{}
	s.name = "Bob"
	fmt.Println(PI)
	fmt.Println(Language)
	//when printing A, B, C  you get 2,3,4 because it was defined at the 2, 3, 4th variable
	fmt.Println(A, B, C)
	//Now these iotas should start from 0
	fmt.Println(D, E, F)
	//These const variables are interpreted as itas
	fmt.Println(G, H, I)
}
