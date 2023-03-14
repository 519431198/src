package main

import (
	"fmt"
	"time"
)

func main() {
	lastUpdate := time.Now()
	for {
		if time.Since(lastUpdate).Seconds() > 2 {
			fmt.Println("nihao")
			lastUpdate = time.Now()
		}
	}

}
