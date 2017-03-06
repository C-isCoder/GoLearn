package main

import "fmt"

/**
 * 随机向ch中写入一个0或者1的过程
 */
func main() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}
}
