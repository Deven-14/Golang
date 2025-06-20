package main

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func WalkAndClose(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go WalkAndClose(t1, ch1)
	go WalkAndClose(t2, ch2)
	for {
		value1, ok1 := <-ch1
		value2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 || value1 != value2 {
			return false
		}
	}
}

func Same2(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for {
		value1, ok1 := <-ch1
		value2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 || value1 != value2 {
			return false
		}
	}
}

func Same3(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go WalkAndClose(t1, ch1)
	go WalkAndClose(t2, ch2)
	for {
		value1, ok1 := <-ch1
		value2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 || value1 != value2 {
			return false
		}
	}
}

func bst() {
	start := time.Now()
	fmt.Println(Same(tree.New(1000000), tree.New(1000000)))
	fmt.Println(Same(tree.New(1000000), tree.New(15000000)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(Same2(tree.New(1000000), tree.New(1000000)))
	fmt.Println(Same2(tree.New(1000000), tree.New(15000000)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(Same3(tree.New(1000000), tree.New(1000000)))
	fmt.Println(Same3(tree.New(1000000), tree.New(15000000)))
	fmt.Println(time.Since(start))
}
