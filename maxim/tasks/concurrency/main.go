package main

import (
	"fmt"
	"time"
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
	go func() {
		SendNumber(43, ch2)
	}()
	res1 = <-ch2
	fmt.Println(res1)
	fmt.Println("----3------")
	fmt.Println(Double(2))
	fmt.Println(Double(-2))
	fmt.Println("----4------")
	fmt.Println(Greet("Friend"))
	fmt.Println(Greet(""))
	fmt.Println("----5------")
	fmt.Println(SumAsync([]int{1, 2, 3, 4, -1, -2, -4}))
	fmt.Println(SumAsync([]int{}))
	fmt.Println("----6------")
	fmt.Println(RuneCountAsync("123456"))
	fmt.Println(RuneCountAsync(""))
	fmt.Println("----7------")
	ch3 := make(chan bool)
	go func() {
		fmt.Println("готово")
		time.Sleep(2 * time.Second)
		ch3 <- true
	}()
	_ = <-ch3
	fmt.Println("конец")
}
