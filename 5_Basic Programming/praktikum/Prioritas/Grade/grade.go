package main

import "fmt"

func main(){

	var input int
	fmt.Print("Masukan nilai : ")
	fmt.Scan(&input)

	if input >= 80 && input <= 100 {
		fmt.Println("A")
	} else if input >= 65 && input <= 79 {
		fmt.Println("B")
	} else if input >= 50 && input <= 64 {
		fmt.Println("C")
	}else if input >= 35 && input <= 49 {
		fmt.Println("D")
	}else if input >= 0 && input <= 34 {
		fmt.Println("E")
	} else if input < 0 || input > 100 {
		fmt.Println("Nilai Invalid")
	}
}