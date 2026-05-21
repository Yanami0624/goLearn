package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Node struct {
	name string
}
var count atomic.Int64
var group sync.WaitGroup

func main() {
	var pool sync.Pool
	pool.New = func() any {
		count.Add(1)
		return Node {"Big"}
	}

	const N int = 1000
	group.Add(N)
	for range N {
		go func() {
			defer group.Done()
			target := pool.Get()
			_ = target.(Node)
			pool.Put(target)
		}()
	}
	group.Wait()
	fmt.Println(count)
}