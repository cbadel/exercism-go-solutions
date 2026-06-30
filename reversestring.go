package reversestring

func Reverse(input string) (c string) {
	karakter := []rune(input)
	s := len(karakter) - 1
	for s >= 0 {
		c += string(karakter[s])
		s--
	}
	return c
}
