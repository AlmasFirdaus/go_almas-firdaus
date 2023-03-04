package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Score int
}

type MinMax struct {
	Average float64
}

func (average MinMax) result(min, max Student) {
	fmt.Printf("Average Score %g \n Min Score of Students : %s (%d) \n Max Score of Students : %s (%d)", average.Average, min.Name, min.Score, max.Name, max.Score)
}

func main() {
	students := []Student{
		{
			Name:  "Rizky",
			Score: 80,
		},
		{
			Name:  "Kobar",
			Score: 75,
		},
		{
			Name:  "Ismail",
			Score: 70,
		},
		{
			Name:  "Umam",
			Score: 75,
		},
		{
			Name:  "Yopan",
			Score: 60,
		},
	}

	var indeks = []int{0, 0}
	avg, max := 0, 0
	min := students[0].Score
	for i, v := range students {
		avg += v.Score
		if v.Score <= min {
			min = v.Score
			indeks[0] = i
		}

		if v.Score >= max {
			max = v.Score
			indeks[1] = i
		}
	}

	averageStudent := MinMax{
		Average: float64(avg / len(students)),
	}

	averageStudent.result(students[indeks[0]], students[indeks[1]])
}
