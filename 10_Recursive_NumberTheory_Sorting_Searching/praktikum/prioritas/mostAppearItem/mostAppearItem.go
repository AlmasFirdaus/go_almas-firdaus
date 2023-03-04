package main

import "fmt"

func mostAppearItem(items []string) map[string]int {
	mapping := map[string]int{}
	// result := []string{}

	for _, item := range items {
		_, exist := mapping[item]
		if !exist {
			mapping[item] = 1
		} else {
			mapping[item] += 1
		}
	}

	// for key, value := range mapping {

	// 	if len(result) == 0 {
	// 		str := key + " -> " + string(value)
	// 		result = append(result, value)
	// 	}
	// }

	return mapping
}

func main() {
	fmt.Println(mostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(mostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(mostAppearItem([]string{"footbal", "basketball", "tenis"}))

}
