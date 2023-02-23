package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allchan chan interface{}
	allchan = make(chan interface{}, 10)
	cat1 := Cat{Name: "Tom", Age: 10}
	cat2 := Cat{Name: "Rayna", Age: 18}
	allchan <- cat1
	allchan <- cat2
	allchan <- 10
	allchan <- "jack"
	close(allchan)
	cat11 := <-allchan
	fmt.Printf("cat11 type is %T", cat11)
	//类型断言
	a := cat11.(Cat)
	fmt.Printf("cat11.Name=%v", a.Name)
}
