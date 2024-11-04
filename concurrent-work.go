package main

import (
	"fmt"
	"sync"
	"time"
)

func doubleNumber(number int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()

	// Simulate work
	time.Sleep(2 * time.Second) // Simulate a delay for the work
	resultChan <- number * 2
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int, len(numbers))
	var wg sync.WaitGroup

	// Start the goroutines to double the numbers
	for _, num := range numbers {
		wg.Add(1)
		go doubleNumber(num, &wg, resultChan)
	}

	// Wait for all goroutines to finish or timeout after 1 minute
	doneChan := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneChan) // Signal that work is done
	}()

	// Use time.After to set a timeout for 1 minute
	timeout := time.After(1 * time.Minute)

	// Wait for either the completion of the work or the timeout
	select {
	case <-doneChan:
		// All goroutines finished successfully
		fmt.Println("All operations completed successfully.")
	case <-timeout:
		// Handle timeout
		fmt.Println("Operation timed out.")
	}

	// Collect and print results
	close(resultChan) // Close results channel if done
	for result := range resultChan {
		fmt.Println("Result:", result)
	}
}
