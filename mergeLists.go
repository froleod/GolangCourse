package main

/* Реализуйте функцию MergeNumberLists(numberLists ...[]int) []int,
которая принимает вариативный список слайсов чисел и объединяет их в 1,
сохраняя последовательность:
MergeNumberLists([]int{1, 2}, []int{3}, []int{4}) // [1, 2, 3, 4]
*/

// MergeNumberLists BEGIN (write your solution here)
func MergeNumberLists(numberLists ...[]int) []int {
	var merged []int
	for _, n := range numberLists {
		if len(n) > 0 {
			merged = append(merged, n...)
		}
	}
	if len(merged) == 0 {
		return []int{}
	}
	return merged
}

// END
