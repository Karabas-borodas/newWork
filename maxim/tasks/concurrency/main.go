package main

import (
	"fmt"
)

func main() {
	// ch1 := make(chan string)
	fmt.Println("----1------")
	ch1 := firstMassage()
	mess := <-ch1
	fmt.Println(mess)

	fmt.Println("----2------")
	ch2 := make(chan int)
	go func() {
		SendNumber(42, ch2)
	}()
	res1 := <-ch2
	fmt.Println(res1)
}
