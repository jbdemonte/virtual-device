package sdl

// based on https://github.com/libsdl-org/SDL/blob/release-2.30.x/src/joystick/usb_ids.h

type Vendor = uint16
type Product = uint16
type Usage = uint16

const (
	USB_VENDOR_8BITDO       Vendor = 0x2dc8
	USB_VENDOR_AMAZON       Vendor = 0x1949
	USB_VENDOR_APPLE        Vendor = 0x05ac
	USB_VENDOR_ASTRO        Vendor = 0x9886
	USB_VENDOR_ASUS         Vendor = 0x0b05
	USB_VENDOR_BACKBONE     Vendor = 0x358a
	USB_VENDOR_GAMESIR      Vendor = 0x3537
	USB_VENDOR_DRAGONRISE   Vendor = 0x0079
	USB_VENDOR_GOOGLE       Vendor = 0x18d1
	USB_VENDOR_HORI         Vendor = 0x0f0d
	USB_VENDOR_HYPERKIN     Vendor = 0x2e24
	USB_VENDOR_LOGITECH     Vendor = 0x046d
	USB_VENDOR_MADCATZ      Vendor = 0x0738
	USB_VENDOR_MAYFLASH     Vendor = 0x33df
	USB_VENDOR_MICROSOFT    Vendor = 0x045e
	USB_VENDOR_NACON        Vendor = 0x146b
	USB_VENDOR_NACON_ALT    Vendor = 0x3285
	USB_VENDOR_NINTENDO     Vendor = 0x057e
	USB_VENDOR_NVIDIA       Vendor = 0x0955
	USB_VENDOR_PDP          Vendor = 0x0e6f
	USB_VENDOR_POWERA       Vendor = 0x24c6
	USB_VENDOR_POWERA_ALT   Vendor = 0x20d6
	USB_VENDOR_QANBA        Vendor = 0x2c22
	USB_VENDOR_RAZER        Vendor = 0x1532
	USB_VENDOR_SAITEK       Vendor = 0x06a3
	USB_VENDOR_SHANWAN      Vendor = 0x2563
	USB_VENDOR_SHANWAN_ALT  Vendor = 0x20bc
	USB_VENDOR_SONY         Vendor = 0x054c
	USB_VENDOR_THRUSTMASTER Vendor = 0x044f
	USB_VENDOR_TURTLE_BEACH Vendor = 0x10f5
	USB_VENDOR_SWITCH       Vendor = 0x2563
	USB_VENDOR_VALVE        Vendor = 0x28de
	USB_VENDOR_ZEROPLUS     Vendor = 0x0c12
)

const (
	USB_PRODUCT_8BITDO_XBOX_CONTROLLER1                Product = 0x2002 /* Ultimate Wired Controller for Xbox */
	USB_PRODUCT_8BITDO_XBOX_CONTROLLER2                Product = 0x3106 /* Ultimate Wireless / Pro 2 Wired Controller */
	USB_PRODUCT_AMAZON_LUNA_CONTROLLER                 Product = 0x0419
	USB_PRODUCT_ASTRO_C40_XBOX360                      Product = 0x0024
	USB_PRODUCT_BACKBONE_ONE_IOS                       Product = 0x0103
	USB_PRODUCT_BACKBONE_ONE_IOS_PS5                   Product = 0x0104
	USB_PRODUCT_GAMESIR_G7                             Product = 0x1001
	USB_PRODUCT_GOOGLE_STADIA_CONTROLLER               Product = 0x9400
	USB_PRODUCT_EVORETRO_GAMECUBE_ADAPTER1             Product = 0x1843
	USB_PRODUCT_EVORETRO_GAMECUBE_ADAPTER2             Product = 0x1846
	USB_PRODUCT_HORI_FIGHTING_COMMANDER_OCTA_SERIES_X  Product = 0x0150
	USB_PRODUCT_HORI_HORIPAD_PRO_SERIES_X              Product = 0x014f
	USB_PRODUCT_HORI_FIGHTING_STICK_ALPHA_PS4          Product = 0x011c
	USB_PRODUCT_HORI_FIGHTING_STICK_ALPHA_PS5          Product = 0x0184
	USB_PRODUCT_LOGITECH_F310                          Product = 0xc216
	USB_PRODUCT_LOGITECH_CHILLSTREAM                   Product = 0xcad1
	USB_PRODUCT_MADCATZ_SAITEK_SIDE_PANEL_CONTROL_DECK Product = 0x2218
	USB_PRODUCT_NACON_REVOLUTION_5_PRO_PS4_WIRELESS    Product = 0x0d16
	USB_PRODUCT_NACON_REVOLUTION_5_PRO_PS4_WIRED       Product = 0x0d17
	USB_PRODUCT_NACON_REVOLUTION_5_PRO_PS5_WIRELESS    Product = 0x0d18
	USB_PRODUCT_NACON_REVOLUTION_5_PRO_PS5_WIRED       Product = 0x0d19
	USB_PRODUCT_NINTENDO_GAMECUBE_ADAPTER              Product = 0x0337
	USB_PRODUCT_NINTENDO_N64_CONTROLLER                Product = 0x2019
	USB_PRODUCT_NINTENDO_SEGA_GENESIS_CONTROLLER       Product = 0x201e
	USB_PRODUCT_NINTENDO_SNES_CONTROLLER               Product = 0x2017
	USB_PRODUCT_NINTENDO_SWITCH_JOYCON_GRIP            Product = 0x200e
	USB_PRODUCT_NINTENDO_SWITCH_JOYCON_LEFT            Product = 0x2006
	USB_PRODUCT_NINTENDO_SWITCH_JOYCON_PAIR            Product = 0x2008 /* Used by joycond */
	USB_PRODUCT_NINTENDO_SWITCH_JOYCON_RIGHT           Product = 0x2007
	USB_PRODUCT_NINTENDO_SWITCH_PRO                    Product = 0x2009
	USB_PRODUCT_NINTENDO_WII_REMOTE                    Product = 0x0306
	USB_PRODUCT_NINTENDO_WII_REMOTE2                   Product = 0x0330
	USB_PRODUCT_NVIDIA_SHIELD_CONTROLLER_V103          Product = 0x7210
	USB_PRODUCT_NVIDIA_SHIELD_CONTROLLER_V104          Product = 0x7214
	USB_PRODUCT_RAZER_ATROX                            Product = 0x0a00
	USB_PRODUCT_RAZER_KITSUNE                          Product = 0x1012
	USB_PRODUCT_RAZER_PANTHERA                         Product = 0x0401
	USB_PRODUCT_RAZER_PANTHERA_EVO                     Product = 0x1008
	USB_PRODUCT_RAZER_RAIJU                            Product = 0x1000
	USB_PRODUCT_RAZER_TOURNAMENT_EDITION_USB           Product = 0x1007
	USB_PRODUCT_RAZER_TOURNAMENT_EDITION_BLUETOOTH     Product = 0x100a
	USB_PRODUCT_RAZER_ULTIMATE_EDITION_USB             Product = 0x1004
	USB_PRODUCT_RAZER_ULTIMATE_EDITION_BLUETOOTH       Product = 0x1009
	USB_PRODUCT_RAZER_WOLVERINE_V2                     Product = 0x0a29
	USB_PRODUCT_RAZER_WOLVERINE_V2_CHROMA              Product = 0x0a2e
	USB_PRODUCT_RAZER_WOLVERINE_V2_PRO_PS5_WIRED       Product = 0x100b
	USB_PRODUCT_RAZER_WOLVERINE_V2_PRO_PS5_WIRELESS    Product = 0x100c
	USB_PRODUCT_RAZER_WOLVERINE_V2_PRO_XBOX_WIRED      Product = 0x1010
	USB_PRODUCT_RAZER_WOLVERINE_V2_PRO_XBOX_WIRELESS   Product = 0x1011
	USB_PRODUCT_ROG_RAIKIRI                            Product = 0x1a38
	USB_PRODUCT_SAITEK_CYBORG_V3                       Product = 0xf622
	USB_PRODUCT_SHANWAN_DS3                            Product = 0x0523
	USB_PRODUCT_SONY_DS3                               Product = 0x0268
	USB_PRODUCT_SONY_DS4                               Product = 0x05c4
	USB_PRODUCT_SONY_DS4_DONGLE                        Product = 0x0ba0
	USB_PRODUCT_SONY_DS4_SLIM                          Product = 0x09cc
	USB_PRODUCT_SONY_DS4_STRIKEPAD                     Product = 0x05c5
	USB_PRODUCT_SONY_DS5                               Product = 0x0ce6
	USB_PRODUCT_SONY_DS5_EDGE                          Product = 0x0df2
	USB_PRODUCT_SWITCH_RETROBIT_CONTROLLER             Product = 0x0575
	USB_PRODUCT_THRUSTMASTER_ESWAPX_PRO                Product = 0xd012
	USB_PRODUCT_TURTLE_BEACH_SERIES_X_REACT_R          Product = 0x7013
	USB_PRODUCT_TURTLE_BEACH_SERIES_X_RECON            Product = 0x7009
	USB_PRODUCT_VICTRIX_FS_PRO                         Product = 0x0203
	USB_PRODUCT_VICTRIX_FS_PRO_V2                      Product = 0x0207
	USB_PRODUCT_XBOX360_XUSB_CONTROLLER                Product = 0x02a1 /* XUSB driver software PID */
	USB_PRODUCT_XBOX360_WIRED_CONTROLLER               Product = 0x028e
	USB_PRODUCT_XBOX360_WIRELESS_RECEIVER              Product = 0x0719
	USB_PRODUCT_XBOX_ONE_ADAPTIVE                      Product = 0x0b0a
	USB_PRODUCT_XBOX_ONE_ADAPTIVE_BLUETOOTH            Product = 0x0b0c
	USB_PRODUCT_XBOX_ONE_ADAPTIVE_BLE                  Product = 0x0b21
	USB_PRODUCT_XBOX_ONE_ELITE_SERIES_1                Product = 0x02e3
	USB_PRODUCT_XBOX_ONE_ELITE_SERIES_2                Product = 0x0b00
	USB_PRODUCT_XBOX_ONE_ELITE_SERIES_2_BLUETOOTH      Product = 0x0b05
	USB_PRODUCT_XBOX_ONE_ELITE_SERIES_2_BLE            Product = 0x0b22
	USB_PRODUCT_XBOX_ONE_S                             Product = 0x02ea
	USB_PRODUCT_XBOX_ONE_S_REV1_BLUETOOTH              Product = 0x02e0
	USB_PRODUCT_XBOX_ONE_S_REV2_BLUETOOTH              Product = 0x02fd
	USB_PRODUCT_XBOX_ONE_S_REV2_BLE                    Product = 0x0b20
	USB_PRODUCT_XBOX_SERIES_X                          Product = 0x0b12
	USB_PRODUCT_XBOX_SERIES_X_BLE                      Product = 0x0b13
	USB_PRODUCT_XBOX_SERIES_X_VICTRIX_GAMBIT           Product = 0x02d6
	USB_PRODUCT_XBOX_SERIES_X_PDP_BLUE                 Product = 0x02d9
	USB_PRODUCT_XBOX_SERIES_X_PDP_AFTERGLOW            Product = 0x02da
	USB_PRODUCT_XBOX_SERIES_X_POWERA_FUSION_PRO2       Product = 0x4001
	USB_PRODUCT_XBOX_SERIES_X_POWERA_SPECTRA           Product = 0x4002
	USB_PRODUCT_XBOX_ONE_XBOXGIP_CONTROLLER            Product = 0x02ff /* XBOXGIP driver software PID */
	USB_PRODUCT_STEAM_VIRTUAL_GAMEPAD                  Product = 0x11ff
)
const (
	/* USB usage pages */
	USB_USAGEPAGE_GENERIC_DESKTOP Usage = 0x0001
	USB_USAGEPAGE_BUTTON          Usage = 0x0009

	/* USB usages for USAGE_PAGE_GENERIC_DESKTOP */
	USB_USAGE_GENERIC_POINTER             Usage = 0x0001
	USB_USAGE_GENERIC_MOUSE               Usage = 0x0002
	USB_USAGE_GENERIC_JOYSTICK            Usage = 0x0004
	USB_USAGE_GENERIC_GAMEPAD             Usage = 0x0005
	USB_USAGE_GENERIC_KEYBOARD            Usage = 0x0006
	USB_USAGE_GENERIC_KEYPAD              Usage = 0x0007
	USB_USAGE_GENERIC_MULTIAXISCONTROLLER Usage = 0x0008
	USB_USAGE_GENERIC_X                   Usage = 0x0030
	USB_USAGE_GENERIC_Y                   Usage = 0x0031
	USB_USAGE_GENERIC_Z                   Usage = 0x0032
	USB_USAGE_GENERIC_RX                  Usage = 0x0033
	USB_USAGE_GENERIC_RY                  Usage = 0x0034
	USB_USAGE_GENERIC_RZ                  Usage = 0x0035
	USB_USAGE_GENERIC_SLIDER              Usage = 0x0036
	USB_USAGE_GENERIC_DIAL                Usage = 0x0037
	USB_USAGE_GENERIC_WHEEL               Usage = 0x0038
	USB_USAGE_GENERIC_HAT                 Usage = 0x0039

	/* Bluetooth SIG assigned Company Identifiers
	   https://www.bluetooth.com/specifications/assigned-numbers/company-identifiers/ */
	BLUETOOTH_VENDOR_AMAZON = 0x0171

	BLUETOOTH_PRODUCT_LUNA_CONTROLLER = 0x0419
)
