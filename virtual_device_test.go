package virtual_device

import (
	"fmt"
	"github.com/jbdemonte/virtual-device/linux"
	"testing"
	"time"
)

func TestVirtualKeyboard(t *testing.T) {

	fmt.Println("starting")
	vd := NewVirtualDevice()
	fmt.Println("NewVirtualDevice ok")

	vd.
		SetBusType(linux.BUS_USB).
		SetProductID(0x02a1).
		SetVendorID(0x045e).
		SetVersion(0x107).
		SetName("Xbox 360 Wireless Receiver (XBOX)").
		SetEventButtons([]linux.Button{
			linux.BTN_NORTH,
			linux.BTN_EAST,
			linux.BTN_SOUTH,
			linux.BTN_WEST,
			linux.BTN_DPAD_UP,
			linux.BTN_DPAD_RIGHT,
			linux.BTN_DPAD_DOWN,
			linux.BTN_DPAD_LEFT,
			linux.BTN_SELECT,
			linux.BTN_START,
		})

	fmt.Println("set done")

	err := vd.Register()
	fmt.Println("Register done")

	if err != nil {
		t.Fatalf("Failed to register the device: %s\n", err)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println("KeyDown...")
		err = vd.KeyDown(int(linux.BTN_NORTH))
		if err != nil {
			t.Fatalf("Failed to send button press. Last error was: %s\n", err)
		}
		fmt.Println("KeyDown... ok")

		time.Sleep(1 * time.Second)
		fmt.Println("KeyUp... ")
		err = vd.KeyUp(int(linux.BTN_NORTH))
		if err != nil {
			t.Fatalf("Failed to send button press. Last error was: %s\n", err)
		}
		fmt.Println("KeyUp... ok")
	}

}
