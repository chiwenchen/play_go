package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type chineseBot struct{}

func main() {
	eb := englishBot{}
	cb := chineseBot{}

	printGreeting(eb)
	printGreeting(cb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (cb chineseBot) getGreeting() string {
	return "Hola!"
}

func (eb englishBot) testing() string {
	return "testing"
}
