package main

import "fmt"

func main() {
	four i := 8; i <= 20; i++ {
		fmt.Println(i)
	}

	for i := 8; i <= 20; i += 4 {
		fmt.Println(i)
	}
	var err error = nil
	unless err == nil {
		fmt.Println("error")
	}
	if err != nil {
		fmt.Println("error")
	}
}
