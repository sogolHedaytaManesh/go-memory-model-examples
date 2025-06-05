package race_conditions

import (
	"fmt"
	"sync"
)

func IncrementWithMutex() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++ // ✅ Safe increment
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value (with mutex):", counter)
}
