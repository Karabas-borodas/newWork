package main

import ()

// NOTE:1
// first massage
func firstMassage() chan string {
	var ch = make(chan string)
	go func() {
		ch <- "hell"
		// close(ch)
	}()
	return ch
}

// NOTE:2
func SendNumber(n int, ch chan int) {
	ch <- n
}

// NOTE:3
func Double(n int) int {
	ch1 := make(chan int)
	go func(n int) {
		if n > 0 {
			n *= n
			ch1 <- n
		} else {
			n *= n * -1
			ch1 <- n
		}
	}(n)

	return <-ch1
}
