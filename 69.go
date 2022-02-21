package main

func main() {
	square := 8
	println(mySqrt(square))
}

func mySqrt(x int) (sqrt int) {
	if x == 0 {
		return 0
	}

	for sqrt = 1; sqrt <= x/2; sqrt++ {
		if sqrt*sqrt <= x && (sqrt+1)*(sqrt+1) > x {
			break
		}
	}

	return
}
