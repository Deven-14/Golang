package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	m map[string]int
	sync.RWMutex
}

func (sc *SafeCounter) increment(key string) {
	sc.Lock()
	sc.m[key] += 1
	sc.Unlock()
}

func (sc *SafeCounter) value(key string) int {
	sc.RLock()
	defer sc.RUnlock()
	return sc.m[key] // value would be zero value of int i.e., 0
}

func mutexExample() {
	sc := SafeCounter{m: make(map[string]int)}
	for range 1000 {
		go sc.increment("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(sc.value("somekey"))
}
