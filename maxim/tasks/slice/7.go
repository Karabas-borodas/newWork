package main

import (
	// "error"
	"fmt"
)

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
	var first = 0
	var last = 0
	sl := make([]int, len(s))
	copy(sl, s)
	if right > len(s)-1 || left < 0 || left >= right {
		return fmt.Errorf("out of range")
	}
	for i := left; i < (right-left+1)/2; i++ {
		first = sl[i]
		last = sl[right-i]
		sl[right-i] = first
		sl[i] = last
	}
	fmt.Println(s)
	return nil
}

// NOTE:13
func RemoveAt(s []int, i int) ([]int, error) {
	if i < 0 || i > len(s) || len(s) == 0 {
		return s, fmt.Errorf("out of range")
	}
	s = append(s[:i], s[i+1:]...)
	return s, nil
}

// NOTE:14
// XXX:чем оличается от 13 если у нас поиск по индексу O(n)?

// NOTE:18
func MoveZerosToEnd(s []int) {
	sl := make([]int, len(s))
	copy(sl, s)
	//NOTE:укорачиваю кописю так как ранее удалял 1 элемнт
	sl = sl[:len(sl)-1]
	lenSlice := len(sl)
	fmt.Println(lenSlice)
	// variable := 0
	for i := lenSlice - 1; i >= 0; i-- {
		if sl[i] == 0 {
			for j := i; j < lenSlice-1; j++ {
				sl[j] = sl[j+1]
			}
			sl[lenSlice-1] = 0
		}
	}
	fmt.Println(sl)
}

// NOTE:27
// XXX:backing array WTF?

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
	// if len(s) == 0 {
	// 	return [][]int{}, nil
	// }
	// if size < 0 {
	// 	return nil, fmt.Errorf("size cant be lou %d", size)
	// }
	//
	// slice := make([][]int,  size)
	// count := size
	// lavelSlice := 0
	// for i, v := range s {
	//
	// 	slice[lavelSlice] = append(slice[lavelSlice], v)
	// 	if i == count {
	// 		count += size + 1
	// 		lavelSlice++
	// 	}
	// }
	// return slice, nil

	if len(s) == 0 {
		return [][]int{}, nil
	}
	if size < 0 {
		return nil, fmt.Errorf("size cant be lou %d", size)
	}

	countChank := (len(s) + size - 1) / size

	slice := make([][]int, countChank)
	count := size
	lavelSlice := 0
	var sl []int
	slice[0] = sl
	for i, v := range s {
		slice[lavelSlice] = append(slice[lavelSlice], v)
		if i == count {
			count += size
			lavelSlice++
			var sl2 []int
			slice[lavelSlice] = sl2
			// slice[lavelSlice] = append(slice[lavelSlice], v)
		}
	}
	return slice, nil
}

// NOTE:49 сделать
func SplitIntoKViews(s []int, k int) ([][]int, error) {
	if len(s) < 1 {
		return [][]int{}, nil
	}
	if k < 1 {
		return nil, fmt.Errorf("aaaaa all bad")
	}
	if k > len(s) {
		k = len(s)
	}
	chankLangth := (len(s)) / k
	chankLitlBit := len(s) % k
	slice := make([][]int, k)
	sl := make([]int, chankLangth+chankLitlBit)
	sliceCount := 0
	slice[sliceCount] = sl
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
	return slice, nil
}

//NOTE:65 сделать

func main() {
	// slice := []float32{2.14, 15, 33.2}
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 9, 8}
	// var summ float32
	fmt.Println("7")
	fmt.Println(IndexOf(slice, 4))
	fmt.Println("12")
	fmt.Println(ReverseRange(slice, 0, 3))
	fmt.Println("13")
	fmt.Println(RemoveAt(slice, 3))
	fmt.Println("18")
	(MoveZerosToEnd(slice))
	MoveZerosToEnd([]int{0, 1, 0, 3, 0, 5, 8})
	fmt.Println("29")
	fmt.Println(Unique([]int{1, 1, 3, 4, 3, 3, 4, 4, 5, 5}))
	fmt.Println("47")
	fmt.Println(ChunkViews([]int{1, 1, 3, 4, 3, 3, 4, 4, 5, 5}, 3))
	fmt.Println(ChunkViews([]int{1, 4, 2, 3}, 3))
	fmt.Println("49")
	fmt.Println(SplitIntoKViews([]int{1, 4, 2, 3}, 3))
}
