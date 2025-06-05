package atomic_operations

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicIncrement() {
	var counter int64 // must be int32 or int64 for atomic operations
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1) // ✅ atomic operation
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value (atomic):", counter)
}
