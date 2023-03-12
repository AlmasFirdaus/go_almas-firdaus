package main

import (
	"fmt"
	"time"
)

func kelipatan(n chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(<-n)
	}
}

func main() {
	n := make(chan int, 10)

	go kelipatan(n)

	for i := 1; i <= 10; i++ {
		n <- 3 * i
	}

	time.Sleep(1 * time.Second)

}
