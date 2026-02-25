package utils

import (
	"sync"
	"testing"
)

func TestAtomicBool_DefaultFalse(t *testing.T) {
	var ab AtomicBool
	if ab.Get() {
		t.Error("expected default false")
	}
}

func TestAtomicBool_SetTrue(t *testing.T) {
	var ab AtomicBool
	ab.Set(true)
	if !ab.Get() {
		t.Error("expected true after Set(true)")
	}
}

func TestAtomicBool_SetFalse(t *testing.T) {
	var ab AtomicBool
	ab.Set(true)
	ab.Set(false)
	if ab.Get() {
		t.Error("expected false after Set(false)")
	}
}

func TestAtomicBool_Concurrent(t *testing.T) {
	var ab AtomicBool
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			ab.Set(true)
		}()
		go func() {
			defer wg.Done()
			_ = ab.Get()
		}()
	}
	wg.Wait()
}
