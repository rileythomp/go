package main 

import "fmt"

func useuntil() {
	until i := 4; i == 0; i-- {
		sayhi()
	}
}

func sayhi() {
	fmt.Println("hello world")
}

