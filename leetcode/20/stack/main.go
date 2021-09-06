package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	sl := strings.Split(s, "")

	stack := []string{}

	hash := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, el := range sl {
		if el == "(" || el == "[" || el == "{" {
			stack = append(stack, el)
		} else {
			if len(stack) != 0 && stack[len(stack)-1] == hash[el] {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
