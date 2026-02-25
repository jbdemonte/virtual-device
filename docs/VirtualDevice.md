# VirtualDevice Documentation

The `VirtualDevice` interface provides a flexible way to define and control virtual input devices on Linux. Below is a detailed guide on how to use it for various operations.

---

## **Interface Definition**

### **Configuration Methods**
These methods allow you to configure the virtual device before registration.

| **Action**           | **Description**                                                                                          |
|----------------------|----------------------------------------------------------------------------------------------------------|
| **`WithPath`**       | Sets the file path for the virtual device. Default is `/dev/uinput`.                                     |
| **`WithMode`**       | Sets the file mode for the device file. Default is `0660`.                                               |
| **`WithQueueLen`**   | Configures the event queue length. Default is `1024`.                                                    |
| **`WithBusType`**    | Sets the device bus type (e.g., `linux.BUS_USB`).                                                        |
| **`WithVendor`**     | Sets the vendor ID of the virtual device. (e.g. `sdl.USB_VENDOR_MICROSOFT`, `0x1234`).                   |
| **`WithProduct`**    | Sets the product ID of the virtual device. (e.g. `sdl.USB_PRODUCT_XBOX_ONE_S`, `0x1234`).                |
| **`WithVersion`**    | Sets the version number for the virtual device. (e.g. `0x01`).                                           |
| **`WithName`**       | Sets the name of the virtual device.                                                                     |
| **`WithKeys`**       | Specifies the keys supported by the device. (e.g. `[]linux.Key{linux.KEY_A, linux.KEY_B, linux.KEY_C}`). |
| **`WithButtons`**    | Specifies the buttons supported by the device. (e.g. `[]linux.Button{linux.BTN_LEFT, linux.BTN_RIGHT}`). |
| **`WithAbsAxes`**    | Configures the absolute axes for the device.                                                             |
| **`WithRelAxes`**    | Configures the relative axes for the device.                                                             |
| **`WithRepeat`**     | Configures key repeat delay and period.                                                                  |
| **`WithLEDs`**       | Specifies the LEDs supported by the device. (e.g. `[]linux.Led{linux.LED_NUML, linux.LED_CAPSL`).        |
| **`WithProperties`** | Sets device-specific properties (e.g., `linux.INPUT_PROP_BUTTONPAD`).                                    |
| **`WithMiscEvents`** | Specifies the miscellaneous events (e.g., `linux.MSC_SCAN`).                                             |


---

### **Lifecycle Methods**
These methods manage the virtual device's lifecycle.

| **Action**       | **Description**                                                                     |
|------------------|-------------------------------------------------------------------------------------|
| **`Register`**   | Registers the virtual device with the system. Must be called before sending events. |
| **`Unregister`** | Unregisters the virtual device, cleaning up resources.                              |


---

### **Event Handling Methods**
These methods send or manipulate input events.

| **Action**              | **Description**                                                 |
|-------------------------|-----------------------------------------------------------------|
| **`Send`**              | Sends a raw input event of the specified type, code, and value. |
| **`Sync`**              | Sends a synchronization event (e.g., `linux.SYN_MT_REPORT`).    |
| **`SyncReport`**        | Sends a default synchronization event (`linux.SYN_REPORT`).     |
| **`PressKey`**          | Simulates a key press event.                                    |
| **`ReleaseKey`**        | Simulates a key release event.                                  |
| **`PressButton`**       | Simulates a button press event.                                 |
| **`ReleaseButton`**     | Simulates a button release event.                               |
| **`SendAbsoluteEvent`** | Sends an absolute axis event with the specified axis and value. |
| **`SendRelativeEvent`** | Sends a relative axis event with the specified axis and value.  |
| **`SetLed`**            | Toggles the state of an LED on the virtual device.              |
| **`SendMiscEvent`**     | Sends a miscellaneous event (e.g., `linux.MSC_SCAN`).           |


## **Usage**

### **1. Configure a Virtual Device**
Use the `With...` methods to configure your virtual device. These methods allow chaining for easy and clear initialization.

Example:
```go
device := virtual_device.NewVirtualDevice().
   WithVendor(sdl.USB_VENDOR_LOGITECH).
   WithProduct(0xc22d).
   WithVersion(0x111).
   WithName("Logitech G510 Gaming Keyboard"),
   WithKeys([]linux.Key{linux.KEY_A, linux.KEY_B, linux.KEY_C}).
   WithLEDs([]linux.Led{linux.LED_CAPSL})
```

### **2. Register the Device**
After configuration, call Register() to create the virtual device

```go
err := device.Register()
if err != nil {
    log.Fatalf("Failed to register virtual device: %v", err)
}
```

### **3. Send Input Events**
The `VirtualDevice` interface allows you to send various types of input events, simulating actions like key presses, mouse movements, and more.

---

#### **Key Events**
Simulate pressing and releasing keyboard keys:
```go
device.PressKey(linux.KEY_A)  // Simulate pressing the "A" key
device.ReleaseKey(linux.KEY_A)  // Simulate releasing the "A" key
```

#### **Button Events**
Simulate pressing and releasing mouse or gamepad buttons:
```go
device.PressButton(linux.BTN_LEFT)  // Simulate pressing the left mouse button
device.ReleaseButton(linux.BTN_LEFT)  // Simulate releasing the left mouse button
```

#### **Absolute Axes Events**
Used for devices like joysticks, touchpads, or drawing tablets. Send specific positions on an axis:
```go
device.SendAbsoluteEvent(linux.ABS_X, 512)  // Move to the middle of the X-axis
device.SendAbsoluteEvent(linux.ABS_Y, 256)  // Move to a lower position on the Y-axis
```

#### **Relative Axes Events**
Used for devices like mice. Send relative movements instead of absolute positions:
```go
device.SendRelativeEvent(linux.REL_X, 10)  // Move cursor 10 units to the right
device.SendRelativeEvent(linux.REL_Y, -5)  // Move cursor 5 units up
```

#### **LED Events**
Control the state of LEDs, such as Caps Lock or Num Lock:
```go
device.SetLed(linux.LED_CAPSL, true)  // Turn on the Caps Lock LED
device.SetLed(linux.LED_CAPSL, false)  // Turn off the Caps Lock LED
```

#### **Misc Events**
Send miscellaneous events, including raw scan codes, to simulate hardware-level key events:
```go
device.SendMiscEvent(linux.MSC_SCAN, 0x1E) // Send the scan code for the "A" key
```

### **4. Synchronize Events**
After sending input events, it’s important to synchronize them to ensure the input subsystem processes them correctly. Synchronization informs the system that a complete input report has been sent.

---

#### **Sync a Specific Event**
Synchronize a specific event type (e.g., `SYN_MT_REPORT`) to signal that the current event sequence is complete:
```go
device.Sync(linux.SYN_MT_REPORT)
```

#### **Sync Report**
Synchronize all pending input events. This is often used after sending a batch of input events to ensure they are processed together:
```go
device.SyncReport()
```

**Why Synchronization is Important?**

Synchronization ensures that all input events sent to the system are properly interpreted and applied. Without synchronization, the system may ignore or misinterpret events, especially when sending multiple inputs in rapid succession.

### **5. Unregister the Device**
When the device is no longer needed, unregister it to release system resources:
```go
err := device.Unregister()
if err != nil {
    log.Fatalf("Failed to unregister virtual device: %v", err)
}
```

#### Example: Simulate Typing
Here’s how to simulate typing a string using a virtual keyboard:

```go
package main

import (
	"fmt"
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

func main() {
	device := virtual_device.NewVirtualDevice().
		WithBusType(linux.BUS_USB).
		WithVendor(0xDEAD).
		WithProduct(0xBABE).
		WithVersion(0x01).
		WithName("Virtual Keyboard").
		WithKeys([]linux.Key{linux.KEY_A, linux.KEY_B, linux.KEY_C})

	device.Register()
	defer device.Unregister()

	device.PressKey(linux.KEY_B)
	device.ReleaseKey(linux.KEY_B)
}
```