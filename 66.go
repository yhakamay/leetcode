/*
INCOMPLETE
*/

package main

import (
	"fmt"
)

func main() {
	digits := []int{1, 2, 9}
	for _, v := range plusOne(digits) {
		fmt.Printf("%d", v)
	}
}

func plusOne(digits []int) (res []int) {
	res = digits

	if onesPlace := res[len(res)-1]; onesPlace < 9 {
		res[len(res)-1] = onesPlace + 1
	} else {
		res[len(res)-1] = 1
		res = append(res, 0)
	}

	return res
}
