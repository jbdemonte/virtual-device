##  **Example of use of NewJoyConIMU**

```go
package main

import (
	"fmt"
	"github.com/jbdemonte/virtual-device/imu"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

func main() {
	jc := imu.NewJoyConIMU(false)

	err := jc.Register()

	if err != nil {
		fmt.Printf("Failed to register the device: %s\n", err)
		return
	}

	// not required, just to get some time to open `evtest`
	time.Sleep(10_000 * time.Millisecond)

	jc.SendMiscEvent(linux.MSC_TIMESTAMP, -979000346)
	jc.SendAbsoluteEvent(linux.ABS_RX, 4914602)
	jc.SendAbsoluteEvent(linux.ABS_RY, 278018)
	jc.SendAbsoluteEvent(linux.ABS_RZ, -5204642)
	jc.SendAbsoluteEvent(linux.ABS_X, -4308)
	jc.SendAbsoluteEvent(linux.ABS_Y, -2008)
	jc.SendAbsoluteEvent(linux.ABS_Z, -10665)
	jc.SyncReport()

	jc.SendMiscEvent(linux.MSC_TIMESTAMP, -978995680)
	jc.SendAbsoluteEvent(linux.ABS_RX, 6176013)
	jc.SendAbsoluteEvent(linux.ABS_RY, -973063)
	jc.SendAbsoluteEvent(linux.ABS_RZ, -5839476)
	jc.SendAbsoluteEvent(linux.ABS_X, -4536)
	jc.SendAbsoluteEvent(linux.ABS_Y, -2311)
	jc.SendAbsoluteEvent(linux.ABS_Z, -11167)
	jc.SyncReport()
}
```