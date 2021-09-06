package main

import "fmt"

func main() {
	runAnyway()
}

func runAnyway() {
	defer fmt.Println("I will run anyway")

	fmt.Println("Run me first")

	return

	fmt.Println("I will not run")
}
