package main

import "fmt"

func f() error {
	// return fmt.Errorf("error")
	return nil
}

func saycode(i int) {
	fmt.Println(i)
}


func main() {
	i := 1
	unless err := f(); err == nil {
		i = 3
	}
	saycode(i)
}
