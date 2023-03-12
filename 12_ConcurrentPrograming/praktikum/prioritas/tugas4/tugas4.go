package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println(getNumber())
}

type SafeNumber struct {
	val int
	m   sync.Mutex
}

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}

}

func (i *SafeNumber) Get() int {
	i.m.Lock()

	defer i.m.Unlock()

	return fibo(i.val)
}

func (i *SafeNumber) Set(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

func getNumber() int {
	i := &SafeNumber{}

	go func() {
		i.Set(10)
	}()
	time.Sleep(time.Second)
	return i.Get()
}
