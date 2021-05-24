package main

import "fmt"

/* use this interface to pass in custom types (englishBot / spanishBot)
** this will call getGreeting() and return a string
** only use when the type has a function called getGreeting() associated with it */
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// return greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// return greeting
	return "Hola!"
}
