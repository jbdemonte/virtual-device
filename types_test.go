package virtual_device

import (
	"math"
	"testing"

	"github.com/jbdemonte/virtual-device/linux"
)

func TestDenormalize_BiDirectional(t *testing.T) {
	axis := AbsAxis{Axis: linux.ABS_X, Min: -32768, Max: 32767}

	tests := []struct {
		name  string
		input float32
		want  int32
	}{
		{"center", 0, 0},          // int32(-32768 + 1*65535/2) = int32(-0.5) = 0
		{"full left", -1, -32768}, // int32(-32768 + 0*65535/2) = -32768
		{"full right", 1, 32767},  // int32(-32768 + 2*65535/2) = 32767
		{"half right", 0.5, 16383},
		{"half left", -0.5, -16384},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := axis.Denormalize(tt.input)
			if got != tt.want {
				t.Errorf("Denormalize(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestDenormalize_UniDirectional(t *testing.T) {
	axis := AbsAxis{Axis: linux.ABS_Z, Min: 0, Max: 255, IsUnidirectional: true}

	tests := []struct {
		name  string
		input float32
		want  int32
	}{
		{"zero", 0, 0},
		{"full", 1, 255},
		{"half", 0.5, 127},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := axis.Denormalize(tt.input)
			if got != tt.want {
				t.Errorf("Denormalize(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestDenormalize_Clamping(t *testing.T) {
	biAxis := AbsAxis{Axis: linux.ABS_X, Min: -32768, Max: 32767}

	// Bi-directional clamps to [-1, 1]
	if got := biAxis.Denormalize(-2); got != -32768 {
		t.Errorf("bi-dir clamp low: got %d, want -32768", got)
	}
	if got := biAxis.Denormalize(2); got != 32767 {
		t.Errorf("bi-dir clamp high: got %d, want 32767", got)
	}

	uniAxis := AbsAxis{Axis: linux.ABS_Z, Min: 0, Max: 255, IsUnidirectional: true}

	// Uni-directional clamps to [0, 1]
	if got := uniAxis.Denormalize(-1); got != 0 {
		t.Errorf("uni-dir clamp low: got %d, want 0", got)
	}
	if got := uniAxis.Denormalize(2); got != 255 {
		t.Errorf("uni-dir clamp high: got %d, want 255", got)
	}
}

func TestDenormalize_ZeroRange(t *testing.T) {
	axis := AbsAxis{Axis: linux.ABS_X, Min: 0, Max: 0}
	got := axis.Denormalize(0.5)
	if got != 0 {
		t.Errorf("zero range: got %d, want 0", got)
	}
}

func TestDenormalize_NaN(t *testing.T) {
	axis := AbsAxis{Axis: linux.ABS_X, Min: -100, Max: 100}
	got := axis.Denormalize(float32(math.NaN()))
	// NaN comparisons fail all branches, so value should not be clamped but will
	// produce some int32 — just ensure no panic.
	_ = got
}
