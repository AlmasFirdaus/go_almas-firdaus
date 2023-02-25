package main

import "fmt"

func primeNumber(number int) bool {
  for i := 2; i < number; i++ {
	if number % i == 0 {
    return false
	}
  }

  return true
}


func main() {
   fmt.Println(primeNumber(1000000007)) // true
   fmt.Println(primeNumber(13))         // true
   fmt.Println(primeNumber(17))         // true
   fmt.Println(primeNumber(20))         // false
   fmt.Println(primeNumber(35))         // false
}