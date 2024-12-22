package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func writeInCh(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return		
		default: 
			num := rand.Int()
			ch <- num 
		}
	}
}

func outFromCh(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return		
		default: 
			select {
			case num := <- ch:
				fmt.Println(num)
			default:
				continue
			}
		}
	}
}

func main() {
	N := 3
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(N) * time.Second)
	ch := make(chan int)
	go writeInCh(ctx, ch)
	outFromCh(ctx, ch)
}