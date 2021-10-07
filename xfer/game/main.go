package main

import "fmt"

func main() {
	fmt.Println(countMoves([]int32{3, 4, 6, 6, 3}))
}

func countMoves(numbers []int32) int64 {
	count := int64(0)
	return counter(numbers, count)

}

func counter(numbers []int32, count int64) int64 {
	if allSameInts(numbers) {
		return 0
	}
	maxNum := int32(0)
	maxIndex := int32(0)
	for i, num := range numbers {
		if num > maxNum {
			maxNum = num
			maxIndex = int32(i)
		}
		numbers[i] = numbers[i] + 1
	}

	numbers[maxIndex] = numbers[maxIndex] - 1

	count = counter(numbers, count) + 1
	return count

}

func allSameInts(a []int32) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}
