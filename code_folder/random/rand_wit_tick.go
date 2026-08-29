package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {

	var i = rand.IntN(100)
	fmt.Println(i)
	tick := time.NewTicker(time.Microsecond * 100)
	time.Sleep(2 * time.Second)
	t := <-tick.C
	fmt.Println(<-tick.C)
	fmt.Println(t.Second() % 60)
}
