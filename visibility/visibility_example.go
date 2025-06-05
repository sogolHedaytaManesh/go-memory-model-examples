package visibility

import (
	"fmt"
	"sync/atomic"
	"time"
)

var running = true

func VisibilityIssue() {
	go func() {
		fmt.Println("Worker started")
		for running {
			// busy wait — may never see updated value!
		}
		fmt.Println("Worker stopped")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Setting running = false")
	running = false

	time.Sleep(1 * time.Second)
}

var runningAtomic int32 = 1

func FixedVisibilityIssue() {
	go func() {
		fmt.Println("Worker (atomic) started")
		for atomic.LoadInt32(&runningAtomic) == 1 {
			// busy wait with atomic read
		}
		fmt.Println("Worker (atomic) stopped")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Setting runningAtomic = 0")
	atomic.StoreInt32(&runningAtomic, 0)

	time.Sleep(1 * time.Second)
}
