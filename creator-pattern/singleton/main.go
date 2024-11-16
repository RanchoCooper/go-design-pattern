package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// 上锁后需要再次检查实例是否为 nil
		if singleInstance == nil {
			// 当前实例为 nil, 创建一个新实例并复制给变量
			singleInstance = &single{}
			fmt.Println("Single instance created.")
		} else {
			fmt.Println("Single instance already created. [in lock range]")
		}
	} else {
		fmt.Println("Single instance already created.[out of lock range]")
	}

	return singleInstance
}
func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
