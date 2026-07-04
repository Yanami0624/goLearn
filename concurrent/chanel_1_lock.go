package concurrent

import "sync"

func CountSum1(N int) int {
	var lock sync.Mutex
	cnt := 0
	var wg sync.WaitGroup
	wg.Add(N)

	for range N {
		go func() {
			defer wg.Done()
			lock.Lock()
			cnt++
			lock.Unlock()
		}()
	}

	wg.Wait()
	return cnt
}
