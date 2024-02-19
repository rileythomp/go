package main

import "fmt"

func f(i int) {
	fmt.Println("in unless block", i)
}

func main() {
	unless true {
		f(1)
		return
	}
	f(5)
}
