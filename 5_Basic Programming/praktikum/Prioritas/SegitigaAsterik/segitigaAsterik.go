package main

import "fmt"

func main(){
	var input int
	fmt.Print("Masukan Tinggi Pyramid : ")
	fmt.Scan(&input)

	for i := 0; i <= input; i++ {
		for j := 0; j < input - i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k < i + 1; k++ {
			fmt.Print("* ")
		}

		fmt.Println("")
	}
}