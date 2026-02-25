package keyboard

import (
	"testing"

	"github.com/jbdemonte/virtual-device/internal/testutil"
	"github.com/jbdemonte/virtual-device/linux"
)

func newTestKeyboard(mock *testutil.MockDevice, keymap KeyMap) VirtualKeyboard {
	return NewVirtualKeyboardFactory().
		WithDevice(mock).
		WithTapDuration(0).
		WithKeyMap(keymap).
		Create()
}

func TestKeyboard_TypePlainChar(t *testing.T) {
	mock := testutil.NewMockDevice()
	kb := newTestKeyboard(mock, qwertyKeyMap)

	kb.Type("a")
	events := mock.Events()

	// press KEY_A, SyncReport, release KEY_A, SyncReport
	if len(events) != 4 {
		t.Fatalf("expected 4 events, got %d: %+v", len(events), events)
	}
	if events[0].EvType != uint16(linux.EV_KEY) || events[0].Code != uint16(linux.KEY_A) || events[0].Value != 1 {
		t.Errorf("event[0] = %+v, want KEY_A press", events[0])
	}
	if events[2].EvType != uint16(linux.EV_KEY) || events[2].Code != uint16(linux.KEY_A) || events[2].Value != 0 {
		t.Errorf("event[2] = %+v, want KEY_A release", events[2])
	}
}

func TestKeyboard_TypeShiftedChar(t *testing.T) {
	mock := testutil.NewMockDevice()
	kb := newTestKeyboard(mock, qwertyKeyMap)

	kb.Type("A")
	events := mock.Events()

	// press LEFTSHIFT, press KEY_A, SyncReport, release KEY_A, release LEFTSHIFT, SyncReport
	if len(events) != 6 {
		t.Fatalf("expected 6 events, got %d: %+v", len(events), events)
	}
	if events[0].Code != uint16(linux.KEY_LEFTSHIFT) || events[0].Value != 1 {
		t.Errorf("event[0] = %+v, want LEFTSHIFT press", events[0])
	}
	if events[1].Code != uint16(linux.KEY_A) || events[1].Value != 1 {
		t.Errorf("event[1] = %+v, want KEY_A press", events[1])
	}
	if events[3].Code != uint16(linux.KEY_A) || events[3].Value != 0 {
		t.Errorf("event[3] = %+v, want KEY_A release", events[3])
	}
	if events[4].Code != uint16(linux.KEY_LEFTSHIFT) || events[4].Value != 0 {
		t.Errorf("event[4] = %+v, want LEFTSHIFT release", events[4])
	}
}

func TestKeyboard_TypeAltGrChar(t *testing.T) {
	mock := testutil.NewMockDevice()
	// Use azerty keymap where '[' requires AltGr
	kb := newTestKeyboard(mock, azertyKeyMap)

	kb.Type("[")
	events := mock.Events()

	// press RIGHTALT, press KEY_5, SyncReport, release KEY_5, release RIGHTALT, SyncReport
	if len(events) != 6 {
		t.Fatalf("expected 6 events, got %d: %+v", len(events), events)
	}
	if events[0].Code != uint16(linux.KEY_RIGHTALT) || events[0].Value != 1 {
		t.Errorf("event[0] = %+v, want RIGHTALT press", events[0])
	}
}

func TestKeyboard_TypeUnmappedChar(t *testing.T) {
	mock := testutil.NewMockDevice()
	// Use a minimal keymap with no mappings
	kb := newTestKeyboard(mock, KeyMap{})

	kb.Type("x")
	events := mock.Events()

	if len(events) != 0 {
		t.Errorf("expected 0 events for unmapped char, got %d", len(events))
	}
}

func TestKeyboard_TypeMultipleChars(t *testing.T) {
	mock := testutil.NewMockDevice()
	kb := newTestKeyboard(mock, qwertyKeyMap)

	kb.Type("ab")
	events := mock.Events()

	// 2 chars * 4 events each = 8
	if len(events) != 8 {
		t.Fatalf("expected 8 events, got %d", len(events))
	}
}

func TestKeyboard_TapKey(t *testing.T) {
	mock := testutil.NewMockDevice()
	kb := newTestKeyboard(mock, qwertyKeyMap)

	kb.TapKey(linux.KEY_ENTER)
	events := mock.Events()

	// press KEY_ENTER, SyncReport, release KEY_ENTER, SyncReport
	if len(events) != 4 {
		t.Fatalf("expected 4 events, got %d: %+v", len(events), events)
	}
	if events[0].Code != uint16(linux.KEY_ENTER) || events[0].Value != 1 {
		t.Errorf("event[0] = %+v, want KEY_ENTER press", events[0])
	}
	if events[2].Code != uint16(linux.KEY_ENTER) || events[2].Value != 0 {
		t.Errorf("event[2] = %+v, want KEY_ENTER release", events[2])
	}
}
