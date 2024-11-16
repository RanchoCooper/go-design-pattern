package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type Single struct {
}

var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating Single instance now.")
			singleInstance = &Single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {

	for i := 0; i < 30; i++ {
		go func() {
			obj := getInstance()
			fmt.Printf("%p\n", obj)
		}()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
