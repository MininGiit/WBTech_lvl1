package main

import (
	"context"
	"fmt"
	"time"
)

var flag = true

func WorckerWithChan(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worcker running")
			time.Sleep(time.Second)
		}
	}
}

func WorckerWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worcker running")
			time.Sleep(time.Second)
		}
	}
}

func WorckerWithClobalVar() {
	for {
		if flag {
			fmt.Println("Worcker running")
			time.Sleep(time.Second)
		} else {
			fmt.Println("Worker stopped")
			return
		}
	}
}

func main() {
	fmt.Println("start worcker1 (with chanel)")
	done := make(chan struct{})
	go WorckerWithChan(done)
	time.Sleep(time.Second * 3)
	done <- struct{}{}
	fmt.Println("stop worcker1")

	time.Sleep(time.Second * 2)
	fmt.Println()

	fmt.Println("start worcker2 (with cancel)")
	ctx, cancel := context.WithCancel(context.Background())
	go WorckerWithContext(ctx)
	time.Sleep(time.Second * 4)
	cancel()
	fmt.Println("stop worcker2")

	time.Sleep(time.Second * 2)
	fmt.Println()

	fmt.Println("start worcker3 (with timeout)")
	ctx, _ = context.WithTimeout(context.Background(), time.Second * 4)
	go WorckerWithContext(ctx)
	<-ctx.Done()
	fmt.Println("stop worcker3")

	time.Sleep(time.Second * 2)
	fmt.Println()

	fmt.Println("start worcker4 (with global var)")	
	go WorckerWithClobalVar()
	time.Sleep(time.Second * 4)
	flag = false
	fmt.Println("stop worcker4")

	time.Sleep(time.Second * 2)
}