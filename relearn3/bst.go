package main

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

type Tree = tree.Tree

func Walk(head *Tree, ch chan int) {
	if head == nil {
		return
	}
	Walk(head.Left, ch)
	ch <- head.Value
	Walk(head.Right, ch)
}

func Same(t1 *Tree, t2 *Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 || val1 != val2 {
			return false
		}
	}

}

func IsBSTSame() {
	fmt.Println("is bst same")
	start := time.Now()
	fmt.Println(Same(tree.New(1000000), tree.New(1000000)))
	fmt.Println(Same(tree.New(1000000), tree.New(15000000)))
	fmt.Println(time.Since(start))
}
