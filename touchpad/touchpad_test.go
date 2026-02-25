package touchpad

import (
	"testing"

	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/internal/testutil"
	"github.com/jbdemonte/virtual-device/linux"
)

func newTestTouchpad(mock *testutil.MockDevice) VirtualTouchpad {
	return NewVirtualTouchpadFactory().
		WithDevice(mock).
		WithClickDelay(0).
		WithDoubleClickDelay(0).
		WithAxes([]virtual_device.AbsAxis{
			{Axis: linux.ABS_X, Min: 0, Max: 1000, IsUnidirectional: true},
			{Axis: linux.ABS_Y, Min: 0, Max: 1000, IsUnidirectional: true},
			{Axis: linux.ABS_PRESSURE, Min: 0, Max: 255, IsUnidirectional: true},
			{Axis: linux.ABS_MT_POSITION_X, Min: 0, Max: 1000, IsUnidirectional: true},
			{Axis: linux.ABS_MT_POSITION_Y, Min: 0, Max: 1000, IsUnidirectional: true},
		}).
		WithButtons([]linux.Button{
			linux.BTN_LEFT, linux.BTN_RIGHT,
			linux.BTN_TOOL_FINGER, linux.BTN_TOOL_DOUBLETAP,
			linux.BTN_TOOL_TRIPLETAP,
		}).
		Create()
}

func TestTouchpad_AssignSlotIfNeeded_NoDuplicates(t *testing.T) {
	mock := testutil.NewMockDevice()
	tp := newTestTouchpad(mock).(*virtualTouchpad)

	slots := []TouchSlot{
		{Slot: 0, X: 0.1, Y: 0.2, Pressure: 0.5},
		{Slot: 1, X: 0.3, Y: 0.4, Pressure: 0.5},
	}
	result := tp.assignSlotIfNeeded(slots)
	if result[0].Slot != 0 || result[1].Slot != 1 {
		t.Errorf("slots changed unexpectedly: %+v", result)
	}
}

func TestTouchpad_AssignSlotIfNeeded_DuplicateSlots(t *testing.T) {
	mock := testutil.NewMockDevice()
	tp := newTestTouchpad(mock).(*virtualTouchpad)

	slots := []TouchSlot{
		{Slot: 0, X: 0.1, Y: 0.2, Pressure: 0.5},
		{Slot: 0, X: 0.3, Y: 0.4, Pressure: 0.5},
	}
	result := tp.assignSlotIfNeeded(slots)
	if result[0].Slot == result[1].Slot {
		t.Errorf("duplicate slots not resolved: both are %d", result[0].Slot)
	}
}

func TestTouchpad_MultiTouchB_SingleFingerPress(t *testing.T) {
	mock := testutil.NewMockDevice()
	tp := newTestTouchpad(mock)

	slots := tp.MultiTouch([]TouchSlot{
		{Slot: 0, X: 0.5, Y: 0.5, Pressure: 0.5},
	})

	if len(slots) != 1 {
		t.Fatalf("expected 1 slot returned, got %d", len(slots))
	}

	events := mock.Events()
	if len(events) == 0 {
		t.Fatal("expected events")
	}

	// Should contain ABS_MT_SLOT, ABS_MT_TRACKING_ID, ABS_MT_POSITION_X/Y, PRESSURE, finger count, SyncReport
	hasSlot := false
	hasTrackingID := false
	for _, e := range events {
		if e.EvType == uint16(linux.EV_ABS) && e.Code == uint16(linux.ABS_MT_SLOT) {
			hasSlot = true
		}
		if e.EvType == uint16(linux.EV_ABS) && e.Code == uint16(linux.ABS_MT_TRACKING_ID) && e.Value >= 0 {
			hasTrackingID = true
		}
	}
	if !hasSlot {
		t.Error("missing ABS_MT_SLOT event")
	}
	if !hasTrackingID {
		t.Error("missing ABS_MT_TRACKING_ID event")
	}
}

func TestTouchpad_MultiTouchB_FingerRelease(t *testing.T) {
	mock := testutil.NewMockDevice()
	tp := newTestTouchpad(mock)

	// Press finger
	tp.MultiTouch([]TouchSlot{
		{Slot: 0, X: 0.5, Y: 0.5, Pressure: 0.5},
	})
	mock.Reset()

	// Release finger (pressure=0)
	tp.MultiTouch([]TouchSlot{
		{Slot: 0, X: 0.5, Y: 0.5, Pressure: 0},
	})
	events := mock.Events()

	// Should have tracking_id = -1 for release
	hasRelease := false
	for _, e := range events {
		if e.EvType == uint16(linux.EV_ABS) && e.Code == uint16(linux.ABS_MT_TRACKING_ID) && e.Value == -1 {
			hasRelease = true
		}
	}
	if !hasRelease {
		t.Error("missing ABS_MT_TRACKING_ID=-1 for release")
	}
}

func TestTouchpad_MultiTouchB_FingerCountToggle(t *testing.T) {
	mock := testutil.NewMockDevice()
	tp := newTestTouchpad(mock)

	// Press 1 finger
	tp.MultiTouch([]TouchSlot{
		{Slot: 0, X: 0.5, Y: 0.5, Pressure: 0.5},
	})

	// Check BTN_TOOL_FINGER was pressed
	events := mock.Events()
	hasFingerBtn := false
	for _, e := range events {
		if e.EvType == uint16(linux.EV_KEY) && e.Code == uint16(linux.BTN_TOOL_FINGER) && e.Value == 1 {
			hasFingerBtn = true
		}
	}
	if !hasFingerBtn {
		t.Error("missing BTN_TOOL_FINGER press for 1 finger")
	}

	mock.Reset()

	// Press 2nd finger
	tp.MultiTouch([]TouchSlot{
		{Slot: 1, X: 0.3, Y: 0.3, Pressure: 0.5},
	})
	events = mock.Events()

	// Should release BTN_TOOL_FINGER and press BTN_TOOL_DOUBLETAP
	hasFingerRelease := false
	hasDoubleTap := false
	for _, e := range events {
		if e.EvType == uint16(linux.EV_KEY) && e.Code == uint16(linux.BTN_TOOL_FINGER) && e.Value == 0 {
			hasFingerRelease = true
		}
		if e.EvType == uint16(linux.EV_KEY) && e.Code == uint16(linux.BTN_TOOL_DOUBLETAP) && e.Value == 1 {
			hasDoubleTap = true
		}
	}
	if !hasFingerRelease {
		t.Error("missing BTN_TOOL_FINGER release when going to 2 fingers")
	}
	if !hasDoubleTap {
		t.Error("missing BTN_TOOL_DOUBLETAP press for 2 fingers")
	}
}

func TestTouchpad_Touch(t *testing.T) {
	mock := testutil.NewMockDevice()
	tp := newTestTouchpad(mock)

	tp.Touch(0.5, 0.5, 0.5)
	events := mock.Events()

	// ABS_X + ABS_Y + ABS_PRESSURE + SyncReport
	if len(events) != 4 {
		t.Fatalf("expected 4 events, got %d: %+v", len(events), events)
	}
	if events[0].EvType != uint16(linux.EV_ABS) || events[0].Code != uint16(linux.ABS_X) {
		t.Errorf("event[0] = %+v, want ABS_X", events[0])
	}
	if events[1].Code != uint16(linux.ABS_Y) {
		t.Errorf("event[1] = %+v, want ABS_Y", events[1])
	}
	if events[2].Code != uint16(linux.ABS_PRESSURE) {
		t.Errorf("event[2] = %+v, want ABS_PRESSURE", events[2])
	}
}
