package main

import (
	"fmt"
)

func main() {
	var line1 = "nihao"
	line2 := fmt.Sprintf(":TEST-SVC-SRI:MDN=86%s,TYPE=0;", line1)
	fmt.Println(line2)
}
