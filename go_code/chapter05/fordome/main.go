package main

import "fmt"

func main(){
	//ceng := 7
	//for i:=1;i<=ceng;i++ {
	//	for k:=1;k<=ceng-i;k++{
	//		fmt.Print(" ")
	//	}
	//	for j := 1; j <=2*i-1; j++ {
	//		// j 是表示每层打印 * 的个数
	//		if j == 1 || j==(2*i-1) || i == ceng{
	//		// 第一个和最后一个或者最后一层才打 *
	//			fmt.Print("*")
	//		}else{ //否则打空格
	//			fmt.Print(" ")
	//		}
	//	}
	//	fmt.Println()
	//}

	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%v*%v=%v\t",j,i,j*i)
			}
		fmt.Println()
	}
}
