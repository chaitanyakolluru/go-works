package main

import (
	"fmt"
	"happy-numbers/pkg/squareAdder"
)

func printResults(resultChn chan int) {
	for result := range resultChn {
		fmt.Printf("Number: %d\n", result)
	}
}

func main() {
	limit := 100
	inputChn := make(chan int, limit)
	resultChn := make(chan int, 10)

	for i := 0; i < limit; i++ {
		inputChn <- i
	}
	close(inputChn)

	for i := 0; i <= 5; i++ {
		go squareAdder.InvokeChecker(i, inputChn, resultChn)

	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Number: %d\n", <-resultChn)
	}
}
