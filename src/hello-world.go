package main

import (
   "fmt"
   "reflect"
 )

//to declare variables at package level it needs to be var?
var (
	name string
	course string 
	module float64
)

func main() {
	fmt.Println("Name is set to ", name, " and is of type ", reflect.TypeOf(name))
	fmt.Println("Module is set to ", module, " and is of type ", reflect.TypeOf(module))
}