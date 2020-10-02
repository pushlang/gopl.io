package main

import (
	"fmt"
	"time"
)

func worker(t time.Duration) chan time.Duration {
	ch := make(chan time.Duration)
	go func() {

		time.Sleep(t * time.Second)
		ch <- t

	}()
	return ch
}

func main() {
	start := time.Now()

	a, b := time.Duration(0), time.Duration(0)
	a = <-worker(3)
	b = <-worker(1)

	dur := time.Since(start).Seconds()

	fmt.Println(a, b)
	fmt.Println(dur)
}
