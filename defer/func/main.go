package main

import "fmt"

func trace(name string) func() {
	fmt.Printf("Entering %s\n", name)
	return func() {
		fmt.Printf("Leaving 1 %s\n", name)
		fmt.Printf("Leaving 2 %s\n", name)
	}
}
func main() {
	defer trace("f")()
	fmt.Println("Doing something")
}
