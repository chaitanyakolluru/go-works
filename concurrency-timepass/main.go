package main

import (
	"math/rand"
	"sync"
)

func getRandomNumber(randChannel chan int) {
	randChannel <- rand.Intn(1000)
}

func processRandomNumber(randChannel chan int, wg *sync.WaitGroup) {
	for {
		println("Random number is: ", <-randChannel)
		wg.Done()
	}

}

func main() {
	wg := sync.WaitGroup{}
	randChannel := make(chan int)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go getRandomNumber(randChannel)
	}

	go processRandomNumber(randChannel, &wg)

	wg.Wait()
}
