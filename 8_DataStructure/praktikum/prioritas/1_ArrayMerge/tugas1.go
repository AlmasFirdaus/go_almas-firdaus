package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string{
	arrayGroup := make([]string, 1)
	copy(arrayGroup, arrayA)

	for _, v := range arrayB {
		exist := false
		for _, x:= range arrayA {
			if v == x {
				exist = true
			}
		}
		
		if exist != true {
			arrayGroup = append(arrayGroup, v)
		}
	}

	return arrayGroup
}

func main()  {
	// Test Case
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve","geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"},[]string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// [""]
}