package main

import(
	"time"
)

func sleep(duration time.Duration) {
	start := time.Now()
	for {
		if time.Now().Sub(start.Add(duration)) > 0 {
			return
		}
	}
}

func main() {
	sleep(time.Second * 4)
}