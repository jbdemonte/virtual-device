package gamepad

import (
	"testing"

	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/internal/testutil"
	"github.com/jbdemonte/virtual-device/linux"
)

func newTestGamepad(mock *testutil.MockDevice) VirtualGamepad {
	return NewVirtualGamepadFactory().
		WithDevice(mock).
		WithDigital(MappingDigital{
			ButtonSouth: linux.BTN_SOUTH,
			ButtonNorth: linux.BTN_NORTH,
			ButtonUp: []InputEvent{
				linux.BTN_TRIGGER_HAPPY3,
				HatEvent{Axis: linux.ABS_HAT0Y, Value: -1},
			},
			ButtonL2: virtual_device.AbsAxis{Axis: linux.ABS_Z, Min: 0, Max: 255},
		}).
		WithLeftStick(MappingStick{
			X: virtual_device.AbsAxis{Axis: linux.ABS_X, Min: -32768, Max: 32767},
			Y: virtual_device.AbsAxis{Axis: linux.ABS_Y, Min: -32768, Max: 32767},
		}).
		Create()
}

func TestGamepad_PressSingleButton(t *testing.T) {
	mock := testutil.NewMockDevice()
	gp := newTestGamepad(mock)

	gp.Press(ButtonSouth)
	events := mock.Events()

	// Expect: PressButton(BTN_SOUTH) + SyncReport
	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d: %+v", len(events), events)
	}
	if events[0].EvType != uint16(linux.EV_KEY) || events[0].Code != uint16(linux.BTN_SOUTH) || events[0].Value != 1 {
		t.Errorf("event[0] = %+v, want EV_KEY BTN_SOUTH press", events[0])
	}
	if events[1].EvType != uint16(linux.EV_SYN) {
		t.Errorf("event[1] = %+v, want SYN_REPORT", events[1])
	}
}

func TestGamepad_ReleaseSingleButton(t *testing.T) {
	mock := testutil.NewMockDevice()
	gp := newTestGamepad(mock)

	gp.Release(ButtonSouth)
	events := mock.Events()

	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}
	if events[0].Value != 0 {
		t.Errorf("expected release (value=0), got %d", events[0].Value)
	}
}

func TestGamepad_PressCompositeButton(t *testing.T) {
	mock := testutil.NewMockDevice()
	gp := newTestGamepad(mock)

	gp.Press(ButtonUp)
	events := mock.Events()

	// ButtonUp maps to: [BTN_TRIGGER_HAPPY3, HatEvent{ABS_HAT0Y, -1}] + SyncReport
	if len(events) != 3 {
		t.Fatalf("expected 3 events, got %d: %+v", len(events), events)
	}
	// BTN_TRIGGER_HAPPY3 press
	if events[0].EvType != uint16(linux.EV_KEY) || events[0].Code != uint16(linux.BTN_TRIGGER_HAPPY3) {
		t.Errorf("event[0] = %+v, want BTN_TRIGGER_HAPPY3 press", events[0])
	}
	// ABS_HAT0Y = -1
	if events[1].EvType != uint16(linux.EV_ABS) || events[1].Code != uint16(linux.ABS_HAT0Y) || events[1].Value != -1 {
		t.Errorf("event[1] = %+v, want ABS_HAT0Y=-1", events[1])
	}
}

func TestGamepad_ReleaseCompositeButton(t *testing.T) {
	mock := testutil.NewMockDevice()
	gp := newTestGamepad(mock)

	gp.Release(ButtonUp)
	events := mock.Events()

	if len(events) != 3 {
		t.Fatalf("expected 3 events, got %d: %+v", len(events), events)
	}
	// Hat release sends value 0
	if events[1].EvType != uint16(linux.EV_ABS) || events[1].Value != 0 {
		t.Errorf("event[1] = %+v, want ABS_HAT0Y=0", events[1])
	}
}

func TestGamepad_PressAbsAxisButton(t *testing.T) {
	mock := testutil.NewMockDevice()
	gp := newTestGamepad(mock)

	// L2 mapped to AbsAxis{ABS_Z, min=0, max=255} — press sends Max
	gp.Press(ButtonL2)
	events := mock.Events()

	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d: %+v", len(events), events)
	}
	if events[0].EvType != uint16(linux.EV_ABS) || events[0].Code != uint16(linux.ABS_Z) || events[0].Value != 255 {
		t.Errorf("event[0] = %+v, want ABS_Z=255", events[0])
	}
}

func TestGamepad_ReleaseAbsAxisButton(t *testing.T) {
	mock := testutil.NewMockDevice()
	gp := newTestGamepad(mock)

	// L2 release sends Min (0)
	gp.Release(ButtonL2)
	events := mock.Events()

	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}
	if events[0].Value != 0 {
		t.Errorf("expected release value 0, got %d", events[0].Value)
	}
}

func TestGamepad_UnmappedButton(t *testing.T) {
	mock := testutil.NewMockDevice()
	gp := newTestGamepad(mock)

	// ButtonEast is not mapped
	gp.Press(ButtonEast)
	if len(mock.Events()) != 0 {
		t.Errorf("expected no events for unmapped button, got %d", len(mock.Events()))
	}
}

func TestGamepad_MoveLeftStick(t *testing.T) {
	mock := testutil.NewMockDevice()
	gp := newTestGamepad(mock)

	gp.MoveLeftStick(0, 0)
	events := mock.Events()

	// SendAbsoluteEvent(ABS_X, denorm(0)) + SendAbsoluteEvent(ABS_Y, denorm(0)) + SyncReport
	if len(events) != 3 {
		t.Fatalf("expected 3 events, got %d: %+v", len(events), events)
	}
	if events[0].EvType != uint16(linux.EV_ABS) || events[0].Code != uint16(linux.ABS_X) {
		t.Errorf("event[0] = %+v, want ABS_X", events[0])
	}
	if events[1].EvType != uint16(linux.EV_ABS) || events[1].Code != uint16(linux.ABS_Y) {
		t.Errorf("event[1] = %+v, want ABS_Y", events[1])
	}
}

func TestGamepad_MoveLeftStick_FullRight(t *testing.T) {
	mock := testutil.NewMockDevice()
	gp := newTestGamepad(mock)

	gp.MoveLeftStickX(1)
	events := mock.Events()

	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}
	if events[0].Value != 32767 {
		t.Errorf("expected 32767 for full right, got %d", events[0].Value)
	}
}

func TestGamepad_Init_ConfiguresDevice(t *testing.T) {
	mock := testutil.NewMockDevice()
	_ = newTestGamepad(mock)

	// init() should have called WithButtons, WithKeys, WithAbsAxes
	if len(mock.Buttons) == 0 {
		t.Error("expected buttons to be configured")
	}
	if len(mock.AbsAxes) == 0 {
		t.Error("expected absolute axes to be configured")
	}
}
