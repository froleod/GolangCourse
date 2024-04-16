package main

import "time"

/*
Реализуйте функцию MaxSum(nums1, nums2 []int) []int из прошлого задания,
используя горутины для расчета каждой суммы слайса.
Не забудьте использовать функцию time.Sleep(100 * time.Millisecond), чтобы сумма успела посчитаться.
В настоящих приложениях используются специальные инструменты, чтобы дожидаться исполнения асинхронного кода,
но для простоты здесь будет использоваться обычный сон.
*/

func MaxSum(nums1, nums2 []int) []int {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		sum := 0
		for _, n := range nums1 {
			sum += n
		}
		time.Sleep(100 * time.Millisecond)
		ch1 <- sum
	}()

	go func() {
		sum := 0
		for _, n := range nums2 {
			sum += n
		}
		time.Sleep(100 * time.Millisecond)
		ch2 <- sum
	}()

	sum1 := <-ch1
	sum2 := <-ch2

	if sum1 >= sum2 {
		return nums1
	} else {
		return nums2
	}
}
