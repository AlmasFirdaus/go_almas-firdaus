package main

import "fmt"

func SimpleEquations(a, b, c int) {
	stop := false
	for x := 1; 1 <= 10000; x++ {
		for y := 1; y <= 10000; y++ {
			for z := 1; z <= 10000; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+z+z == c {
					fmt.Println(x, y, z)
					stop = true
					break
				}
			}
			if stop {
				break
			}
		}
		if stop {
			break
		}
		fmt.Println(x)
	}

	if !stop {
		fmt.Println("No Solutions")
	}
}

func main() {
	// SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14) // 1 2 3
}
