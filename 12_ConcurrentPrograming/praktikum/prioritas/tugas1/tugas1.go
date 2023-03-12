package main

import (
	"fmt"
	"time"
)

func kelipatan(number int) {
	for i := 1; i > 0; i++ {
		fmt.Println(number * i)
	}
}

func main() {
	go kelipatan(3)
	time.Sleep(3 * time.Second)
}
