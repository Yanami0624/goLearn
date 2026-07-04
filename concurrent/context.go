package concurrent

import (
	"context"
	"fmt"
	"time"
)

func Sleepfor(ctx context.Context) {
	ddl := time.Second * 3
	start := time.Now()
	select {
	case <-ctx.Done():
		fmt.Printf("task finished normally, cost %d milliseconds\n", time.Since(start).Microseconds())
		return
	case <-time.After(ddl):
		fmt.Printf("time set in %d milliseconds, task failed\n", ddl.Milliseconds())
		return
	}
}
