package main

import "fmt"

func caesar(offset int, input string) string {
	result := ""
	for _, v := range input {
		rune := v + rune(offset)
		if rune > 122 {
			rune = rune - 122
			if rune > 26 {
				rune = 96 + (rune % 26)
				if rune == 96 {
					rune = 96 + 26
				}
			} else {
				rune = 96 + rune

			}
		}
		result += string(rune)

	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
