package main

import (
	"fmt"
	"sync"
	// "time"
)

// NOTE:1
// first massage
func firstMassage() chan string {
	var ch = make(chan string)
	go func() {
		ch <- "hello"
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
			n *= 2
			ch1 <- n
		} else {
			n *= 2 * -1
			ch1 <- n
		}
	}(n)
	return <-ch1
}

// NOTE:4
func gr(str string, ch chan<- string) {
	ch <- fmt.Sprintf("Привет %s", str)
}

func Greet(name string) string {
	ch1 := make(chan string)
	go gr(name, ch1)
	return <-ch1
}

// NOTE:5
func SumAsync(nums []int) int {
	chan1 := make(chan int)
	go func(n []int) {
		summ := 0
		for _, v := range n {

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
func SumNonNegative(nums []int) int {
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

// NOTE:11
func Send3Mess() chan int {
	// var wg sync.WaitGroup
	ch1 := make(chan int)
	go func() {
		// wg.Add(3)
		ch1 <- 50
		ch1 <- 20
		ch1 <- 30
		// for i := 0; i < 1000; i++ {
		// 	ch1 <- i
		// }
		// wg.Done()
		// wg.Done()
		// wg.Done()
		close(ch1)
	}()
	// wg.Wait()
	return ch1
}

// NOTE:12
func CopyThroughChannel(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	ch1 := make(chan int)
	go func(nums []int) {
		for _, v := range nums {
			ch1 <- v
		}
	}(nums)
	slice := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		slice = append(slice, <-ch1)
	}
	return slice
}

// NOTE:15
func TwoResults(a, b int) []int {
	ch := make(chan int)
	go func(a int) {
		ch <- a
	}(a)

	go func(b int) {
		ch <- b
	}(b)
	slice := []int{}
	for i := 0; i < 2; i++ {
		slice = append(slice, <-ch)
	}
	return slice
}

// NOTE:17
func SumHalves(nums []int) int {
	length := len(nums) / 2
	ch1 := make(chan int, 2)
	go func(l int, nums []int) {
		summ1 := 0
		for i := 0; i < l; i++ {
			summ1 += nums[i]
		}
		ch1 <- summ1
	}(length, nums)
	go func(l int, nums []int) {
		summ1 := 0
		for i := len(nums) - 1; i >= l; i-- {
			summ1 += nums[i]
		}
		ch1 <- summ1
	}(length, nums)
	summ := 0
	for i := 0; i < 2; i++ {
		summ += <-ch1
	}
	return summ
}

// NOTE:18
func Echo(x int) int {
	ch1 := make(chan int)
	go func(x int) {
		ch1 <- x
	}(x)
	return <-ch1
}

// NOTE:20
func CountSteps(steps []int) int {
	// counter := 0
	// var mutex sync.Mutex
	// ch1 := make(chan int)
	// for i := 0; i < len(steps); i++ {
	// 	go func(x int, m *sync.Mutex) {
	// 		// mutex.Lock()
	// 		// counter += x
	// 		// mutex.Unlock()
	// 		ch1 <- x
	// 	}(steps[i], &mutex)
	// }
	// for i := 0; i < len(steps); i++ {
	// 	counter += <-ch1
	// }
	// return counter

	counter := 0
	var mutex sync.Mutex
	// ch1 := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < len(steps); i++ {
		wg.Add(1)
		go func(x int, m *sync.Mutex, counter *int, wg *sync.WaitGroup) {
			mutex.Lock()
			*counter += x
			mutex.Unlock()

			wg.Done()
		}(steps[i], &mutex, &counter, &wg)
	}
	// for i := 0; i < len(steps); i++ {
	// 	counter += <-ch1
	// }
	wg.Wait()
	// time.Sleep(5 * time.Microsecond)
	return counter
}

// NOTE:24
func BufferedSequence(n int) []int {
	if n <= 0 || n > 100 {
		return []int{}
	}
	ch1 := make(chan int, 2)
	go func(n int) {
		for i := 1; i < n+1; i++ {
			ch1 <- i
		}
	}(n)
	fmt.Println(cap(ch1))
	slice := make([]int, n)
	fmt.Println(len(ch1))
	for i := 0; i < n; i++ {
		slice[i] = <-ch1
	}
	return slice
}

// NOTE:26
func StartSquare(n int) <-chan int {
	ch1 := make(chan int, 1)
	go func() {
		ch1 <- n * n
	}()
	return ch1
}
