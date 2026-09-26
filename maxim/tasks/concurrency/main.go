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
	a := RuneCountAsync("123445")
	// time.Sleep(2 * time.Second)
	fmt.Println(a)
	fmt.Println(RuneCountAsync(""))
	fmt.Println("----7------")
	ch3 := make(chan bool)
	go func() {
		fmt.Println("готово")

		time.Sleep(0 * time.Second)
		ch3 <- true
	}()
	_ = <-ch3
	fmt.Println("конец")
	fmt.Println("----8------")
	n := Rectangle{3, 5}
	ch4 := make(chan int)
	go n.SendArea(ch4)
	fmt.Println(<-ch4)
	fmt.Println("----9------")
	fmt.Println(AnalyzePair(1, 1))
	fmt.Println(AnalyzePair(2, 1))
	// go fmt.Println(AnalyzePair(1, 1))
	fmt.Println("----10------")
	fmt.Println(SuNonNegative([]int{-3, -1, 2, -5}))
	fmt.Println("----11------")
	ch5 := Send3Mess()
	fmt.Println(<-ch5, <-ch5, <-ch5)
	// for v := range ch5 {
	// 	fmt.Println(v)
	// }
}
