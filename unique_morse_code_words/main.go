package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// 26 letters of the English alphabet
var morseCodes = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations_1(words []string) int {
	exist := make(map[string]struct{})

	for _, word := range words {
		code := ""

		for _, c := range word {
			code += morseCodes[c-'a']
		}

		if _, ok := exist[code]; !ok {
			exist[code] = struct{}{}
		}
	}

	return len(exist)
}
