package concurrent

import (
	"sync"
)

func CountSum0(N int) int {
	lock := make(chan struct{}, 1)
	cnt := 0

	var wg sync.WaitGroup
	wg.Add(N)
	for range N {
		go func() {
			defer wg.Done()
			lock <- struct{}{}
			cnt++
			<-lock
		}()
	}
	wg.Wait()
	return cnt
}
