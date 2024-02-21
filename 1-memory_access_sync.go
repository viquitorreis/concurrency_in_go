package main

import (
	"fmt"
	"sync"
)

func memoryAccess1() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("o valor é %v.\n", data)
	}
}

func memoryAccess2() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		fmt.Printf("o valor é %v.\n", value)
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("o valor é %v.\n", value)
	} else {
		fmt.Printf("o valor é %v.\n", value)
	}
	memoryAccess.Unlock()
}
