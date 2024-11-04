package main

import (
	"fmt"
	"sync"
	"time"
)

func doubleNumber(number int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done() // Signal that this goroutine is done

	// Simulate some work with a sleep
	//time.Sleep(2 * time.Second)
	resultChan <- number * 2 // Send the result to the channel
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}            // Numbers to double
	resultChan := make(chan int, len(numbers)) // Channel for results
	var wg sync.WaitGroup

	// Start the goroutines
	for _, num := range numbers {
		wg.Add(1)                             // Increment the WaitGroup counter
		go doubleNumber(num, &wg, resultChan) // Start a goroutine
	}

	// Wait for all goroutines to finish in a separate goroutine
	go func() {
		wg.Wait()         // Wait for all goroutines to finish
		close(resultChan) // Close resultChan after sending all results
	}()

	// Collect results with a timeout
	timeout := time.After(1 * time.Second) // Set a timeout for 1 minute
myloop:
	for {
		select {
		case result, ok := <-resultChan: // Try to read from the channel
			if ok {
				fmt.Println("Result:", result) // Print the result
			} else {
				// Channel has been closed, exit the loop
				break myloop
			}
		case <-timeout: // Handle the timeout case
			fmt.Println("Operation timed out.")
			break myloop
		}
	}

	// Additional code to run after processing results
	fmt.Println("Finished processing results. Proceeding with additional logic...")
	// ... more code can go here
}
