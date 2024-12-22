package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[string] int, 10) 
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(){
			mutex.Lock()
			data["qwe"] += i 
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(data["qwe"])
}