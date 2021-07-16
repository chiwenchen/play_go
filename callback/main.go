// callback： 將func以參數的形式傳入另一個func

package main

import "fmt"

func sum(f func(...int) []int, nums ...int) int {
	funcResult := f(nums...)

	result := 0

	for _, num := range funcResult {
		result += num
	}

	return result
}

func odd(nums ...int) []int {
	result := []int{}

	for _, num := range nums {
		if (num % 2) != 0 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(sum(odd, numbers...))
}
