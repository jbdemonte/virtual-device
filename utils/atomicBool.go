package utils

import "sync/atomic"

type AtomicBool struct {
	value int32 // Utilise int32 car les opérations atomiques ne supportent pas directement les booléens
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
