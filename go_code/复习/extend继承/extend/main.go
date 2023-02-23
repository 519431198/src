package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score float64) {
	stu.Score = score
}

// Pupil 学生考试系统
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

// Graduate 大学生考试
type Graduate struct {
	Student
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试...")
}
func main() {
	var pupil = &Pupil{
		Student{
			Name: "tom",
			Age:  10,
		},
	}
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

	var graduate = &Graduate{
		Student{Name: "marry",
			Age: 10,
		},
	}
	graduate.testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
}
