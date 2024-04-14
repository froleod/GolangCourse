package main

import (
	"strings"
	"unicode"
)

/*
В пакете unicode есть функция unicode.Is(unicode.Latin, rune), которая проверяет,
что руна является латинским символом или нет.

Реализуйте функцию LatinLetters(s string) string,
которая возвращает только латинские символы из строки s.
Например:
LatinLetters("привет world!") // "world"
*/

// LatinLetters BEGIN (write your solution here)
func LatinLetters(s string) string {
	sb := strings.Builder{}
	for _, char := range s {
		if unicode.Is(unicode.Latin, char) {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}

// END
