package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		// go 重要关键字开启协程模式（goroutine）由于多线程，
		// add函数还未执行完毕 mani方法已经结束了。没有输出。
		go Add(i, i)
	}
}
