//go:build integration

package gamepad

import (
	"os"
	"testing"
)

func TestIntegration_XBox360_Lifecycle(t *testing.T) {
	gp := NewXBox360()

	if err := gp.Register(); err != nil {
		t.Fatalf("Register: %v", err)
	}
	defer gp.Unregister()

	path := gp.EventPath()
	if path == "" {
		t.Fatal("EventPath is empty after Register")
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("event file %s does not exist: %v", path, err)
	}
}
