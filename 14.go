package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	strs := []string{"", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) (ans string) {
	var minLen int

	// Calcurate min length
	for _, str := range strs {
		if len(str) == 0 {
			return
		}

		if utf8.RuneCountInString(str) < minLen || minLen == 0 {
			minLen = utf8.RuneCountInString(str)
		}
	}

	for i := 0; i < minLen; i++ {
		var currentChar string

		for _, str := range strs {
			// If encountered the first rune, set it as the current rune
			if currentChar == "" {
				currentChar = string(str[i])
				continue
			}

			// If the first rune is not equal to the current rune,
			// return immediately
			if string(str[i]) != currentChar {
				return
			}
		}

		ans += string(currentChar)
	}

	return
}
