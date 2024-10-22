package main

import (
	"fmt"
)

func generateNumbers(numbers chan<- int) {
	for i := 1; i <= 10; i++ {
		numbers <- i
	}
	close(numbers)
}

func squareNumbers(numbers <-chan int, squares chan<- int) {
	for num := range numbers {
		squares <- num * num
	}
	close(squares)
}

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go generateNumbers(numbers)
	go squareNumbers(numbers, squares)

	for sq := range squares {
		fmt.Println(sq)
	}
}
