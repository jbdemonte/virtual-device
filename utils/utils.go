package utils

import (
	"fmt"
	"os"
	"time"
)

// WaitForEventFile polls until the given event file exists or the timeout elapses.
func WaitForEventFile(eventPath string, timeout time.Duration) error {
	start := time.Now()
	for {
		file, err := os.Open(eventPath)
		if err == nil {
			file.Close()
			return nil
		}

		if time.Since(start) > timeout {
			return fmt.Errorf("event file %s is not ready within the timeout", eventPath)
		}

		time.Sleep(10 * time.Millisecond)
	}
}
