package main

import (
	"fmt"
	"sync"
)

func doubleNumber(n int, wg *sync.WaitGroup, resultCh chan<- int) {
	defer wg.Done()
	resultCh <- n * 2
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	wg := &sync.WaitGroup{}
	resultCh := make(chan int, len(numbers))

	for _, v := range numbers {
		wg.Add(1)
		go doubleNumber(v, wg, resultCh)
	}

	wg.Wait()
	close(resultCh)

	for v := range resultCh {
		fmt.Println("Result", v)
	}
}
