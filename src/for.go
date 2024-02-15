package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("hello world")
	}

	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	for a := 0; a < 2; a++ {
		b := 0
		for b < 2 {
			fmt.Println(a, b)
			b++
		}
	}
}
