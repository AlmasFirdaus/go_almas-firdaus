package main

import (
	"fmt"
)

func Absolute(num int) int {
	if num < 0 {
		return num * -1
	} else {
		return num
	}
}

func sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func frog(stones []int) int {
	all_step := []int{}
	step_odd := []int{}
	step_even := []int{Absolute(stones[0] - stones[1])}
	for i := 0; i < len(stones)-1; i++ {
		all_step = append(all_step, Absolute(stones[i]-stones[i+1]))
	}
	for i := 0; i < len(stones)-2; i += 2 {
		step_odd = append(step_odd, Absolute(stones[i]-stones[i+2]))

	}
	for i := 1; i < len(stones)-2; i += 2 {
		step_even = append(step_even, Absolute(stones[i]-stones[i+2]))

	}
	fmt.Println(step_odd)
	fmt.Println(step_even)

	if len(stones)%2 == 0 {
		if sum(step_odd)+Absolute(step_odd[len(step_odd)]-all_step[len(all_step)-1]) < sum(step_even) {
			return sum(step_odd) + Absolute(step_odd[len(step_odd)]-all_step[len(all_step)-1])
		} else {
			return sum(step_even)
		}
		fmt.Println("yes")
	} else {
		if sum(step_odd) < sum(step_even) {
			return sum(step_odd)
		} else {
			return sum(step_even)
		}
	}
}

func main() {
	fmt.Println(frog([]int{10, 30, 40, 20}))
}
