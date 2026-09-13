package main

import ()

// NOTE:1
// first massage
func firstMassage() chan string {
	var ch = make(chan string)
	go func() {
		ch <- "hell"
		// close(ch)
	}()
	return ch
}
