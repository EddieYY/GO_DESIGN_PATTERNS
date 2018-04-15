package channel_singleton

import (
	"fmt"
	"testing"
	"time"
)

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

func init() {
	var count int
	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)
}

type singleton struct{}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()
	n := 5000
	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}
	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())
	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	singleton.Stop()
}
