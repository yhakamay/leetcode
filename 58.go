package main

func main() {
	//s := "Hello World"
	//s := "   fly me   to   the moon  "
	s := "luffy is still joyboy"
	println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) (length int) {
	for pos := len(s) - 1; pos >= 0; pos-- {
		if s[pos] != ' ' {
			length++
		} else if length != 0 {
			break
		}
	}

	return
}
