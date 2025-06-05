package race_conditions

import (
	"fmt"
	"sync"
)

func UnsynchronizedIncrement() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++ // Data race occurs here
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value (unsynchronized):", counter)
}
