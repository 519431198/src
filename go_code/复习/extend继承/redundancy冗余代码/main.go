package main

import "fmt"

// Pupil 学生考试系统
type Pupil struct {
	Name  string
	Age   int
	Score float64
}

func (p *Pupil) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n", p.Name, p.Age, p.Score)
}

func (p *Pupil) SetScore(score float64) {
	p.Score = score
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

//大学生考试
type Graduate struct {
	Name  string
	Age   int
	Score float64
}

func (p *Graduate) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n", p.Name, p.Age, p.Score)
}

func (p *Graduate) SetScore(score float64) {
	p.Score = score
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试...")
}
func main() {
	var pupil = &Pupil{
		Name: "tom",
		Age:  10,
	}
	pupil.testing()
	pupil.SetScore(90)
	pupil.ShowInfo()

	var graduate = &Graduate{
		Name: "tom",
		Age:  10,
	}
	graduate.testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
}
