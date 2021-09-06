package main

func main() {

}

func maxArea(height []int) int {
	maxContainer := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			maxContainer = Max(Min(height[i], height[j])*(j-i), maxContainer)
		}
	}

	return maxContainer
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
