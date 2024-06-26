package main

import "sort"

/*
Реализуйте функцию UniqueSortedUserIDs(userIDs []int64) []int64,
которая возвращает отсортированный слайс, состоящий из уникальных
идентификаторов userIDs.
*/

// BEGIN (write your solution here)
func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}

// END
