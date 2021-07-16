// minimize scope variable

package main

import "fmt"

func main() {

	// example 1
	{
		y := 42
		fmt.Println(y)
	}

	// this will not work, because it can not access the value inside closure
	// fmt.Println(y)

	// example 2

	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
