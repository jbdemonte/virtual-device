# Inertial Measurement Unit (IMU)

## Nintendo Switch Joy-Con

The Nintendo Switch Joy-Con controllers are equipped with an IMU (Inertial Measurement Unit), which includes:  

An accelerometer that measures linear acceleration on three axes: X, Y, and Z.  
A gyroscope that measures rotational velocity on three axes: RX, RY, and RZ.  
These sensors allow the Joy-Con to detect motion, orientation, and rotation, enabling advanced motion-based controls in games.

### Event Codes
 
| **Event Code** | **Description**              | **Range**                  | **Resolution**       |
|----------------|------------------------------|----------------------------|----------------------|
| `ABS_X`        | Acceleration on X-axis       | -32,767 to +32,767         | 4096 steps per g     |
| `ABS_Y`        | Acceleration on Y-axis       | -32,767 to +32,767         | 4096 steps per g     |
| `ABS_Z`        | Acceleration on Z-axis       | -32,767 to +32,767         | 4096 steps per g     |
| `ABS_RX`       | Rotational velocity around X | -32,767,000 to +32,767,000 | 14,247 steps per °/s |
| `ABS_RY`       | Rotational velocity around Y | -32,767,000 to +32,767,000 | 14,247 steps per °/s |
| `ABS_RZ`       | Rotational velocity around Z | -32,767,000 to +32,767,000 | 14,247 steps per °/s |


### Reading Data

#### Accelerometer Data
The accelerometer measures the acceleration along the X, Y, and Z axes. This includes both dynamic forces (e.g., movement, vibration) and static forces (e.g., gravity).  
To convert the raw data into acceleration in units of g:

```
                             Raw Value  
Acceleration (g) =   -------------------------
                      Resolution (steps per g)
```

For example, if the raw value on the X-axis is 8192 and the resolution is 4096 steps per g, the acceleration is:

```
                   8192
Acceleration 𝑋 = -------- = 2𝑔
                   4096
```

#### Gyroscope Data

The gyroscope measures the rotational velocity (angular speed) around the X, Y, and Z axes. To convert the raw data into degrees per second (°/s):

```
                                    Raw Value
Angular Velocity (°/s) = -------------------------------
                           Resolution (steps per °/s)
```

For example, if the raw value on the RX-axis is 14247 and the resolution is 14247 steps per °/s, the angular velocity is:

```
                        14247
Angular Velocity 𝑅𝑋 = --------- = 1 °/𝑠 
                        14247
```

#### Synchronizing Data

The Joy-Con provides a timestamp event (MSC_TIMESTAMP) for each batch of accelerometer and gyroscope data. This timestamp helps synchronize motion data with other input events, such as button presses or haptic feedback, ensuring accurate motion tracking in time-sensitive applications.