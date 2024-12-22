package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mutex 	sync.Mutex
	val 	int
}

func New(startValue int) *Counter{
	var mutex sync.Mutex
	return &Counter{
		mutex: 	mutex,
		val:	startValue,
	}
}

func (c *Counter) Increase() {
	c.mutex.Lock()
	c.val++
	c.mutex.Unlock()

}

func (c *Counter) GetVal() int{
	return c.val
}

func main() {
	startVal := 0
	counter := New(startVal)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func ()  {
			for j := 0; j < 100; j++{
				counter.Increase()
			}
			wg.Done()	
		} ()
	}
	wg.Wait()
	fmt.Println(counter.GetVal())
}