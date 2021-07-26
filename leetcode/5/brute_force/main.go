package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("abab"))
}

func longestPalindrome(s string) string {

	length := len(s)
	MaxSubStrLen := 0
	MaxSubString := ""

	for i := 0; i < length; i++ {
		for j := length; j > 0; j-- {
			if j < i {
				continue
			}
			subStr := s[i:j]
			if MaxSubStrLen < len(subStr) && isPalindrome(subStr) {
				MaxSubStrLen = len(subStr)
				MaxSubString = subStr
			}
		}
	}

	return MaxSubString
}

func isPalindrome(s string) bool {
	var mid int
	n := len(s)
	if n%2 == 0 {
		mid = (n / 2) - 1
	} else {
		mid = n / 2
	}
	for i := 0; i <= mid; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
