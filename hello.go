package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
}

func main() {
	u := User{}
	u.ID = 22
	u.Email = "test@test.com"
	u.FirstName = "John"

	bs, _ := json.Marshal(u)

	fmt.Println(string(bs)) // {"id":22,"email":"test@test.com","first_name":"John"}

	words := []string{"abc", "def", "sdasd", "abc", "defefe", "def"}

	m := map[string]int{}

	for _, word := range words {
		m[word]++
	}

	fmt.Println(m)

	maxVal := 0

	ans := ""

	for key, value := range m {
		if value > maxVal {
			maxVal = value
			ans = key
		}
	}

	fmt.Println(ans)

}
