## VirtualDevice Documentation

The `VirtualDevice` interface provides a flexible way to define and control virtual input devices on Linux. Below is a detailed guide on how to use it for various operations.

---

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

#### **Scan Codes**
Send raw scan codes to simulate hardware-level key events:
```go
device.SendScanCode(0x1E)  // Send the scan code for the "A" key
```

### **4. Synchronize Events**
After sending input events, it’s important to synchronize them to ensure the input subsystem processes them correctly. Synchronization informs the system that a complete input report has been sent.

---

#### **Sync a Specific Event**
Synchronize a specific event type (e.g., `SYN_REPORT`) to signal that the current event sequence is complete:
```go
device.Sync(linux.SYN_REPORT)
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
```