package main

import (
	"fmt"
)

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Printf("checkInclusion(s1, s2): %v\n", checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	charsS1 := make(map[byte]bool)

	for i, charS2 := range s2 {
		// If v is not included in charsS2, continue
		for _, charS1 := range s1 {
			if charS2 == charS1 {
				for j := i; j < len(s2); j++ {
				}
			}
		}

		// Else, search more in inner loop, staring v to len(charsS1)
	}
}
