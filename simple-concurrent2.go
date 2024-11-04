package main

import (
	"fmt"
	"sync"
	"time"
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

	doneChan := make(chan struct{})
	go func() {
		wg.Wait()       // Wait for all goroutines to finish
		close(resultCh) // Close results channel after sending all results
		close(doneChan) // Signal that work is done
	}()

	t := time.After(2 * time.Second)
	select {
	case <-doneChan:
		for v := range resultCh {
			fmt.Println("Result", v)
		}
	case <-t:
		fmt.Println("Op timed out")
	}
}
