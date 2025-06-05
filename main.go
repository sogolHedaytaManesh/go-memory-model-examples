package main

import (
	"fmt"
	"time"

	"github.com/sogolHedaytaManesh/go-memory-model-examples/atomic_operations"
	"github.com/sogolHedaytaManesh/go-memory-model-examples/happens_before"
	"github.com/sogolHedaytaManesh/go-memory-model-examples/race_conditions"
	"github.com/sogolHedaytaManesh/go-memory-model-examples/visibility"
)

func main() {
	fmt.Println("== Race Condition Example ==")
	race_conditions.UnsynchronizedIncrement()
	race_conditions.IncrementWithMutex()
	fmt.Println()

	fmt.Println("== Happens-Before Example ==")
	happens_before.HappensBeforeViaChannel()
	fmt.Println()

	fmt.Println("== Atomic Operations Example ==")
	atomic_operations.AtomicIncrement()
	fmt.Println()

	fmt.Println("== Visibility Example ==")
	visibility.VisibilityIssue()
	// Wait a moment so goroutine can run
	time.Sleep(2 * time.Second)

	visibility.FixedVisibilityIssue()
	time.Sleep(2 * time.Second)
}
