package SleepSort

import (
	"sync"
	"time"
)

func GoSleep(nums []int) []int {
	sortc := make(chan int)
	sorted := []int{}

	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		go Sleeping(num, sortc, &wg)
	}

	go func() {
		wg.Wait()
		close(sortc)
	}()

	for num := range sortc {
		sorted = append(sorted, num)
	}

	return sorted
}

func Sleeping(num int, sortc chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	duration := time.Duration(num*30) * time.Millisecond
	time.Sleep(duration)
	sortc <- num
}
