package main

import "fmt"

func main() {
	haystack, needle := "hel9gralehollo", "ll"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if haystack == "" || needle == "" {
		return 0
	}

	firstCharOfNeedle := string(needle[0])

	// Loop of the haystack
	for i := 0; i < len(haystack)-1; i++ {
		if string(haystack[i]) == firstCharOfNeedle {

			// Loop of the needle
			for j := 0; j < len(needle); j++ {
				// Found it!
				if j == len(needle) && haystack[i+j] == needle[j] {
					return i
				}

				// This is wrong...
				if haystack[i+j] != needle[j] {
					break
				}
			}
		}
	}

	return -1
}
