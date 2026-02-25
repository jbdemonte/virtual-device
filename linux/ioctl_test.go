package linux

import (
	"testing"
	"unsafe"
)

func TestIOC_KnownValues(t *testing.T) {
	// UI_DEV_CREATE = 0x5501 = _IO('U', 1)
	got := _IO('U', 1)
	if got != 0x5501 {
		t.Errorf("_IO('U', 1) = 0x%x, want 0x5501", got)
	}

	// UI_DEV_DESTROY = 0x5502 = _IO('U', 2)
	got = _IO('U', 2)
	if got != 0x5502 {
		t.Errorf("_IO('U', 2) = 0x%x, want 0x5502", got)
	}
}

func TestIOR(t *testing.T) {
	// EVIOCGVERSION = _IOR('E', 0x01, int) = 0x80044501
	// int size on linux/amd64 is 8 bytes, but the constant 0x80044501 implies 4 bytes (int32)
	// The actual Go value depends on platform int size.
	// Test with int32 to match the known constant.
	var i int32
	got := _IOR('E', 0x01, i)
	want := uintptr(0x80044501)
	if got != want {
		t.Errorf("_IOR('E', 0x01, int32) = 0x%x, want 0x%x", got, want)
	}
}

func TestIOW(t *testing.T) {
	// EVIOCSREP = _IOW('E', 0x03, unsigned int[2]) = 0x40084503
	var arr [2]uint32
	got := _IOW('E', 0x03, arr)
	want := uintptr(0x40084503)
	if got != want {
		t.Errorf("_IOW('E', 0x03, [2]uint32) = 0x%x, want 0x%x", got, want)
	}
}

func TestIOC_Direction(t *testing.T) {
	// _IOC with _IOC_NONE should have zero in the direction bits
	none := _IOC(_IOC_NONE, 'U', 1, 0)
	dirBits := none >> _IOC_DIRSHIFT
	if dirBits != 0 {
		t.Errorf("expected direction bits 0, got %d", dirBits)
	}

	// _IOC with _IOC_READ should have 2 in direction bits
	var i int32
	read := _IOC(_IOC_READ, 'E', 0x01, uintptr(unsafe.Sizeof(i)))
	dirBits = read >> _IOC_DIRSHIFT
	if dirBits != _IOC_READ {
		t.Errorf("expected direction bits %d, got %d", _IOC_READ, dirBits)
	}
}
