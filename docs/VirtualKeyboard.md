## VirtualKeyboard Documentation

The `VirtualKeyboard` interface provides a simple way to simulate keyboard inputs and control keyboard-related features, such as LEDs, using a virtual device. 
The `VirtualKeyboardFactory` is used to configure and create instances of `VirtualKeyboard`.

---

### **VirtualKeyboard**

The `VirtualKeyboard` interface defines the core functionalities of a virtual keyboard.

#### **Methods**

| **Action**     | **Description**                                                                      |
|----------------|--------------------------------------------------------------------------------------|
| **Register**   | Registers the virtual keyboard device with the system.                               |
| **Unregister** | Unregisters the virtual keyboard device, releasing system resources.                 |
| **PressKey**   | Simulates pressing a specific key.                                                   |
| **ReleaseKey** | Simulates releasing a specific key.                                                  |
| **Type**       | Simulates typing a string of characters using the virtual keyboard.                  |
| **SetLed**     | Controls the state of a keyboard LED (e.g., Caps Lock or Num Lock).                  |
| **Send**       | Sends a raw input event of the specified type, code, and value.                      |


####  **Keyboard Layout Detection in `Type`**

The `Type` function is designed to simulate typing a string on the virtual keyboard. 
It automatically detects the system's keyboard layout (e.g., `fr`, `us`, `de`) and maps it to a predefined keymap (e.g., AZERTY, QWERTY). 
This ensures that the characters typed match the expected output for the system's configured layout.

You can shunt this process by providing your own KeyMap with the factory function `WithKeyMap`.

---

### **How Layout Detection Works**

1. **System Layout Detection**:  
   The package uses system commands (e.g., `localectl` or `setxkbmap`) to query the current keyboard layout. The detected layout code (e.g., `fr` for French, `us` for US English) determines which keymap to use.

2. **Keymap Association**:  
   The detected layout code is associated with a predefined hardcoded keymap in the package. For example:
  - `fr` → **AZERTY** (French)
  - `us` → **QWERTY** (US English)

These keymaps cover many common configurations but can be extended or replaced to support additional layouts or custom configurations. (planned)

3. **Typing Simulation**:  
   When a string is passed to the `Type` function, the package uses the selected keymap to determine the appropriate key codes and modifiers (e.g., Shift, AltGr) for each character.

---

### **VirtualKeyboardFactory**

The `VirtualKeyboardFactory` is used to configure and create instances of `VirtualKeyboard`. 
It supports method chaining for easy setup.

#### **Methods**

| **Action**           | **Description**                                                            |
|----------------------|----------------------------------------------------------------------------|
| **`WithDevice`**     | Attaches an existing `VirtualDevice` to the keyboard.                      |
| **`WithKeys`**       | Configures the list of supported keys for the keyboard.                    |
| **`WithLEDs`**       | Configures the LEDs supported by the keyboard.                             |
| **`WithRepeat`**     | Sets the repeat delay and period for held keys.                            |
| **`WithKeyMap`**     | Specifies a custom keymap to use with the keyboard.                        |
| **`WithMiscEvents`** | Specifies the miscellaneous events (e.g., `linux.MSC_SCAN`).               |
| **`Create`**         | Creates an instance of `VirtualKeyboard` with the specified configuration. |


---

### **Example Usage**

Here’s how to configure and use a `VirtualKeyboard`:

```go
package main

import (
   "fmt"
   virtual_device "github.com/jbdemonte/virtual-device"
   "github.com/jbdemonte/virtual-device/keyboard"
   "github.com/jbdemonte/virtual-device/linux"
   "log"
)

func main() {
   kb := keyboard.NewVirtualKeyboardFactory().
      WithDevice(
         virtual_device.NewVirtualDevice().
            WithBusType(linux.BUS_USB).
            WithVendor(0xDEAD).
            WithProduct(0xBEEF).
            WithVersion(0x01).
            WithName("My Virtual Keyboard"),
      ).
      WithKeys([]linux.Key{linux.KEY_A, linux.KEY_B, linux.KEY_C}).
      WithLEDs([]linux.Led{linux.LED_CAPSL}).
      WithRepeat(250, 33).
      Create()

   err := kb.Register()
   if err != nil {
      log.Fatalf("Failed to register virtual keyboard: %v", err)
   }
   defer kb.Unregister()

   kb.Type("Hello, world!")

   kb.SetLed(linux.LED_CAPSL, true)
}
```

This documentation outlines the essential steps for configuring, registering, and using a `VirtualKeyboard` to simulate keyboard inputs and control related features.