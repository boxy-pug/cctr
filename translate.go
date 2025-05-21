package main

import "unicode"

func ToDigit(r rune) rune {
	if unicode.IsDigit(r) {
		return r
	}
	if unicode.IsLetter(r) {
		return rune(r % 10)
	}
	return 0
}
