package main

import (
	// "error"
	"fmt"
)

// NOTE: 4
func CountEven(s []int) int {

	count := 0
	// var summ float32
	for _, v := range s {
		if v%2 == 0 {
			count += 1
		}
	}
	return count
}

// NOTE: 7
func IndexOf(s []int, target int) int {
	index := -1
	for i, v := range s {
		if v == target {
			index = i
			break
		}
	}
	return index
}

// NOTE: 12
// XXX:что значит не меняя слайс?
func ReverseRange(s []int, left, right int) error {
	// var first = 0
	// var last = 0
	if right > len(s) || left < 0 || left > right {
		return fmt.Errorf("out of range")
	}
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--

	}
	// for i := left; i < left+(right-left+1)/2; i++ {
	// 	j := left + right - i
	// 	fmt.Println("----------------------")
	// 	fmt.Println(i)
	// 	fmt.Println(left)
	// 	fmt.Println()
	// 	fmt.Println("----------------------")
	// 	s[i], s[j] = s[j], s[i]
	// 	// first = s[i]
	// 	// last = s[right-i]
	// 	// s[right-i] = first
	// 	// s[i] = last
	// }
	fmt.Println(s)
	return nil
}

// NOTE:13
func RemoveAt(s []int, i int) ([]int, error) {
	if i < 0 || i > len(s)-1 || len(s) == 0 {
		return nil, fmt.Errorf("out of range")
	}
	s = append(s[:i], s[i+1:]...)
	return s, nil
}

// NOTE:14
// XXX:чем оличается от 13 если у нас поиск по индексу O(n)?
func RemoveAtFast(s []int, i int) ([]int, error) {
	if i < 0 || i > len(s)-1 || len(s) == 0 {
		return nil, fmt.Errorf("out of range")
	}
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	return s, nil
}

// NOTE:18
func MoveZerosToEnd(s []int) {
	// sl := make([]int, len(s))
	// copy(sl, s)
	//NOTE:укорачиваю кописю так как ранее удалял 1 элемнт
	// s = s[:len(s)-1]
	lenSlice := len(s)
	// fmt.Println(lenSlice)
	// variable := 0
	for i := lenSlice - 1; i >= 0; i-- {
		if s[i] == 0 {
			for j := i; j < lenSlice-1; j++ {
				s[j] = s[j+1]
			}
			s[lenSlice-1] = 0
		}
	}
	fmt.Println(s)
}

// NOTE:18-2
// FIX: разобраться
func MoveZerosToEnd2(s []int) {
	var i = 0

	for _, v := range s {
		if v != 0 {
			s[i] = v
			i++
		}
	}
	for i < len(s) {
		s[i] = 0
		i++
	}
	fmt.Println(s)
}

// NOTE:27
// XXX:backing array WTF?
func Concat(a, b []int) []int {
	slice := make([]int, len(a)+len(b))

	copy(slice, a)
	copy(slice[len(a):], b)
	return slice
}

// NOTE:29 сделать
func Unique(s []int) []int {
	//NOTE:my logic
	// slice := make([]int, len(s))
	// copy(slice, s)
	// ma := make(map[int]bool)
	// count := 0
	// for _, v := range slice {
	// 	if _, ok := ma[v]; !ok {
	//
	// 		ma[v] = true
	// 		slice[count] = v
	// 		count++
	// 	}
	// }
	// slice = slice[:count]
	// return slice
	//NOTE:neyro
	slice := make([]int, 0)
	ma := make(map[int]struct{})
	for _, v := range s {
		if _, ok := ma[v]; !ok {
			slice = append(slice, v)
			ma[v] = struct{}{}
		}
	}
	return slice
}

// NOTE:47 сделать
func ChunkViews(s []int, size int) ([][]int, error) {
	//NOTE:мое решение
	if len(s) == 0 {
		return [][]int{}, nil
	}
	if size < 0 {
		return nil, fmt.Errorf("size cant be lou %d", size)
	}

	countChank := (len(s) + size - 1) / size

	slice := make([][]int, countChank)
	lastIndex := size
	lavelSlice := 1
	// var sl []int
	// slice[0] = sl
	for i := 0; i < len(s); i += size {
		if lastIndex > len(s) {
			lastIndex = len(s)
		}
		slice[lavelSlice-1] = s[i:lastIndex]
		lavelSlice++
		lastIndex = lavelSlice * size
	}
	return slice, nil
	//NOTE:нейронка

	// if len(s) == 0 {
	// 	return [][]int{}, nil
	// }
	// if size < 0 {
	// 	return nil, fmt.Errorf("size cant be lou %d", size)
	// }
	// // countChank := (len(s) + size - 1) / size
	// slice := make([][]int, 0)
	// for i := 0; i < len(s); i += size {
	// 	end := i + size
	// 	if end > len(s) {
	// 		end = len(s)
	// 	}
	// 	slice = append(slice, s[i:end])
	//
	// }
	// slice[0][0] = 9
	// return slice, nil

}

// NOTE:49 сделать
// BUG: несовсем понятно что именлсь ввиду (уточнить исплавления)
func SplitIntoKViews(s []int, k int) ([][]int, error) {
	if len(s) < 1 {
		return nil, fmt.Errorf("aaaaa all bad")
	}
	if k < 1 {
		return nil, fmt.Errorf("aaaaa all bad")
	}
	// if k > len(s) {
	// 	k = len(s)
	// }
	// ka := k
	chankLangth := (len(s)) / k
	chankLitlBit := len(s) % k
	slice := make([][]int, k)
	// sliceCount := 0
	// for i := 0; i < k; i++ {
	// 	slice[i] = []int{}
	// }

	start := 0
	for i := 0; i < k; i++ {
		size := chankLangth
		if i < chankLitlBit {
			size++
		}
		chank := s[start : size+start]
		slice[i] = chank
		start += size
	}
	// fmt.Println(len(slice))
	// if len(slice) < ka {
	// 	for len(slice) <= k+1 {
	// 		sl := make([]int, 0)
	// 		slice = append(slice, sl)
	// 	}
	// }
	return slice, nil
}

// TODO:: после созвона 03.09.26
// NOTE:50
// Верни все подряд идущие окна длины size как subslice исходного backing array. Для size > len(s)
// верни пустой результат.
func SlidingWindowViews(s []int, size int) ([][]int, error) {
	if size < 0 {
		return nil, fmt.Errorf("size <0")
	}
	if size > len(s) {
		return nil, fmt.Errorf("size >%d", len(s))
	}
	var slice [][]int
	// countChank := len(s) / size
	// lastChankLength := len(s) - countChank*size
	i := 0
	lastIndex := size
	for lastIndex <= len(s) {
		chank := s[i:lastIndex]
		slice = append(slice, chank)
		i++
		lastIndex++
	}
	// if lastChankLength > 0 {
	// 	chank := s[len(s)-lastChankLength : len(s)]
	// 	slice = append(slice, chank)
	// }
	// slice[0][0] = -9
	// fmt.Println(s)
	return slice, nil
}

// NOTE:52
func SplitOnValue(s []int, sep int) [][]int {
	slice := make([][]int, 0)
	last := 0
	first := 0
	for i, v := range s {
		if v == sep && i != 0 && i != len(s) {

			last = i
			slice = append(slice, s[first:last])
			// slice = append(s[first:last])
			first = i + 1
		}
		if v == sep && i == 0 {
			first = 1
		}
	}
	if first != len(s) {

		slice = append(slice, s[first:])
	}
	return slice
}

// NOTE:54
func ExpandToCap(s []int) []int {
	fmt.Println(cap(s))
	// sl := make([]int, cap(s))
	if s == nil {

		return nil
	}
	sl := s[:len(s)]
	return sl
}

// NOTE:55
// Верни slice view с теми же элементами, но cap результата должен быть ровно len(s). Новый backing
// array создавать нельзя.
func ClipCapacity(s []int) []int {
	if s == nil {
		return nil
	}
	sa := s[0:len(s):len(s)]
	return sa
}

// NOTE:56
// Верни первые n элементов как view, причём cap результата должен быть ровно n, чтобы
// последующий append не мог затереть элементы s после n.
func TakeWithLimitedCap(s []int, n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("n <0")
	}
	if n > len(s) {
		return nil, fmt.Errorf("n >%d", len(s))
	}

	sl := s[:n:n]
	return sl, nil
}

// NOTE:57
// Измени длину до n. При увеличении используй существующий cap без аллокации, если его
// хватает; новые элементы должны быть нулевыми. При уменьшении верни reslice.
func Resize(s []int, n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("length<0")
	}
	if n < len(s) {
		slice := make([]int, n, n)
		slice = s[:n]
		return slice, nil
	}
	sl := s[:len(s)]
	for range n - len(s) {
		sl = append(sl, 0)
	}
	return sl, nil
}

// NOTE:58
// Гарантируй cap >= minCap, сохранив len и значения. Если cap уже хватает, верни slice на том же
// backing array; иначе выдели новый.
func EnsureCapacity(s []int, minCap int) ([]int, error) {
	if minCap < 0 {
		return nil, fmt.Errorf("alarma min<0")
	}
	if cap(s) >= minCap {
		return s, nil
	}
	m := make([]int, len(s), minCap)
	copy(m, s)
	return m, nil

}

// NOTE:60
// Верни результат append, который гарантированно не alias-ит исходный backing array s, даже если
// в s было достаточно свободного cap.
func AppendCopy(s []int, values ...int) []int {
	sl := make([]int, len(s))
	copy(sl, s)
	// for range values {
	sl = append(sl, values...)
	// }
	return sl
}

// NOTE:61
// Скопируй участок [left:right) начиная с индекса dst внутри того же слайса. Корректно обрабатывай
// перекрывающиеся диапазоны и проверяй границы.
func CopyWithin(s []int, dst, left, right int) error {
	// if left > right || left < 0 || right > len(s) || dst < 0 || dst > len(s) || (dst+(right-left)) > len(s) {
	// 	return fmt.Errorf("ou shet bab deta")
	// }
	// // fmt.Printf("%d --dst %d --left %d right %d", len(s), dst, left, right)
	// sl := make([]int, right-left)
	// length := 0
	// if dst+(left-right) > len(s) {
	// 	length = len(s) - dst
	// } else {
	// 	length = right - left
	// }
	// // fmt.Println()
	// // fmt.Println(length)
	// copy(sl, s[left:length+left])
	// count := 0
	// // fmt.Println(sl)
	// // fmt.Println(sl)
	// for i := dst; i < dst+length; i++ {
	// 	s[i] = sl[count]
	// 	count++
	// }
	// return nil
	//NOTE:нейронка
	if left > right || left < 0 || right > len(s) {
		return fmt.Errorf("out of range")
	}
	n := right - left
	if dst < 0 || dst+n > len(s) {
		return fmt.Errorf("out of range sdt")
	}

	if dst < left {
		for i := 0; i < n; i++ {
			s[dst+i] = s[left+i]
		}
	} else {
		for i := n - 1; i >= 0; i-- {
			s[dst+i] = s[left+i]
		}
	}
	return nil
}

// NOTE:62
// Реализуй вставку values через увеличение длины и copy. values может alias-ить s; функция должна
// оставаться корректной.
func InsertWithCopy(s []int, i int, values []int) ([]int, error) {
	if i < 0 || i > len(s) {
		return nil, fmt.Errorf(" out of range insert place")
	}
	fmt.Println(cap(s))
	// if len(s) < len(values)+cap(s) {
	sl := make([]int, len(s)+len(values))
	copy(sl, s)
	copy(sl[i:], values)
	copy(sl[i+len(values):], s[i:])
	return sl, nil
	// } else {

	// }
}

// return nil, nil
// }
// NOTE:63
// Удаляй [left:right) in-place, сохраняя порядок, а освободившийся хвост backing array обнуляй через
// clear, чтобы не удерживать ссылки на удалённые объекты.
func DeleteAndClear[T any](s []T, left, right int) ([]T, error) {
	if left < 0 || left > right || right < 0 || left > len(s) || right > len(s) {
		return nil, fmt.Errorf("alarma out of range ")
	}
	count := right - left
	copy(s[left:], s[right:])
	clear(s[len(s)-count:])
	return s[:len(s)-count], nil
}

//NOTE:65 сделать

func main() {
	// slice := []float32{2.14, 15, 33.2}
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 9, 8}
	slice2 := []int{1, 2, 3, 4}
	slice3 := []int{1, 2}
	slice4 := []int{1}
	// var summ float32
	fmt.Println("7")
	fmt.Println(IndexOf(slice, 4))
	fmt.Println("-----12")
	fmt.Println(ReverseRange(slice2, 0, 3))
	fmt.Println(ReverseRange(slice, 0, 3))
	fmt.Println(ReverseRange(slice4, 0, 0))
	fmt.Println("13")
	fmt.Println(RemoveAt(slice3, 2))
	fmt.Println("14")
	fmt.Println(RemoveAtFast(slice2, 2))
	// fmt.Println(RemoveAtFast(slice2, 8))
	fmt.Println("18")
	MoveZerosToEnd(slice)
	MoveZerosToEnd([]int{0, 1, 0, 3, 0, 5, 8})
	MoveZerosToEnd([]int{0, 1, 0, 3, 0, 5, 8})
	fmt.Println("18-2")
	(MoveZerosToEnd2([]int{0, 1, 0, 3, 0, 5, 8}))
	fmt.Println("27")
	fmt.Println(Concat(slice2, slice3))
	fmt.Println("29")
	fmt.Println(Unique([]int{1, 1, 3, 4, 3, 3, 4, 4, 5, 5}))
	fmt.Println("47")
	fmt.Println(ChunkViews([]int{1, 1, 3, 4}, 7))
	fmt.Println(ChunkViews([]int{1, 1, 3, 4}, 4))
	fmt.Println(ChunkViews([]int{1, 4, 2, 3}, 3))
	fmt.Println(ChunkViews([]int{1, 4, 2, 3}, 2))
	fmt.Println(ChunkViews([]int{1, 4, 2, 3, 4, 5}, 2))
	fmt.Println("49")
	fmt.Println(SplitIntoKViews([]int{1, 4, 2, 3}, 3))
	fmt.Println(SplitIntoKViews([]int{1, 2}, 4))
	fmt.Println(SplitIntoKViews([]int{}, 4))
	fmt.Println(SplitIntoKViews([]int{}, -2))
	// TODO:: после созвона 03.09.26
	fmt.Println("50")
	slice1 := []int{1, 2, 3, 4}
	slice5, _ := SlidingWindowViews(slice1, 2)
	fmt.Println(slice5)
	slice5[0][0] = 8
	fmt.Println(slice1)
	fmt.Println(SlidingWindowViews([]int{1, 2, 3, 4}, 3))
	fmt.Println("52")
	fmt.Println(SplitOnValue([]int{1, 2, 3, 4}, 3))
	fmt.Println(SplitOnValue([]int{1, 2, 3, 4, 3, 6}, 3))
	fmt.Println(SplitOnValue([]int{1, 2, 3, 4}, 1))
	fmt.Println(SplitOnValue([]int{1, 2, 3, 4}, 4))
	fmt.Println("54")
	slice1 = []int{}
	slice1 = append(make([]int, 0, 18), 1, 2, 3, 4, 5, 6)

	slice2 = []int{}
	slice2 = ExpandToCap(slice1)
	slice2[3] = 19
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("len slice1--", len(slice1))
	fmt.Println("cap slice1--", cap(slice1))
	fmt.Println("len slice2--", len(slice2))
	fmt.Println("cap slice2--", cap(slice2))
	q := ExpandToCap(nil)
	if q == nil {
		fmt.Println("q==nil")
	}
	fmt.Println("55")
	slice1 = []int{}
	slice1 = append(make([]int, 0, 18), 1, 2, 3, 4, 5, 6)

	slice2 = []int{}
	slice2 = ClipCapacity(slice1)
	slice2[3] = 19
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("len slice1--", len(slice1))
	fmt.Println("cap slice1--", cap(slice1))
	fmt.Println("len slice2--", len(slice2))
	fmt.Println("cap slice2--", cap(slice2))
	fmt.Println(ClipCapacity(nil))
	fmt.Println("56")
	sl := make([]int, 3, 15)
	fmt.Println(TakeWithLimitedCap(sl, 5))
	fmt.Println(TakeWithLimitedCap([]int{1, 2, 3, 4, 5, 6}, 9))
	fmt.Println(TakeWithLimitedCap([]int{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println("57")
	fmt.Println(Resize(sl, 2))
	fmt.Println(Resize([]int{}, 3))
	fmt.Println(Resize(sl, 5))
	fmt.Println(Resize([]int{1, 2, 3, 4, 5, 6}, 9))
	fmt.Println(Resize([]int{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println("58")
	a, err := EnsureCapacity([]int{1, 2, 3, 4, 5, 6}, 3)
	if err == nil {
	}
	fmt.Printf("a== %v, cap==%d\n", a, cap(a))
	a, err = EnsureCapacity([]int{1, 2, 3, 4, 5, 6}, 9)
	if err != nil {
	}
	fmt.Printf("a== %v, cap==%d\n", a, cap(a))
	var s []int
	a, err = EnsureCapacity(s, -2)
	if err != nil {
		fmt.Printf("alara")
	}
	fmt.Printf("a== %v, cap==%d\n", a, cap(a))
	slice1 = []int{}
	slice1 = append(make([]int, 0, 18), 1, 2, 3, 4, 5, 6)

	// slice2 = []int{}
	slice7, _ := EnsureCapacity(slice1, 3)
	// slice2 = EnsureCapacity(slice1, 3)
	slice2[3] = 19
	fmt.Println(slice7)
	fmt.Println(slice7)
	fmt.Println("len slice1--", len(slice1))
	fmt.Println("cap slice1--", cap(slice1))
	fmt.Println("len slice2--", len(slice7))
	fmt.Println("cap slice2--", cap(slice7))
	fmt.Println(ClipCapacity(nil))
	fmt.Println("------------59")
	fmt.Println("------------59")
	fmt.Println(AppendCopy([]int{1, 2}, 1, 2, 3))
	fmt.Println("------------61")
	a1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(CopyWithin(a1, 2, 0, 4))
	// fmt.Println(a1)
	// fmt.Println(CopyWithin(a1, 1, 3, 6))
	// fmt.Println(a1)
	// a2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // len = 10
	// CopyWithin(a2, 8, 0, 5)                   // dst=8, копируем 5 элементов в позиции 8..12
	// fmt.Println("------------62")
	// a1 = make([]int, 15, 15)
	// // a1 = []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(InsertWithCopy(a1, 2, []int{0, 1, 2}))
	// fmt.Println("------------63")
	// fmt.Println(DeleteAndClear([]string{"1", "2", "3", "4",
	// 	"5"}, 1, 2))
	// fmt.Println(DeleteAndClear([]string{"1", "2", "3", "4",
	// 	"5"}, 1, 4))
}
