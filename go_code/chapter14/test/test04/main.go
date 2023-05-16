package main

import (
	"fmt"
	"os"
)

func PathExits(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
func main() {
	file1Path := "/Users/wangyi/utils/test3.txt"
	//判断文件是否存在
	exits, err := PathExits(file1Path)
	if err == nil {
		fmt.Println(exits)
		return
	} else {
		fmt.Println(exits)
	}

}
