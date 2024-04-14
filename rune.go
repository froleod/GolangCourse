package main

/*
Реализуйте функцию IsASCII(s string) bool, которая возвращает true,
если строка s состоит только из ASCII символов.
*/

import "unicode"

// BEGIN (write your solution here)
func IsASCII(s string) bool {
	for _, ch := range s {
		if ch > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// END
