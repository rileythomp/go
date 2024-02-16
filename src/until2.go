package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	until {
		i++
		time.Sleep(1 * time.Second)
		if i % 2 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
		if i > 8 {
			break
		}
	}
}
