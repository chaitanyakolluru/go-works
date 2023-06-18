package main

import (
	"fmt"
	"happy-numbers/pkg/squareAdder"
	"sync"
)

func printResults(resultChn chan int, wg *sync.WaitGroup) {
	for {
		fmt.Printf("Number: %d", <-resultChn)
		wg.Done()
	}
}

func main() {
	limit := 100
	var wg sync.WaitGroup
	inputChn := make(chan int, limit)
	resultChn := make(chan int, limit)

	for i := 0; i < limit; i++ {
		wg.Add(1)
		inputChn <- i
	}
	close(inputChn)

	for i := 0; i <= 10; i++ {
		go squareAdder.InvokeChecker(i, inputChn, resultChn, &wg)
	}

	go printResults(resultChn, &wg)

	wg.Wait()

}
