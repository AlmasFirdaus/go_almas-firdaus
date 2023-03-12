package main

import (
	"fmt"
	"time"
)

func frekuensiString(text string) {
	mapping := map[string]int{}

	for _, txt := range text {

		_, isExist := mapping[string(txt)]
		if !isExist {
			mapping[string(txt)] = 1
		} else if isExist {
			mapping[string(txt)] += 1
		}

		for key, value := range mapping {
			fmt.Println(key, ":", value)
		}
	}
}

func main() {

	go frekuensiString("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
	time.Sleep(time.Second)
}
