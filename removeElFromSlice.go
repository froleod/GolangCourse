package main

/*
В Go нет встроенной функции удаления элемента из слайса.
	Реализуйте функцию Remove(nums []int, i int) []int,
которая удаляет элемент по индексу i из слайса nums.
	Если приходит несуществующий индекс, то из функции возвращается исходный слайс.
	Порядок элементов в итоговом слайсе после удаления может быть изменен.
*/

// Remove BEGIN (write your solution here)
func Remove(nums []int, i int) []int {
	if i >= len(nums) || i < 0 {
		return nums
	}
	newNums := make([]int, 0, len(nums)-1)

	for j := 0; j < i; j++ {
		newNums = append(newNums, nums[j])
	}

	for k := i + 1; k < len(nums); k++ {
		newNums = append(newNums, nums[k])
	}

	return newNums

}

// END
