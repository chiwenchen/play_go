package main

func main() {

}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxContainer := 0

	for {
		if l == r {
			break
		}

		maxContainer = Max(Min(height[l], height[r])*(r-l), maxContainer)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}

	return maxContainer
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
