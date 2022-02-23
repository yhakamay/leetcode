package main

import "fmt"

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2

	// Print the map
	for _, v := range image {
		for _, w := range v {
			fmt.Printf("%d ", w)
		}
		fmt.Println()
	}

	fmt.Println()

	newImage := floodFill(image, sr, sc, newColor)

	for _, v := range newImage {
		for _, w := range v {
			fmt.Printf("%d ", w)
		}
		fmt.Println()
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	prevColor := image[sr][sc]
	m, n := len(image), len(image[0])

	if prevColor == newColor {
		return image
	}

	floodFillUtil(image, sr, sc, prevColor, newColor, m, n)
	return image
}

func floodFillUtil(image [][]int, sr, sc, prevColor, newColor, m, n int) {
	if !shouldRepaint(image, sr, sc, prevColor, newColor, m, n) {
		return
	}

	image[sr][sc] = newColor

	floodFillUtil(image, sr+1, sc, prevColor, newColor, m, n)
	floodFillUtil(image, sr-1, sc, prevColor, newColor, m, n)
	floodFillUtil(image, sr, sc+1, prevColor, newColor, m, n)
	floodFillUtil(image, sr, sc-1, prevColor, newColor, m, n)
}

func shouldRepaint(image [][]int, sr, sc, prevColor, newColor, m, n int) bool {
	if sr < 0 || sr >= m || sc < 0 || sc >= n {
		return false
	}
	if image[sr][sc] != prevColor {
		return false
	}
	if image[sr][sc] == newColor {
		return false
	}

	return true
}
