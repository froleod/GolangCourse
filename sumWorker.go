package main

/*
Реализуйте функцию-воркера SumWorker(numsCh chan []int, sumCh chan int),
которая суммирует переданные числа из канала numsCh и передает результат в канал sumCh:

numsCh := make(chan []int)
sumCh := make(chan int)

go SumWorker(numsCh, sumCh)
numsCh <- []int{10, 10, 10}

res := <- sumCh // 30
*/

// SumWorker BEGIN (write your solution here)
func SumWorker(numsCh chan []int, sumCh chan int) {
	go func() {
		for nums := range numsCh {
			sum := 0
			for _, num := range nums {
				sum += num
			}
			sumCh <- sum
		}
	}()
}

// END
