package main

import (
	"fmt"
	"strconv"
)

const X = 122222222122222222

func main() {
	fmt.Println(isPalindrome(X))
}

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)

	for i, cX1 := range strX {
		if i > len(strX)/2 {
			return true
		}

		cX2 := strX[len(strX)-i-1]
		if byte(cX1) != cX2 {
			return false
		}
	}

	return true
}
