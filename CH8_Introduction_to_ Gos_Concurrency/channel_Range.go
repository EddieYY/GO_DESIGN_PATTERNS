package main

import "time"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
		close(ch) // must to close, range will keep iterating until channel is close
	}()
	//Range is very useful in taking data from a channel, and it's commonly used in fan-in
	//   patterns where many different Goroutines send data to the same channel.
	for v := range ch {
		println(v)
	}
}
