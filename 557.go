package main

import "strings"

func main() {
	s := "Dog ding"
	println(reverseWords(s))
}

func reverseWords(s string) (reversedWords string) {
	words := strings.Split(s, " ")
	tmpWords := make([]string, 0)

	for _, word := range words {
		tmpChars := []byte(word)
		for i, j := 0, len(tmpChars)-1; i < j; i, j = i+1, j-1 {
			tmpChars[i], tmpChars[j] = tmpChars[j], tmpChars[i]
		}
		tmpWords = append(tmpWords, string(tmpChars))
	}

	reversedWords = strings.Join(tmpWords, " ")

	return
}
