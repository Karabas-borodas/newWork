package main

import "fmt"

func hasSum(items []int, sum int) bool {
	ma := make(map[int]struct{})
	for _, v := range items {
		// variab := sum - v
		//valiab=11-2=9
		_, ok := ma[sum-v]
		if ok == true {
			return true
		} else {
			ma[v] = struct{}{}
		}
	}
	return false
}

func main() {
	fmt.Println(hasSum([]int{3, 2, 8, 5}, 11)) // -> true
	fmt.Println(hasSum([]int{3, 2, 8, 5}, 12)) // -> false
	fmt.Println(hasSum([]int{3}, 3))           // -> false
	fmt.Println(hasSum([]int{3, 4}, 6))        // -> false
	fmt.Println(hasSum([]int{3, 3}, 6))        // -> true
	fmt.Println(hasSum([]int{}, 6))            // -> false
}

//что такое мапа швейцарские таблицы
// что tакео как устроены
//раздел по api
//по задачам конкурентность
