package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var intSlice = []int{3, 57, 4, 33, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对结构体切片进行排序
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄: %d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Println(v)
	}
}

// Hero 声明一个结构体
type Hero struct {
	Name string
	Age  int
}

// HeroSlice 声明一个结构体切片
type HeroSlice []Hero

// Len 实现 Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

//Less 决定你使用什么标准进行排序
//例如按照 Hero 的 Age 从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
