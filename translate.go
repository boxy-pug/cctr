package main

import "unicode"

func ToDigit(r rune) rune {
	if unicode.IsDigit(r) {
		return r
	}
	if unicode.IsLetter(r) {
		lowerCaseR := unicode.ToLower(r)
		if lowerCaseR >= 'a' && lowerCaseR <= 'i' {
			return rune('0' + (lowerCaseR - 'a'))
		} else {
			return '9'
		}
	}
	return '9'
}

func ToPunct(r rune) rune {
	if unicode.IsPunct(r) {
		return r
	}
	return '.'
}

func ToSpace(r rune) rune {
	if unicode.IsSpace(r) {
		return r
	}
	return ' '
}
