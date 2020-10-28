package main

func main() {}

func reverseString_1(s []byte) {
	if len(s) == 0 || len(s) == 1 {
		return
	}

	reverseString_1(s[1:])

	final := s[0]

	for i := 1; i < len(s); i++ {
		s[i-1] = s[i]
	}

	s[len(s)-1] = final
}
