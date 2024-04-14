package main

import "fmt"

/*
Реализуйте функцию GenerateSelfStory(name string, age int, money float64) string,
которая генерирует предложение, подставляя переданные данные в возвращаемую строку.
Например:
GenerateSelfStory("Vlad", 25, 10.00000025) // "Hello! My name is Vlad.
I'm 25 y.o. And I also have $10.00 in my wallet right now."
*/

// BEGIN (write your solution here)
func GenerateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}

// END
