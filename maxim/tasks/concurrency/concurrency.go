package main

import (
	"fmt"
	"sync"
	"unicode"
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

// NOTE:28

func SumWithBuffer(nums []int) int {
	ch1 := make(chan int, 1)
	if len(nums) == 0 {
		return 0
	}
	go func(nums []int) {

		for i := 0; i < len(nums); i++ {
			ch1 <- nums[i]
		}
	}(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		count += <-ch1
	}
	return count
}

// NOTE:30
func BufferedPair(a, b string) map[string]int {
	if a == "" || b == "" {
		return nil
	}
	ch1 := make(chan string, 2)
	go func(a string) {
		ch1 <- a
	}(a)
	go func(a string) {
		ch1 <- a
	}(b)

	ma := make(map[string]int)
	for i := 0; i < 2; i++ {
		ma[<-ch1]++
	}

	return ma

}

// NOTE:32
func ReceiveN(in <-chan int, n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		// result = append(result, <-in)
		result[i] = <-in
	}
	return result
}

// NOTE:33
func SendAndClose(value int, out chan<- int) {
	out <- value
	close(out)
}

// NOTE:35
func CollectRange(in <-chan int) []int {
	out := []int{}
	for v := range in {
		out = append(out, v)
	}
	return out
}

// NOTE:37
func EmptyInts() <-chan int {
	ch1 := make(chan int)
	// go func() {
	// ch1 <- 3
	close(ch1)
	// }()
	return ch1
}

// NOTE:40

func As() <-chan int {
	fmt.Println("обработка завершена")
	ch2 := make(chan int)
	go func() {

		ch2 <- 1
		close(ch2)
	}()
	return ch2
}

// NOTE:41 -1
func SumStream(in <-chan int) int {
	var n int = 0
	for v := range in {
		n += v
	}
	return n
}

// NOTE:43 -2
func WordCounts(in <-chan string) map[string]int {
	ma := make(map[string]int)
	for v := range in {
		ma[v]++
	}
	return ma
}

// NOTE:44 -3
func JoinStream(in <-chan string, sep string) string {
	var str string = ""
	for v := range in {
		str += v
		str += sep
	}
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}

// NOTE:46 -4
func LongestWord(in <-chan string) (string, int) {
	var length int
	var word []rune
	var index int
	var indexstring int
	for v := range in {
		va := []rune(v)
		if length < len(va) {
			length = len(va)
			word = va
			indexstring = index
		}
		index++
	}
	exit := string(word)
	return exit, indexstring

}

// NOTE:47 -1.1
type Run struct {
	Value, Count int
}

func Runs(in <-chan int) []Run {
	var r []Run
	for v := range in {
		if len(r) == 0 {
			r = append(r, Run{Value: v, Count: 1})
			continue
		}

		if v == r[len(r)-1].Value {
			r[len(r)-1].Count += 1
		} else {
			var va Run
			va.Value = v
			va.Count = 1
			r = append(r, va)
		}
	}

	return r
}

// NOTE:50 -1.2
func TripleSums(in <-chan int) []int {
	// var insert []int
	// for v := range in {
	// 	insert = append(insert, v)
	// }
	// var out []int
	// var summ int
	// for i := 0; i < len(insert)-2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		summ += insert[i+j]
	// 	}
	// 	out = append(out, summ)
	// 	summ = 0
	// }
	// return out
	//NOTE: нейронка

	var insert [3]int
	var out []int
	count := 0
	for v := range in {
		insert[count%3] = v
		count++
		if count >= 3 {
			out = append(out, insert[0]+insert[1]+insert[2])
		}
	}
	return out
}

// NOTE:51 -1.3

type Entry struct {
	Key   string
	Value int
}

func LastValues(in <-chan Entry) map[string]int {
	m := make(map[string]int)
	for v := range in {
		m[v.Key] = v.Value
	}
	return m
}

// NOTE:53 -1.4
func Balanced(in <-chan rune) bool {
	var ret bool = true
	count := 0
	for v := range in {
		if v == rune('(') {
			count++

		}
		if v == rune(')') {
			count--
		}
	}
	if count != 0 {
		ret = false
	}
	return ret
}

// NOTE:54 -1.5
func FirstUnique(in <-chan string) (string, bool) {
	var sl []string
	var m = make(map[string]int)
	for v := range in {
		sl = append(sl, v)
		m[v] += 1
	}
	for _, v := range sl {
		if c := m[v]; c == 1 {
			return v, true
		}
	}
	return "", false
}

// NOTE:56 -1.6
func RowSums(rows [][]int) []int {
	//NOTE:решенеи на мютекасах
	// l := len(rows)
	// var summ = make([]int, l)
	// var wg sync.WaitGroup
	// var mutex sync.Mutex
	// wg.Add(len(rows))
	// for i := 0; i < len(rows); i++ {
	// 	go func(i int, summ []int, mutex *sync.Mutex) {
	// 		defer wg.Done()
	// 		//WARNING: знаю что раком ставлю всю скорость подсчета из за блокируиз операций
	// 		mutex.Lock()
	// 		for j := 0; j < len(rows[i]); j++ {
	// 			summ[i] += rows[i][j]
	// 		}
	// 		mutex.Unlock()
	//
	// 	}(i, summ, &mutex)
	// }
	// wg.Wait()
	// return summ
	//NOTE: решение на каналах
	l := len(rows)
	// 1. Обязательно создаем срез нужной длины через make!
	summ := make([]int, l)

	result := make(chan struct {
		index int
		summ  int
	}, l) // Буферизированный канал

	for i := 0; i < l; i++ {
		go func(i int) {
			// 2. Локальная переменная для суммы строки (называем rowSum, чтобы не путать)
			rowSum := 0

			// 3. ВАЖНО: используем _, v для перебора ЗНАЧЕНИЙ, а не индексов!
			for _, v := range rows[i] {
				rowSum += v
			}

			// Отправляем результат в канал
			result <- struct {
				index int
				summ  int
			}{index: i, summ: rowSum}
		}(i)
	}

	// 4. Собираем результаты
	for i := 0; i < l; i++ {
		res := <-result
		// 5. Записываем по правильному индексу
		summ[res.index] = res.summ
	}

	return summ
}

// NOTE:57 -1.7

func OrderedSquares(nums []int) []int {
	// //NOTE:мое решенеи
	// type Str struct {
	// 	index int
	// 	value int
	// }
	// var wg sync.WaitGroup
	// ret := make([]int, len(nums))
	// chans := make(chan Str)
	// wg.Add(len(nums))
	// for i := 0; i < len(nums); i++ {
	// 	go func(i int) {
	// 		defer wg.done()
	// 		chans <- str{index: i, value: nums[i] * nums[i]}
	// 	}(i)
	// }
	// go func() {
	// 	wg.Wait()
	// 	close(chans)
	// }()
	// for v := range chans {
	// 	ret[v.index] = v.value
	// }
	// return ret
	//NOTE: через fan-out
	type Str struct {
		index int
		value int
	}
	ret := make([]int, len(nums))

	// 1. Канал для данных
	chans := make(chan Str)

	// 2. Канал для сигналов о завершении горутин
	done := make(chan struct{}, len(nums))

	for i := 0; i < len(nums); i++ {
		go func(i int) {
			// Отправляем результат
			chans <- Str{index: i, value: nums[i] * nums[i]}
			// Отправляем сигнал, что эта горутина закончила
			done <- struct{}{}
		}(i)
	}

	// 3. Запускаем фоновую горутину, которая ждет все сигналы и закрывает канал
	go func() {
		for i := 0; i < len(nums); i++ {
			<-done // Ждем ровно len(nums) сигналов
		}
		close(chans) // Теперь безопасно закрыть
	}()

	// 4. Главный поток читает данные
	for v := range chans {
		ret[v.index] = v.value
	}
	return ret
}

// NOTE:59 -1.8

func FirstArrival(values []string) (string, bool) {
	ch := make(chan string)
	for _, v := range values {
		go func(v string) {
			ch <- v
		}(v)
	}

	go func() {

		for i := 0; i < len(values); i++ {
			<-ch
		}
	}()
	return <-ch, true
}

// NOTE:60 -1.9
func AllNonEmpty(values []string) bool {
	// var ok = true
	// ch := make(chan bool)
	// ch2 := make(chan struct{})
	// for i := 0; i < len(values); i++ {
	// 	go func(val string) {
	// 		if val == "" {
	// 			ch <- false
	// 		} else {
	// 			ch <- true
	// 		}
	// 		ch2 <- struct{}{}
	// 	}(values[i])
	// }
	// go func() {
	//
	// 	for i := 0; i < len(values); i++ {
	// 		<-ch2
	// 	}
	// 	close(ch)
	// }()
	//
	// for i := 0; i < len(values); i++ {
	// 	c := <-ch
	// 	if c == false {
	// 		ok = false
	// 	}
	// }
	// return ok
	//NOTE: 1 канал
	if len(values) == 0 {
		return true
	}

	ch := make(chan bool)

	for i := 0; i < len(values); i++ {
		go func(val string) {
			ch <- (val != "") // Отправляем true, если не пустая
		}(values[i])
	}

	ok := true
	for i := 0; i < len(values); i++ {
		if !<-ch { // Если пришел false
			ok = false
		}
	}
	return ok
}

// NOTE:61 -1.10
type Metric struct {
	Name  string
	Value int
}

func TextMetrics(s string) map[string]int {
	ma := make(map[string]int)
	ch := make(chan int)
	countR := 0
	countA := 0
	go func(count int) {
		r := []rune(s)
		count = len(r)
		ch <- count
		fmt.Println("countr -- ", count)
	}(countR)

	go func(count int) {
		for _, sp := range s {
			if unicode.IsSpace(sp) {
				count++
			}
		}
		ch <- count
	}(countA)
	countA = <-ch
	countR = <-ch
	ma["runes"] = countR
	ma["spaces"] = countA
	return ma
}
