package mouse

import (
	"testing"

	"github.com/jbdemonte/virtual-device/internal/testutil"
	"github.com/jbdemonte/virtual-device/linux"
)

func newTestMouse(mock *testutil.MockDevice) VirtualMouse {
	return NewVirtualMouseFactory().
		WithDevice(mock).
		WithClickDelay(0).
		WithDoubleClickDelay(0).
		Create()
}

func newTestMouseWithHighRes(mock *testutil.MockDevice, vStep int32) VirtualMouse {
	return NewVirtualMouseFactory().
		WithDevice(mock).
		WithClickDelay(0).
		WithDoubleClickDelay(0).
		WithHighResStepVertical(vStep).
		Create()
}

func TestMouse_Move(t *testing.T) {
	mock := testutil.NewMockDevice()
	m := newTestMouse(mock)

	m.Move(10, -5)
	events := mock.Events()

	// REL_X + REL_Y + SyncReport
	if len(events) != 3 {
		t.Fatalf("expected 3 events, got %d: %+v", len(events), events)
	}
	if events[0].EvType != uint16(linux.EV_REL) || events[0].Code != uint16(linux.REL_X) || events[0].Value != 10 {
		t.Errorf("event[0] = %+v, want REL_X=10", events[0])
	}
	if events[1].EvType != uint16(linux.EV_REL) || events[1].Code != uint16(linux.REL_Y) || events[1].Value != -5 {
		t.Errorf("event[1] = %+v, want REL_Y=-5", events[1])
	}
}

func TestMouse_ScrollVertical_NoHighRes(t *testing.T) {
	mock := testutil.NewMockDevice()
	m := newTestMouse(mock)

	m.ScrollVertical(1)
	events := mock.Events()

	// REL_WHEEL + SyncReport (no hi-res)
	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d: %+v", len(events), events)
	}
	if events[0].Code != uint16(linux.REL_WHEEL) || events[0].Value != 1 {
		t.Errorf("event[0] = %+v, want REL_WHEEL=1", events[0])
	}
}

func TestMouse_ScrollVertical_WithHighRes(t *testing.T) {
	mock := testutil.NewMockDevice()
	m := newTestMouseWithHighRes(mock, 120)

	m.ScrollVertical(1)
	events := mock.Events()

	// REL_WHEEL + REL_WHEEL_HI_RES + SyncReport
	if len(events) != 3 {
		t.Fatalf("expected 3 events, got %d: %+v", len(events), events)
	}
	if events[1].Code != uint16(linux.REL_WHEEL_HI_RES) || events[1].Value != 120 {
		t.Errorf("event[1] = %+v, want REL_WHEEL_HI_RES=120", events[1])
	}
}

func TestMouse_ButtonPress(t *testing.T) {
	mock := testutil.NewMockDevice()
	m := newTestMouse(mock)

	m.ButtonPress(linux.BTN_LEFT)
	events := mock.Events()

	// PressButton + SyncReport
	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}
	if events[0].EvType != uint16(linux.EV_KEY) || events[0].Code != uint16(linux.BTN_LEFT) || events[0].Value != 1 {
		t.Errorf("event[0] = %+v, want BTN_LEFT press", events[0])
	}
}

func TestMouse_Click(t *testing.T) {
	mock := testutil.NewMockDevice()
	m := newTestMouse(mock)

	m.Click(linux.BTN_LEFT)
	events := mock.Events()

	// ButtonPress (press+sync) + ButtonRelease (release+sync) = 4
	if len(events) != 4 {
		t.Fatalf("expected 4 events, got %d: %+v", len(events), events)
	}
	if events[0].Value != 1 {
		t.Errorf("expected press first, got value %d", events[0].Value)
	}
	if events[2].Value != 0 {
		t.Errorf("expected release, got value %d", events[2].Value)
	}
}

func TestMouse_DoubleClick(t *testing.T) {
	mock := testutil.NewMockDevice()
	m := newTestMouse(mock)

	m.DoubleClick(linux.BTN_LEFT)
	events := mock.Events()

	// 2 clicks * 4 events = 8
	if len(events) != 8 {
		t.Fatalf("expected 8 events, got %d", len(events))
	}
}

func TestMouse_MoveX(t *testing.T) {
	mock := testutil.NewMockDevice()
	m := newTestMouse(mock)

	m.MoveX(42)
	events := mock.Events()

	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}
	if events[0].Code != uint16(linux.REL_X) || events[0].Value != 42 {
		t.Errorf("event[0] = %+v, want REL_X=42", events[0])
	}
}

func TestMouse_ScrollDown(t *testing.T) {
	mock := testutil.NewMockDevice()
	m := newTestMouse(mock)

	m.ScrollDown()
	events := mock.Events()

	if len(events) < 1 {
		t.Fatal("expected at least 1 event")
	}
	if events[0].Code != uint16(linux.REL_WHEEL) || events[0].Value != -1 {
		t.Errorf("event[0] = %+v, want REL_WHEEL=-1", events[0])
	}
}
