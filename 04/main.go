// для остановки горутин используется context with cancel
// с помоощью signal отлавливается сигнал при нажатии ctrl + C и вызывается cancel()
// при отлавливании сигнала вызывается cancel(),
// в worker() отслеживает был ли вызван cancel(), если да, то выполнение функции завершается
// Данный подхд обеспечивает возможность успешно завершить выполнение обработки данных, полученных из канала
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"os/signal"
	"syscall"
	"os"
)

func mainThread(ch chan<- int) {
	go func ()  {
		for i := 0;; i++ {
			ch <- i
		}
	}()
}

func worker(ctx context.Context, ch <-chan int, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worler %d stopped\n", i)
			return		
		default: 
			num := <-ch
			time.Sleep(time.Second)
			fmt.Printf("worker %d: %d\n", i, num)
		}
	}
}

func startWorkers(ctx context.Context, ch <-chan int, N int) {
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func ()  {
			worker(ctx, ch, i)
			wg.Done()
		} ()
	}
	wg.Wait()
}

func main() {
	var N int
	fmt.Print("enter the number of workers: ")
	fmt.Scan(&N)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	ch := make(chan int, 10)
	mainThread(ch)
	ctx, cancel := context.WithCancel(context.Background())
	go func ()  {
		sig := <- c
		fmt.Println(" Signal received: ", sig)
		cancel()
	} ()

	startWorkers(ctx, ch, N)
	fmt.Println("program completed successfully")	
}