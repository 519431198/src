package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student信息:\n  name:[%v] gender:[%v] age:[%v] id:[%v] score:[%v]",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

func main() {
	var str = Student{
		name:   "tom",
		gender: "male",
		age:    17,
		id:     775,
		score:  98,
	}
	fmt.Println(str.say())
}
