package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student Student) say() string {
	infoStr := fmt.Sprintf("信息是 name=[%v],gender=[%v],age=[%v],id=[%v],score=[%v]",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

func main() {
	stu := Student{
		name:   "mary",
		gender: "man",
		age:    18,
		id:     101,
		score:  98,
	}
	infoStr := stu.say()
	fmt.Println(infoStr)
}
