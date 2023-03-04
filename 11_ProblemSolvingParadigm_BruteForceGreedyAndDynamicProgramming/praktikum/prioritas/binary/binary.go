package main

import (
	"fmt"
	"strconv"
)

func reverse(binn string) string {
	reverse := ""

	for i := len(binn) - 1; i >= 0; i-- {
		reverse += string(binn[i])
	}
	return reverse
}

func getBinary(number int) string {
	binn := ""
	for number > 0 {
		binn += strconv.Itoa(number % 2)
		number /= 2
	}
	return reverse(binn)
}

func main() {
	input := 0
	fmt.Print("input : ")
	fmt.Scan(&input)

	arr := []int{}

	for i := 0; i < input+1; i++ {
		value, _ := strconv.Atoi(getBinary(i))

		arr = append(arr, value)
	}
	fmt.Println("output :", arr)
}
