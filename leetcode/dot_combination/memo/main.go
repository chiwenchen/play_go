package main

import "fmt"

func main() {
	fmt.Println(DotCombination(3, 2))
	fmt.Println(DotCombination(5, 3))
	fmt.Println(DotCombination(99, 10))
}

func DotCombination(target int, max int) [][]int {
	if target == 0 {
		return [][]int{[]int{}}
	}

	if target < 0 {
		return nil
	}

	result := make([][]int, 0)

	for i := 1; i <= max; i++ {
		remaining := target - i
		memo := DotCombination(remaining, i)
		for _, val := range memo {
			val := append(val, i)
			result = append(result, val)
		}
	}

	return result
}
