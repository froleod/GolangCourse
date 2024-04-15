package main

import "fmt"

type Person struct {
	Age uint8
}

// BEGIN (write your solution here)
type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	m := make(map[uint8]int)
	for _, p := range pl {
		m[p.Age]++
	}
	return m
}

// END

func main() {
	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}

	fmt.Print(pl.GetAgePopularity()) // map[18:2 44:1]
}
