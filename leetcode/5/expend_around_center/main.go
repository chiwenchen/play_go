package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	startFlag := 0
	endFlag := 0
	maxLen := 0

	for i := 0; i < len(s); i++ {

		lenForOdd := findMaxLength(s, i, i)    // aba
		lenForEven := findMaxLength(s, i, i+1) // abba
		leng := int(math.Max(float64(lenForOdd), float64(lenForEven)))
		if maxLen < leng {
			maxLen = leng
			startFlag = i - (leng-1)/2
			endFlag = i + (leng / 2)
		}
	}

	return s[startFlag : endFlag+1]
}

func findMaxLength(s string, i, j int) int {
	for i >= 0 && j >= i && j < len(s) {
		if s[i] != s[j] {
			break
		}
		i--
		j++
	}
	return j - i + 1 - 2
	// +1 是因為ji相減代表距離(間距數)，但個數為距離+1，例如 i = 0, j = 0的話，個數還是有一個
	// -2 是因為上述演算法是在不符合的情況下才跳出，所以實際上符合的個數為上一輪，所以是 i,j 個往回縮一格，所以一共-2
}
