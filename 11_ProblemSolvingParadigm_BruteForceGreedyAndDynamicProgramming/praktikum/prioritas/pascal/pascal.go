package main

import "fmt"

func pascal(num int) [][]int {
	var result [][]int
	for i := 1; i <= num; i++ {
		if i == 1 {
			result = append(result, []int{1})
		} else if i == 2 {
			result = append(result, []int{1, 1})
		} else {
			var temp = []int{1}
			for j := 0; j < len(result[len(result)-2]); j++ {
				temp = append(temp, result[len(result)-1][j]+result[len(result)-1][j+1])
			}
			temp = append(temp, 1)
			result = append(result, temp)
		}
		fmt.Println(result[len(result)-1])
	}
	return result
}

func main() {
	pascal(5)
	// fmt.Print(pascal(5))
	// var tes = [][]int{{1}, {1, 1}, {1, 2, 1}}
	// fmt.Print(tes[len(tes)-1:])
}
