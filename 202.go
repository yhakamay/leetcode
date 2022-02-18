package main

import (
	"strconv"
)

func main() {
	isHappy(2)
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	if n < 3 {
		return false
	}

	var nStr string = strconv.Itoa(n)

	n = 0

	for _, c := range nStr {
		digit, _ := strconv.Atoi(string(c))
		n += digit * digit
	}

	return isHappy(n)
}
