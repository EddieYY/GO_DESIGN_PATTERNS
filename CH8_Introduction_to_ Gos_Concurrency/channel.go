package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan string)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(ch chan<- string) {
		ch <- "Hello World! 1"
		println("Finishing goroutine1")
		//channel <- "Hello World! 2"
		//println("Finishing goroutine2")
		waitGroup.Done()
	}(channel)
	time.Sleep(time.Second * 10)
	message := <-channel
	fmt.Println(message)
	waitGroup.Wait()
}
