package main

type Stu struct {
	Name  string
	Score int
}

type StuScore []Stu

func (stu StuScore) Len() int {
	return len(stu)
}
func (stu StuScore) Less(i, j int) bool {
	return stu[i].Score < stu[j].Score
}

func (stu StuScore) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {

}
