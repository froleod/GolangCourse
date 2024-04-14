package main

/*
Реализуйте функцию ShiftASCII(s string, step int) string,
принимает на вход состоящую из ASCII символов строку s и возвращает новую строку,
где каждый символ из входящей строки сдвинут вперед на число step.
Например:
ShiftASCII("abc", 0) // "abc"
ShiftASCII("abc1", 1) // "bcd2"
ShiftASCII("bcd2", -1) // "abc1"
ShiftASCII("hi", 10) // "rs"
*/

// BEGIN (write your solution here)
func ShiftASCII(s string, step int) string {
	new_str := ""
	for i := 0; i < len(s); i++ {
		ch := byte(s[i])
		ch = byte(int(ch) + step)
		new_str += string(ch)
	}
	return new_str
}

// END
