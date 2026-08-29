package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	// "math/rand"
	// "io"
	"sort"
)

func sortMagic(slices []int) []int {
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
	return slReturn
}

func main() {
	file, err := os.OpenFile("sorts_file.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("error create file %v", err)
		os.Exit(1)
	}
	defer file.Close()
	sli := []int{1, 4, 3, 5, 3, 2, 2, 3, 4, 1, 6}
	dlReturn := sortMagic(sli)
	fmt.Println(dlReturn)
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprint(dlReturn) + "\n")
	if err != nil {
		fmt.Errorf("error erite file %v", err)
	}
	writer.Flush()

	//NOTE: чтение файла чере os
	slic, err := os.ReadFile("sorts_file.txt")
	for _, v := range slic {
		fmt.Println(v)

	}
	for _, v := range slic {
		fmt.Printf("%c", v)

	}
	//NOTE: перевод курсора в начало файла
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Errorf("error chene place cursor%v", err)
	}
	//NOTE:чтение чрезе пакет io
	sl, err := io.ReadAll(file)
	if err != nil {
		fmt.Errorf("error erite file %v", err)
	}

	fmt.Println(string(sl))
	for _, v := range sl {

		fmt.Println(v)

	}
	for _, v := range sl {
		fmt.Printf("---%c", v)

	}
}
