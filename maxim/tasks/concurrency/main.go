package main

import (
	"fmt"
)

func main() {
	// ch1 := make(chan string)
	ch1 := firstMassage()
	mess := <-ch1
	fmt.Println(mess)
}
