package main

import "fmt"

func diagonalSum(arr [][]int) {
	toRight := 0
	toLeft := 0

	x := 1
	for i,v := range arr {
		toRight += v[i]
		toLeft += v[len(v)-x]
		x++
	}


	fmt.Println("to Right",toRight)
	fmt.Println("to Left",toLeft)
	fmt.Println("===========")
	fmt.Println("selisih",toRight, "-", toLeft, "=", (toRight - toLeft) * -1)
}

func main(){
	arr := [][]int{
		{1,2,3},
		{4,5,6},
		{9,8,9},
	}

	diagonalSum(arr)
}

