package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	// 加锁
	lock.Lock()
	counter++
	fmt.Println(counter)
	// 执行完，解锁
	lock.Unlock()
}
/**
 * C语言处理多线程加锁的处理翻版。c的思路go的实现。
 */
func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()

		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
