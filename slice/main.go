package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(sum(xi...))

	fmt.Println("I am last element:", xi[len(xi)-1])

}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
