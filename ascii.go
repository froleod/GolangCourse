package main

/*
Реализуйте функции NextASCII(b byte) byte и PrevASCII(b byte) byte,
которые возвращают следующий или предыдущий символ типа byte из ASCII таблицы.
Например:
NextASCII(byte('a')) // 98
PrevASCII(byte('b')) // 97
Допускаем, что в функцию PrevASCII не может прийти нулевой символ,
а в функцию NextASCII — последний символ ASCII таблицы.
*/

// BEGIN (write your solution here)
func NextASCII(b byte) byte {
	return byte(b + 1)
}

func PrevASCII(b byte) byte {
	return byte(b - 1)
}

// END
