package main

import "fmt"

type Person struct {
	Name string
}

type Dancer interface {
	Dance() string
}

type Employee interface {
	getSalary()
	work()
}

func (p Person) Dance() string {
	return p.Name
}

// Performing is a function for dancer interface
func Performing(d Dancer) {

	fmt.Println(d.Dance(), "I can dance on stage")
}

// 這個會錯，因為 interface 不能當作 receiver
// func (d Dancer) notPerforming() {
// 	fmt.Println(d.Dance(), "I can dance on stage")
// }

func main() {
	ivan := &Person{"Ivan"}

	Performing(ivan)
}
