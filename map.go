package main

/*
Реализуйте функцию UniqueUserIDs(userIDs []int64) []int64,
которая возвращает слайс, состоящий из уникальных идентификаторов userIDs.
Порядок слайса должен сохраниться.

UniqueUserIDs([]int64{55, 55, 33, 22, 55})  # [55 33 22]
*/

// UniqueUserIDs BEGIN (write your solution here)
func UniqueUserIDs(userIDs []int64) []int64 {
	if len(userIDs) == 0 {
		return []int64{}
	}

	uniqueMap := make(map[int64]bool)
	var uniqueIDs []int64

	for _, id := range userIDs {
		if !uniqueMap[id] {
			uniqueMap[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	return uniqueIDs
}

// END
