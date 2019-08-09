package main

import "fmt"

type bot interface {
	getGreeting() float64
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

func (eb englishBot) getGreeting() float64 {
	// VER CUSTOM LOGIC FOR GENERATING AN ENGLISH GREETING

	return (0.5 * 10 * 2)
}

func (sb spanishBot) getGreeting() float64 {
	return (10.5 * 10.5)
}
