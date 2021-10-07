package main

import "fmt"

func main() {
	fmt.Println(constructSequence(5, 1, 2))
}

func constructSequence(n int32, lo int32, hi int32) []int32 {
	if -1 {
		if 
		p + 1 = recursive
	}
}

func withPointer(n int32, lo int32, hi int32 p int) []int32 {
	sl := make([]int32, int(n))
	for i := 0; i < int(n); i++ {
		newNum := hi - (p - i)
		if newNum < lo {
			return []int32{-1}
		}
		sl[i] = newNum
	}

	return sl
}


if i == 0 {
	sl[i] = hi - 1
} else {
	num := int32((int(hi) - (i - 1)))
	if num < lo {
		return []int32{-1}
	}
	sl[i] = num
}
