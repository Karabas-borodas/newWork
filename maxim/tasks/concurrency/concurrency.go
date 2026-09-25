package main

import "fmt"

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

// NOTE:4
func Greet(name string) string {
	ch1 := make(chan string)
	go func(name string) {
		ch1 <- fmt.Sprintf("Привет %s", name)
	}(name)
	return <-ch1
}

// NOTE:5
func SumAsync(nums []int) int {
	chan1 := make(chan int)
	summ := 0
	go func([]int) {
		for _, v := range nums {

			summ += v
		}
		chan1 <- summ
	}(nums)
	return <-chan1
}

// NOTE:6
func RuneCountAsync(s string) int {
	r := []rune(s)
	ch1 := make(chan int)
	count := 0
	go func() {
		for range r {
			count++
		}
		ch1 <- count
	}()
	return <-ch1
}

// NOTE:7
// NOTE:8
type Rectangle struct{ Width, Height int }

func (r Rectangle) SendArea(ch chan int) {
	ch <- (r.Height * r.Width)
}

// NOTE:9
type PairInfo struct {
	Sum   int
	Equal bool
}

func AnalyzePair(a, b int) PairInfo {
	Par := PairInfo{}
	ch1 := make(chan PairInfo)
	go func() {
		Par.Sum = a + b
		if a == b {
			Par.Equal = true
		} else {
			Par.Equal = false
		}
		ch1 <- Par
	}()
	return <-ch1
}

// NOTE:10
func SuNonNegative(nums []int) int {
	summ := 0
	ch1 := make(chan int)
	go func() {
		for _, v := range nums {
			if v > 0 {
				summ += v
			}
		}
		ch1 <- summ
	}()
	return <-ch1
}
