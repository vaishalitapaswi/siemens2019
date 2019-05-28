package main

import (
	"fmt"
	"time"
)

var count int = 0

func m1() {
	for i := 1; i <= 100; i++ {
		locvar := count
		time.Sleep(10 * time.Millisecond)
		locvar += 1
		count = locvar
	}
	fmt.Println("Current Count = ", count)
}
func main() {
	go m1()
	fmt.Println("Current Count = ", count)
	go m1()
	fmt.Println("Current Count = ", count)
	go m1()
	fmt.Println("Current Count = ", count)
	for {
	}
}
