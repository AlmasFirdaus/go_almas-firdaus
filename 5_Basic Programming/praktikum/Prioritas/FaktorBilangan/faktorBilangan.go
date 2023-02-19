package main

import "fmt"

func main(){
	
	var input int
	fmt.Print("Faktor Bilangan : ")
	fmt.Scan(&input)

	for i := 1; i <= input; i++ {
		if input % i == 0 {
			fmt.Println(i)
		}
	}
}