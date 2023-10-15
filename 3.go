package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Printf("lengthOfLongestSubstring(s): %v\n", lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) (max int) {
	if len(s) == 0 {
		return 0
	}

	max = 1
	var tmpCount int

	for i := 0; i < len(s); {
		if tmpCount > max {
			max = tmpCount
		}

		tmpCount = 0
		substr := make(map[byte]bool)

		if substr[s[i]] == true {
			i = tmpCount
			break
		}

		tmpCount++
		substr[s[i]] = true
	}

	return
}
