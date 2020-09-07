package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type frenchBot struct{}

func main() {
	eb := englishBot{}
	fb := frenchBot{}
	printGreeting(eb)
	printGreeting(fb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi"
}

func (frenchBot) getGreeting() string {
	return "Bonjour"
}
