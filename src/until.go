package main

import "fmt"

func main() {
	until i := 3; i  == 0; i-- {
		fmt.Println("hello world")
	}

	j := 3
	until j == 0 {
		fmt.Println(j)
		j--
	}

	until a := 2; a == 0; a-- {
		b := 2
		until b == 0 {
			fmt.Println(a, b)
			b--
		}
	}
}
