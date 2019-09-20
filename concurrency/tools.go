package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 3)
	go func() {
		for greeting := range c {
			fmt.Println(greeting)
		}
	}()
	c <- "1"
	c <- "2"
	c <- "3"

	//greeting := <-c

	//go heeloWorld()
	time.Sleep(1 * time.Millisecond)
}

func heeloWorld() {
	fmt.Println("hello Worle!")
}
