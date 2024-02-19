package main

import "fmt"

func odd() {
	fmt.Println("odd")
}

func even() {
	fmt.Println("even")
}

func main() {
	until i := 10; i == 3; i-- {
		unless i%2 == 0 {
			odd()
			continue
		}
		even()
	}
}
