package happens_before

import (
	"fmt"
)

func HappensBeforeViaChannel() {
	var data int
	done := make(chan struct{})

	// Writer goroutine
	go func() {
		data = 42          // Write to shared memory
		done <- struct{}{} // Synchronization point (happens-before)
	}()

	<-done // Reader waits for the signal
	fmt.Println("Data read after synchronization:", data)
}
