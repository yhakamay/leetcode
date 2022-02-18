package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string = "}}"

	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := make([]rune, 0, 1)
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	if utf8.RuneCountInString(s)%2 != 0 {
		return false
	}

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[r] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
