package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	// "strings"
)

func main() {
	s := "kalabanga--"

	file, err := os.Create("wasya.txt")
	if err != nil {
		fmt.Printf("aaaaa %v", err)
	}
	wr := bufio.NewWriter(file)
	wr.WriteString(s)
	err = wr.Flush()
	if err != nil {
		fmt.Printf("aaaaa %v", err)
	}

	file.Close()
	data, err := os.ReadFile("wasya.txt")
	if err != nil {
		fmt.Printf("aaaaa %v", err)
	}
	fmt.Println(string(data))
}
