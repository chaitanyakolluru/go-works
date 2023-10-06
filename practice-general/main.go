package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func f(from string, done chan int) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
	done <- 1
}

func channelSync() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}

func channels() {
	messages := make(chan string, 2)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

}

func goroutines() {

	done := make(chan int, 2)

	go f("direct", done)

	go f("goroutine", done)

	<-done
	<-done

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// time.Sleep(time.Second)
	// fmt.Println("done")
}

func main() {
	goroutines()
	channels()
	channelSync()
}
