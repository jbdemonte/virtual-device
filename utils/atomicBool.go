package utils

import "sync/atomic"

// AtomicBool is a thread-safe boolean backed by atomic int32 operations.
type AtomicBool struct {
	value int32 // Uses int32 because atomic operations do not directly support booleans
}

func (ab *AtomicBool) Set(value bool) {
	var intValue int32
	if value {
		intValue = 1
	}
	atomic.StoreInt32(&ab.value, intValue)
}

func (ab *AtomicBool) Get() bool {
	return atomic.LoadInt32(&ab.value) == 1
}
