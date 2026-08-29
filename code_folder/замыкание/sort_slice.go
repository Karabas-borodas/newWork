package main

import (
	"fmt"
	// "os"
	// "math/rand"
	"sort"
)

func sortMagic(slices []int) {
	slReturn := []int{}
	sl1 := []int{}
	sl2 := []int{}
	for _, v := range slices {
		if v%2 == 0 {
			sl1 = append(sl1, v)
		} else {
			sl2 = append(sl2, v)
		}

	}
	sort.Sort(sort.Reverse(sort.IntSlice(sl1)))
	sort.Sort(sort.Reverse(sort.IntSlice(sl2)))
	for i := 0; i < len(sl1); i++ {
		slReturn = append(slReturn, sl1[i])
	}
	for i := 0; i < len(sl2); i++ {
		slReturn = append(slReturn, sl2[i])
	}
	fmt.Println(slReturn)
}

func main() {
	// count := 10
	// sli := make([]int, count)
	// for i := 0; i < count; i++ {
	// 	sli[i] = rand.Intn(100)
	// }
	sli := []int{3, 1, 4, 1}
	// sl := []int{1, 4, 3, 5, 3, 2, 2, 3, 4, 1, 6}
	sortMagic(sli)
}
