package main

import "fmt"

type Student struct {
	name string
	solvedProblems int
	scoreForOneTask float64
	passingScore float64
}

func (s Student) IsExcellentStudent() bool {
	if float64(s.solvedProblems) * s.scoreForOneTask >= s.passingScore {
		return true
	}
	return false
}

func main() {
	student1 := Student{
		name: "masha",
		solvedProblems: 12,
		scoreForOneTask: 4,
		passingScore: 30,
	}

	student2 := Student{
		name: "misha",
		solvedProblems: 5,
		scoreForOneTask: 4,
		passingScore: 30,
	}

	fmt.Println(student1.IsExcellentStudent()) // true
	fmt.Println(student2.IsExcellentStudent()) // false
}
