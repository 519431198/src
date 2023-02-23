package main

import "fmt"

func main(){
	var str string = "hello,world!上海"
	//str2 := []rune(str)
	//for i := 0; i < len(str2);{
	//	fmt.Printf("%c %v\n", str2[i],i)
	//	i++
	//}

	fmt.Println()
	for index,val :=range str {
		fmt.Printf("index=%d,val=%c\n",index,val)
	}
}
