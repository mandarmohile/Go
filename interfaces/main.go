package main

import "fmt"

type bot interface {
	getGreeting() string
}

type bot2 interface {
	getGreeting(int) string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	// Practice
	somevar := make(map[string]bot)
	somevar["english"] = eb
	fmt.Println(somevar)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
