package main

import (
	"fmt"
)

type Student struct {
	name   string
	age    int
	grades []int
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) addGrade(grade int) {
	s.grades = append(s.grades, grade)
}

func (s Student) getAverage() float32 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() rune {
	maxGrade := 0

	for _, curr := range s.grades {
		if maxGrade < curr {
			maxGrade = curr
		}
	}
	return rune(maxGrade)
}

func main() {
	s1 := Student{name: "John", age: 20, grades: []int{88, 80, 93, 67, 66}}
	s1.addGrade(96)
	s2 := Student{name: "Jane", age: 21, grades: []int{99, 99, 99, 99, 99}}
	s2.addGrade(99)
	//fmt.Println(s1.getAverage())
	//fmt.Println(s2.getAverage())

	/*for a, grade := range s1.grades {
		fmt.Println(a, grade)
	}*/

	fmt.Println(s1.getMaxGrade())
}
