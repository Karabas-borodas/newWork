package main

import (
	"fmt"
)

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
	fmt.Println("#18")
	fmt.Println(AreAnagrams("aA", "Aa"))
	fmt.Println(AreAnagrams("ab", "Aa"))
	fmt.Println("#19")
	fmt.Println(WordFrequency("aa aa aa ad ad "))
	fmt.Println(WordFrequency(""))
}
