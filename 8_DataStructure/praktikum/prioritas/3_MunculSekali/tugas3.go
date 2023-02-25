package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {

	var mapping = make(map[string]int)
	var result = []int{}

	for _,v := range angka{
		str := string(v)
		_, exist := mapping[str]

		if !exist {
			mapping[str] = 1
		} else {
			mapping[str] += 1
		}
	}
	
	for key, value := range mapping {
		keyInt, err := strconv.Atoi(key)
		if err == nil && value == 1 {
			result = append(result, keyInt)
		}
	}
	return result
}

func main(){
	fmt.Println(munculSekali("1234123")) // [4]
	fmt.Println(munculSekali("76523752")) // [6, 3]
	fmt.Println(munculSekali("12345")) // [1, 2, 3, 4, 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504")) // [8, 7, 2, 5, 4]
}