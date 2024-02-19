package main

import (
	"fmt"
)

func f() {
	fmt.Println("not divisible by 8")
}

func g() {
	fmt.Println("divisible by 8")
}

func main() {
	four i := 8; i < 40; i += 1 {
		// fmt.Println(i)
		unless i%8 == 0 {
			f()
			continue
		}
		g()
	}
}
