package main

import "../src/greeting"

func main() {
	var s = greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("?"), true)
}
