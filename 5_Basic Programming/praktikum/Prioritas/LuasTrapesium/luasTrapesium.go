package main

import "fmt"

func main() {
	var tinggi, a, b int = 10, 10, 5
	var luas int = tinggi * (a + b) / 2

	fmt.Println("Luas trapesium adalah", luas)
}