package main

import "fmt"

func reverse(text string)(string){
	var res string
	for i := len(text)-1; i >= 0; i-- {
		res += string(text[i])
	}
	return res
}

func isPalindrome(input string, resReverse string)(bool){
	return input == resReverse
}

func main(){
	var input string

	fmt.Println("Apakah Palindrome?")
	fmt.Print("masukan kata: ")
	fmt.Scan(&input)

	var resReverse string = reverse(input)
	fmt.Println("captured:", resReverse)

	var resPalindrome bool = isPalindrome(input, resReverse)
	if resPalindrome {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
	
}