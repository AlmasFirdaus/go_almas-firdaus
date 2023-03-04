package main

import "fmt"

func primeX(number int) int {

	arr := make([]int, number)

	loop := true
	x := 3
	index := 1

	for loop {

		if arr[number-1] != 0 {
			loop = false
		} else if arr[0] == 0 {
			arr[0] = 2
		} else {
			faktor := 0
			for i := 1; i <= x; i++ {
				if x%i == 0 {
					faktor++
				}
			}
			if faktor == 2 {
				arr[index] = x
				index++
			}
			x++
		}

	}
	return arr[number-1]

}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
