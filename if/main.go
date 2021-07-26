package main

import "fmt"

func main() {
	// one line if statement.
	if b, _ := compareNumtoHundard(99); b {
		fmt.Println("I am greater than 100")
	} else {
		fmt.Println("I am less, Fuxk")
	}
}

func compareNumtoHundard(num int) (bool, int) {
	if num > 100 {
		return true, num
	} else {
		return false, 100
	}
}
