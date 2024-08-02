package vcp

// vcp code with read only
type readonly int

func (code readonly) Code() int { return int(code) }

const (
	// ColorTemperatureIncrement
	// Color temperature increment used by feature 0Ch Color Temperature Request
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: COLOR
	//
	ColorTemperatureIncrement readonly = _ColorTemperatureIncrement

	// ActiveControl
	// Read id of one feature that has changed, 0x00 indicates no more
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	//
	ActiveControl readonly = _ActiveControl

	// LUTSize
	// Provides the size (number of entries and number of bits/entry) for the Red, Green, and Blue LUT in the display.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Miscellaneous
	// ddcutil feature subsets: LUT
	//
	LUTSize readonly = _LUTSize

	// DisplayIdentificationOperation
	//   Causes a selected 128 byte block of Display Identification Data (EDID or Display ID) to be read
	//   MCCS versions: 2.1, 3.0, 2.2
	//   MCCS specification groups: Miscellaneous
	//   ddcutil feature subsets:
	//   Attributes (v2.1): Read Only, Table (normal)
	//   Attributes (v3.0): Read Only, Table (normal)
	//   Attributes (v2.2): Read Only, Table (normal)
	DisplayIdentificationOperation readonly = _DisplayIdentificationOperation

	// HorizontalFrequency
	// Horizontal sync signal frequency as determined by the display
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	//
	HorizontalFrequency readonly = _HorizontalFrequency

	// VerticalFrequency
	// Vertical sync signal frequency as determined by the display, in .01 hz
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	//
	VerticalFrequency readonly = _VerticalFrequency

	// MonitorStatus
	// Video mode and status of a DPVL capable monitor
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: DPVL
	// ddcutil feature subsets: DPVL
	//
	MonitorStatus readonly = _MonitorStatus

	// DisplayUsageTime
	// Active power on time in hours
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	//
	DisplayUsageTime readonly = _DisplayUsageTime

	// DisplayDescriptorLength
	// Length in bytes of non-volatile storage in the display available for writing a display descriptor, max 256
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	//
	DisplayDescriptorLength readonly = _DisplayDescriptorLength

	// ApplicationEnableKey
	// A 2 byte value used to allow an application to only operate with known products.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	//
	ApplicationEnableKey readonly = _ApplicationEnableKey

	// DisplayFirmwareLevel
	// 2 byte firmware level
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Control, Miscellaneous
	// ddcutil feature subsets:
	//
	DisplayFirmwareLevel readonly = _DisplayFirmwareLevel

	// AuxiliaryDisplaySize
	// Rows and characters/row of auxiliary display
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	//
	AuxiliaryDisplaySize readonly = _AuxiliaryDisplaySize

	// Version
	// MCCS version
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	//
	Version readonly = _Version
)

// vcp code with read only, and non-continuous value
const (
	// ScreenOrientation
	//
	// Indicates screen orientation
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Geometry
	// ddcutil feature subsets:
	// Attributes: Read Only, Non-Continuous (simple)
	//
	// Simple NC values:
	// 	0x01: 0 degrees
	// 	0x02: 90 degrees
	// 	0x03: 180 degrees
	// 	0x04: 270 degrees
	// 	0xff: Display cannot supply orientation
	ScreenOrientation readonly = _ScreenOrientation

	// FlatPanelSubPixelLayout
	//
	// LCD sub-pixel structure
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	// Attributes: Read Only, Non-Continuous (simple)
	//
	// Simple NC values:
	// 	0x00: Sub-pixel layout not defined
	// 	0x01: Red/Green/Blue vertical stripe
	// 	0x02: Red/Green/Blue horizontal stripe
	// 	0x03: Blue/Green/Red vertical stripe
	// 	0x04: Blue/Green/Red horizontal stripe
	// 	0x05: Quad pixel, red at top left
	// 	0x06: Quad pixel, red at bottom left
	// 	0x07: Delta (triad)
	// 	0x08: Mosaic
	FlatPanelSubPixelLayout readonly = _FlatPanelSubPixelLayout

	// DisplayTechnologyType
	//
	// Indicates the base technology type
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	// Attributes: Read Only, Non-Continuous (simple)
	//
	// Simple NC values:
	// 	0x01: CRT (shadow mask)
	// 	0x02: CRT (aperture grill)
	// 	0x03: LCD (active matrix)
	// 	0x04: LCos
	// 	0x05: Plasma
	// 	0x06: OLED
	// 	0x07: EL
	// 	0x08: MEM
	DisplayTechnologyType readonly = _DisplayTechnologyType

	// DisplayControllerType
	//
	// Mfg id of controller and 2 byte manufacturer-specific controller type
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Control, Miscellaneous
	// ddcutil feature subsets:
	// Attributes: Read Only, Non-Continuous (complex)
	//
	// Simple NC values:
	// 	0x01: Conexant
	// 	0x02: Genesis
	// 	0x03: Macronix
	// 	0x04: IDT
	// 	0x05: Mstar
	// 	0x06: Myson
	// 	0x07: Phillips
	// 	0x08: PixelWorks
	// 	0x09: RealTek
	// 	0x0a: Sage
	// 	0x0b: Silicon Image
	// 	0x0c: SmartASIC
	// 	0x0d: STMicroelectronics
	// 	0x0e: Topro
	// 	0x0f: Trumpion
	// 	0x10: Welltrend
	// 	0x11: Samsung
	// 	0x12: Novatek
	// 	0x13: STK
	// 	0x14: Silicon Optics
	// 	0x15: Texas Instruments
	// 	0x16: Analogix
	// 	0x17: Quantum Data
	// 	0x18: NXP Semiconductors
	// 	0x19: Chrontel
	// 	0x1a: Parade Technologies
	// 	0x1b: THine Electronics
	// 	0x1c: Trident
	// 	0x1d: Micros
	// 	0xff: Not defined - a manufacturer designed controller
	DisplayControllerType readonly = _DisplayControllerType
)
