package gamepad

import (
	"testing"

	"github.com/jbdemonte/virtual-device/linux"
)

func TestMinMax_SingleValue(t *testing.T) {
	min, max, err := minMax([]int32{5})
	if err != nil {
		t.Fatal(err)
	}
	if min != 5 || max != 5 {
		t.Errorf("got min=%d max=%d, want 5, 5", min, max)
	}
}

func TestMinMax_MultipleValues(t *testing.T) {
	min, max, err := minMax([]int32{3, -1, 7, 0})
	if err != nil {
		t.Fatal(err)
	}
	if min != -1 || max != 7 {
		t.Errorf("got min=%d max=%d, want -1, 7", min, max)
	}
}

func TestMinMax_Empty(t *testing.T) {
	_, _, err := minMax([]int32{})
	if err == nil {
		t.Error("expected error for empty slice")
	}
}

func TestConvertHatToAbsAxis_Basic(t *testing.T) {
	hats := []HatEvent{
		{Axis: linux.ABS_HAT0Y, Value: -1},
		{Axis: linux.ABS_HAT0Y, Value: 1},
	}
	result := convertHatToAbsAxis(hats)
	if len(result) != 1 {
		t.Fatalf("expected 1 axis, got %d", len(result))
	}
	a := result[0]
	if a.Axis != linux.ABS_HAT0Y {
		t.Errorf("axis = %d, want ABS_HAT0Y", a.Axis)
	}
	if a.Min != -1 || a.Max != 1 {
		t.Errorf("min=%d max=%d, want -1, 1", a.Min, a.Max)
	}
	if a.Value != 0 {
		t.Errorf("value=%d, want 0 (midpoint)", a.Value)
	}
}

func TestConvertHatToAbsAxis_MultipleAxes(t *testing.T) {
	hats := []HatEvent{
		{Axis: linux.ABS_HAT0X, Value: -1},
		{Axis: linux.ABS_HAT0X, Value: 1},
		{Axis: linux.ABS_HAT0Y, Value: -1},
		{Axis: linux.ABS_HAT0Y, Value: 1},
	}
	result := convertHatToAbsAxis(hats)
	if len(result) != 2 {
		t.Fatalf("expected 2 axes, got %d", len(result))
	}
}
