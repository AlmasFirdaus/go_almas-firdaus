package main

import (
	"fmt"
	"time"
)

func kelipatan(n chan int) {
	data := <-n
	for i := 1; i <= 10; i++ {
		fmt.Println(data * i)
	}
}

func main() {
	n := make(chan int)

	go kelipatan(n)

	n <- 3
	time.Sleep(time.Millisecond)
}
