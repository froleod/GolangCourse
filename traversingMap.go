package main

/*
Реализуйте функцию MostPopularWord(words []string) string,
которая возвращает самое часто встречаемое слово в слайсе.
Если таких слов несколько, то возвращается первое из них.
*/

// BEGIN (write your solution here)
func MostPopularWord(words []string) string {
	m := make(map[string]int)

	for _, word := range words {
		m[word]++
	}

	maxCount := 0
	mostPopularWord := ""
	for word, count := range m {
		if count > maxCount {
			maxCount = count
			mostPopularWord = word
		}
	}
	return mostPopularWord
}

// END
