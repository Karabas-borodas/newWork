package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var mas = map[string]int{
	"d": 0}
var ma = map[string]int{
	"d": 4,
	"b": 2,
	"c": 3,
	"a": 1,
	"k": 5,
}
var empty = map[string]int{}

// NOTE:1
func CountInts(nums []int) map[int]int {
	ma := make(map[int]int)
	for _, v := range nums {
		ma[v] += 1
	}
	return ma
}

// NOTE:5
func SortedKeys(m map[string]int) []string {
	slice := make([]string, 0)
	for i, _ := range m {
		slice = append(slice, i)
	}

	sort.StringSlice(slice).Sort()
	return slice
}

// NOTE:6
func SumValues(m map[string]int) int {
	summ := 0
	for _, v := range m {
		summ += v
	}
	return summ
}

// NOTE:7
func Invert(m map[string]int) (map[int]string, error) {
	ma := map[int]string{}
	for i, v := range m {
		if _, ok := ma[v]; ok {
			return ma, fmt.Errorf("second element ")
		} else {
			ma[v] = i
		}
	}
	return ma, nil
}

// NOTE:9
func DeleteZeroValues(m map[string]int) {
	for i, v := range m {
		if v == 0 {
			delete(m, i)
		}
	}
}

// NOTE:10
func Clone(m map[string][]int) map[string][]int {
	ma := make(map[string][]int)
	for i, v := range m {
		ma[i] = v
	}
	// ma["1"] = []int{5, 6}
	// m["1"] = []int{7, 7}
	return ma
}

// NOTE:12
func Difference(a, b []int) []int {
	ma := map[int]struct{}{}
	for _, v := range b {
		ma[v] = struct{}{}
	}
	slice := []int{}
	for _, v := range a {
		if _, ok := ma[v]; !ok {
			slice = append(slice, v)
			ma[v] = struct{}{}
		}
	}
	return slice
}

// NOTE:15
func UniqueFold(words []string) []string {
	var str []string
	ma := map[string]struct{}{}
	for i := 0; i < len(words); i++ {
		s := strings.ToLower(words[i])
		if _, ok := ma[s]; !ok {
			ma[s] = struct{}{}
			str = append(str, words[i])
		}
	}
	return str
}

// NOTE:16
func MostFrequent(nums []int) (value int, count int, ok bool) {
	if len(nums) == 0 {
		return 0, 0, false
	}
	value = math.MaxInt
	m := map[int]int{}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	for v := range m {
		if m[v] >= count && v < value {
			count = m[v]
			value = v
		}
	}
	return value, count, true
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 2, 3, 1, 2}
	fmt.Println(slice)
	fmt.Println("#1")
	fmt.Println(CountInts(slice))
	fmt.Println("#5")
	fmt.Println(SortedKeys(ma))
	fmt.Println("#6")
	fmt.Println(SumValues(ma))
	fmt.Println(SumValues(empty))
	fmt.Println("#7")
	m, err := Invert(ma)
	fmt.Println(m)
	fmt.Println(err)
	fmt.Println("#9")
	DeleteZeroValues(ma)
	fmt.Println(mas)
	fmt.Println("#10")
	maq := map[string][]int{"1": {1, 2, 3},
		"2": {2, 3, 4}}
	maa := Clone(maq)
	maq["1"] = []int{3}
	fmt.Println((maa))
	fmt.Println(maq)
	fmt.Println("#12")
	b := []int{2, 5, 2}
	c := []int{}
	fmt.Println(Difference(b, c))
	fmt.Println("#15")
	result := UniqueFold([]string{"aABbcCssS", "aaaa", "AAAA", "B"})
	fmt.Println(result)
	fmt.Println("#16")
	fmt.Println(MostFrequent([]int{1, 1, 1, 2, 2}))
	fmt.Println(MostFrequent([]int{2, 2, 1, 1, 3}))
}
