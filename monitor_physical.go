//go:build windows

package ddcci

import (
	"errors"
	"fmt"
	"slices"
	"syscall"
	"unsafe"

	"github.com/niluan304/ddcci/vcp"
)

// PhysicalMonitor
// Contains a handle and text description corresponding to a physical monitor.
//
// See [microsoft-windows doc]
//
// [microsoft-windows doc]: https://learn.microsoft.com/en-us/windows/win32/api/physicalmonitorenumerationapi/ns-physicalmonitorenumerationapi-physical_monitor
type PhysicalMonitor struct {
	dxva2 *syscall.DLL

	// Handle to the physical monitor.
	handle syscall.Handle

	// Text description of the physical monitor.
	// A physical monitor description is always an array of 128 characters.
	description string
}

// Description Text description of the physical monitor.
// A physical monitor description is always an array of 128 characters.
func (m *PhysicalMonitor) Description() string {
	return m.description
}

// NewPhysicalMonitor
// Retrieves the physical monitors associated with an HMONITOR monitor handle.
//
// [in] hMonitor
// A monitor handle. PhysicalMonitor handles are returned by
// several Multiple Display PhysicalMonitor functions,
// including EnumDisplayMonitors and MonitorFromWindow,
// which are part of the graphics device interface (GDI).
//
// [out] pPhysicalMonitorArray
// Pointer to an array of PHYSICAL_MONITOR structures. The caller must allocate the array.
func NewPhysicalMonitor(m *SystemMonitor) (*PhysicalMonitor, error) {
	num, err := m.GetNumberOfPhysicalMonitorsFromHMONITOR()
	if err != nil {
		return nil, err
	}

	data := make([]byte, 128)
	err = m.dxva2Call(GetPhysicalMonitorsFromHMONITOR,
		uintptr(num),
		uintptr(unsafe.Pointer(&data[0])),
	)
	if err != nil {
		return nil, err
	}

	description := slices.DeleteFunc(data[8:], func(b byte) bool {
		return b == 0
	})

	return &PhysicalMonitor{
		dxva2:       m.dxva2,
		handle:      syscall.Handle(data[0]),
		description: string(description),
	}, nil
}

// PhysicalMonitor0
// return the first PhysicalMonitor in your Windows computer, so your computer must be connected to a monitor.
//
// it is short name: NewPhysicalMonitor(NewSystemMonitors()[0]),
// if err != nil during in the function, the result will be nil
func PhysicalMonitor0() *PhysicalMonitor {
	monitors, err := NewSystemMonitors()
	if err != nil {
		return nil
	}

	m, err := NewPhysicalMonitor(&monitors[0])
	if err != nil {
		return nil
	}
	return m
}

// short name: m.dxva2.FindProc(name).Call(m.handle, args...)
func (m *PhysicalMonitor) call(name string, args ...uintptr) error {
	proc, err := m.dxva2.FindProc(name)
	if err != nil {
		return errors.Join(err, fmt.Errorf("dxva2 fail to load proc(%s)", name))
	}

	args = slices.Insert(args, 0, uintptr(m.handle))
	r1, r2, errno := proc.Call(args...)
	if !errors.Is(errno, errno0) {
		return errors.Join(errno, fmt.Errorf("proc(%s) fail to call(%v)", name, args))
	}

	_ = r1
	_ = r2

	return nil
}

// GetVCPFeatureAndVCPFeatureReply
// Retrieves the current value, maximum value,
// and code type of Virtual Control Panel (VCP) code for a monitor.
//
// [in] vcpCode
// VCP code to query.
// The VCP codes are Include the VESA PhysicalMonitor Control Command Set (MCCS) standard,
// versions 1.0 and 2.0. This parameter must specify a continuous or non-continuous VCP,
// or a vendor-specific code. It should not be a table control code.
//
// [out] pdwCurrentValue
// Receives the current value of the VCP code. This parameter can be NULL.
//
// [out] pdwMaximumValue
// If vcpCode specifies a continuous VCP code, this parameter receives the maximum value of the VCP code.
// If vcpCode specifies a non-continuous VCP code, the value received in this parameter is undefined.
// This parameter can be NULL.
func (m *PhysicalMonitor) GetVCPFeatureAndVCPFeatureReply(coder vcp.Coder) (current, maxValue int, err error) {
	// type vcpCodeType int
	// const (
	//	// Momentary VCP code. Sending a command of this type causes the monitor to initiate a self-timed operation
	//	// and then revert to its original state. Examples include display tests and degaussing.
	//	Momentary vcpCodeType = iota
	//	// SetParameter Set Parameter VCP code.
	//	// Sending a command of this type changes some aspect of the monitor's operation.
	//	SetParameter
	// )

	// [out] pvct
	//
	// Receives the VCP code type, as a member of the MC_VCP_CODE_TYPE enumeration.
	// This parameter can be NULL.
	codeType := 0

	// // Confuse:
	// // unable to get minValue from GetVCPFeatureAndVCPFeatureReply
	// minValue := 99

	err = m.call(GetVCPFeatureAndVCPFeatureReply,
		uintptr(coder.Code()),
		uintptr(unsafe.Pointer(&codeType)),
		uintptr(unsafe.Pointer(&current)),
		uintptr(unsafe.Pointer(&maxValue)),
		// uintptr(unsafe.Pointer(&minValue)),
	)
	if err != nil {
		return 0, 0, err
	}

	return current, maxValue, nil
}

// SetVCPFeature
// Handle to a physical monitor. To get the monitor handle,
// call GetPhysicalMonitorsFromHMONITOR or GetPhysicalMonitorsFromIDirect3DDevice9.
//
// [in] vcpCode
// VCP code to set.
// The VCP codes are defined in the VESA PhysicalMonitor Control Command Set (MCCS) standard, version 1.0 and 2.0.
// This parameter must specify a continuous or non-continuous VCP,
// or a vendor-specific code. It should not be a table control code.
//
// [in] value
// Value of the VCP code.
//
// [out]
// If the function succeeds, the return value is TRUE.
// If the function fails, the return value is FALSE.
// To get extended error information, call GetLastError.
func (m *PhysicalMonitor) SetVCPFeature(code vcp.VCP) error {
	err := m.call(SetVCPFeature,
		uintptr(code.Code()),
		uintptr(code.Value()),
	)
	if err != nil {
		return err
	}

	return nil
}

// GetBrightness
// Retrieves a monitor's minimum, maximum, and current brightness settings.
//
// [out] pdwMinimumBrightness
// Receives the monitor's minimum brightness.
//
// [out] pdwCurrentBrightness
// Receives the monitor's current brightness.
//
// [out] pdwMaximumBrightness
// Receives the monitor's maximum brightness.
func (m *PhysicalMonitor) GetBrightness() (minValue, current, maxValue int, err error) {
	err = m.call(GetMonitorBrightness,
		uintptr(unsafe.Pointer(&minValue)),
		uintptr(unsafe.Pointer(&current)),
		uintptr(unsafe.Pointer(&maxValue)),
	)

	return
}

// GetCapabilities
// Retrieves the configuration capabilities of a monitor.
// Call this function to find out which high-level monitor configuration functions
// are supported by the monitor.
//
// [out] pdwMonitorCapabilities
// Receives a bitwise OR of capabilities flags. See Remarks.
//
// [out] pdwSupportedColorTemperatures
// Receives a bitwise OR of color temperature flags. See Remarks.
func (m *PhysicalMonitor) GetCapabilities() (capabilities, supportedColorTemperatures int, err error) {
	err = m.call(GetMonitorCapabilities,
		uintptr(unsafe.Pointer(&capabilities)),
		uintptr(unsafe.Pointer(&supportedColorTemperatures)),
	)

	return
}

// GetColorTemperature
// Retrieves a monitor's current color temperature.
//
// [out] pctCurrentColorTemperature
// Receives the monitor's current color temperature,
// specified as a member of the MC_COLOR_TEMPERATURE enumeration.
func (m *PhysicalMonitor) GetColorTemperature() (current int, err error) {
	err = m.call(GetMonitorColorTemperature,
		uintptr(unsafe.Pointer(&current)),
	)

	return
}

// GetContrast
// Retrieves a monitor's minimum, maximum, and current contrast settings.
//
// [out] pdwMinimumContrast
// Receives the monitor's minimum contrast.
//
// [out] pdwCurrentContrast
// Receives the monitor's current contrast.
//
// [out] pdwMaximumContrast
// Receives the monitor's maximum contrast.
func (m *PhysicalMonitor) GetContrast() (minValue, current, maxValue int, err error) {
	err = m.call(GetMonitorContrast,
		uintptr(unsafe.Pointer(&minValue)),
		uintptr(unsafe.Pointer(&current)),
		uintptr(unsafe.Pointer(&maxValue)),
	)

	return
}

type PositionType int

const (
	HorizontalPosition PositionType = iota
	VerticalPosition
)

// GetDisplayAreaPosition
// Retrieves a monitor's minimum, maximum, and current horizontal or vertical position.
//
// [in] ptPositionType
//
// A member of the MC_POSITION_TYPE enumeration,
// specifying whether to retrieve the horizontal position or the vertical position.
//
// [out] pdwMinimumPosition
// Receives the minimum horizontal or vertical position.
//
// [out] pdwCurrentPosition
// Receives the current horizontal or vertical position.
//
// [out] pdwMaximumPosition
// Receives the maximum horizontal or vertical position.
func (m *PhysicalMonitor) GetDisplayAreaPosition(positionType PositionType) (minValue, current, maxValue int, err error) {
	err = m.call(GetMonitorDisplayAreaPosition,
		uintptr(positionType),
		uintptr(unsafe.Pointer(&minValue)),
		uintptr(unsafe.Pointer(&current)),
		uintptr(unsafe.Pointer(&maxValue)),
	)

	return
}

type SizeType int

const (
	Width SizeType = iota
	Height
)

// GetDisplayAreaSize
//
// [out] pdwMinimumWidthOrHeight
// Receives the minimum width or height.
//
// [out] pdwCurrentWidthOrHeight
// Receives the current width or height.
//
// [out] pdwMaximumWidthOrHeight
// Receives the maximum width or height.
func (m *PhysicalMonitor) GetDisplayAreaSize(sizeType SizeType) (minValue, current, maxValue int, err error) {
	err = m.call(GetMonitorDisplayAreaSize,
		uintptr(sizeType),
		uintptr(unsafe.Pointer(&minValue)),
		uintptr(unsafe.Pointer(&current)),
		uintptr(unsafe.Pointer(&maxValue)),
	)

	return
}

type DriveType int

const (
	RedDrive DriveType = iota
	GreenDrive
	BlueDrive
)

// GetRedGreenOrBlueDrive
// Retrieves a monitor's red, green, or blue drive value.
//
// [in] dtDriveType
// A member of the McDriveType enumeration, specifying whether to retrieve the red, green, or blue drive value.
//
// [out] pdwMinimumDrive
// Receives the minimum red, green, or blue drive value.
//
// [out] pdwCurrentDrive
// Receives the current red, green, or blue drive value.
//
// [out] pdwMaximumDrive
// Receives the maximum red, green, or blue drive value.
func (m *PhysicalMonitor) GetRedGreenOrBlueDrive(driveType DriveType) (minValue, current, maxValue int, err error) {
	err = m.call(GetMonitorRedGreenOrBlueDrive,
		uintptr(driveType),
		uintptr(unsafe.Pointer(&minValue)),
		uintptr(unsafe.Pointer(&current)),
		uintptr(unsafe.Pointer(&maxValue)),
	)

	return
}

// GetRedGreenOrBlueGain
// Retrieves a monitor's red, green, or blue gain value.
//
// [in] dtDriveType
// A member of the McDriveType enumeration, specifying whether to retrieve the red, green, or blue drive value.
//
// [out] pdwMinimumGain
// Receives the minimum red, green, or blue gain value.
//
// [out] pdwCurrentGain
// Receives the current red, green, or blue gain value.
//
// [out] pdwMaximumGain
// Receives the maximum red, green, or blue gain value.
func (m *PhysicalMonitor) GetRedGreenOrBlueGain(driveType DriveType) (minValue, current, maxValue int, err error) {
	err = m.call(GetMonitorRedGreenOrBlueGain,
		uintptr(driveType),
		uintptr(unsafe.Pointer(&minValue)),
		uintptr(unsafe.Pointer(&current)),
		uintptr(unsafe.Pointer(&maxValue)),
	)

	return
}

// GetTechnologyType
//
// [out] pdtyDisplayTechnologyType
// Receives the technology type, specified as a member of the MC_DISPLAY_TECHNOLOGY_TYPE enumeration.
func (m *PhysicalMonitor) GetTechnologyType() (value int, err error) {
	err = m.call(GetMonitorTechnologyType,
		uintptr(unsafe.Pointer(&value)),
	)

	return
}

// Degauss
// Degausses a monitor.
// Degaussing improves a monitor's image quality and color fidelity by demagnetizing the monitor.
func (m *PhysicalMonitor) Degauss() error {
	err := m.call(DegaussMonitor)
	if err != nil {
		return err
	}

	return nil
}

// RestoreFactoryColorDefaults
// Restores a monitor's color settings to their factory defaults.
func (m *PhysicalMonitor) RestoreFactoryColorDefaults() error {
	err := m.call(RestoreMonitorFactoryColorDefaults)
	if err != nil {
		return err
	}

	return nil
}

// RestoreFactoryDefaults
// Restores a monitor's settings to their factory defaults.
func (m *PhysicalMonitor) RestoreFactoryDefaults() error {
	err := m.call(RestoreMonitorFactoryDefaults)
	if err != nil {
		return err
	}

	return nil
}

// SaveCurrentMonitorSettings
// Saves the current monitor settings to the display's nonvolatile storage.
func (m *PhysicalMonitor) SaveCurrentMonitorSettings() error {
	err := m.call(SaveCurrentMonitorSettings)
	if err != nil {
		return err
	}

	return nil
}

// SetBrightness
// Sets a monitor's brightness value.
// Increasing the brightness value makes the display on the monitor brighter,
// and decreasing it makes the display dimmer.
//
// [in] dwNewBrightness
// Brightness value. To get the monitor's minimum and maximum brightness values, call GetMonitorBrightness.
//
// It is same as PhysicalMonitor.SetVCPFeature(vcp.Brightness. Value(value))
func (m *PhysicalMonitor) SetBrightness(value int) error {
	err := m.call(SetMonitorBrightness,
		uintptr(value),
	)
	if err != nil {
		return err
	}

	return nil
}

// SetContrast
// Sets a monitor's contrast value.
//
// [in] value
// Contrast value. To get the monitor's minimum and maximum contrast values, call GetMonitorContrast.
//
// It is same as PhysicalMonitor.SetVCPFeature(vcp.Contrast. Value(value))
func (m *PhysicalMonitor) SetContrast(value int) error {
	err := m.call(SetMonitorContrast,
		uintptr(value),
	)
	if err != nil {
		return err
	}

	return nil
}

// GetCapabilitiesStringLength
//
// Retrieves the length of a monitor's capabilities string.
//
// [out] pdwCapabilitiesStringLengthInCharacters
// Receives the length of the capabilities string, in characters, including the terminating null character.
func (m *PhysicalMonitor) GetCapabilitiesStringLength() (size int, err error) {
	err = m.call(GetCapabilitiesStringLength,
		uintptr(unsafe.Pointer(&size)),
	)

	return
}

// CapabilitiesRequestAndCapabilitiesReply
// Retrieves a string describing a monitor's capabilities.
//
// [out] pszASCIICapabilitiesString
// Pointer to a buffer that receives the monitor's capabilities string.
// The caller must allocate this buffer.
func (m *PhysicalMonitor) CapabilitiesRequestAndCapabilitiesReply() (info string, err error) {
	size, err := m.GetCapabilitiesStringLength()
	if err != nil {
		return
	}

	data := make([]byte, size)
	err = m.call(CapabilitiesRequestAndCapabilitiesReply,
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(size),
	)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// TimingReport
// Contains information from a monitor's timing report.
type TimingReport struct {
	// The monitor's horizontal synchronization frequency in Hz.
	HorizontalFrequency uint32

	// The monitor's vertical synchronization frequency in Hz.
	VerticalFrequency uint32

	// Timing status byte.
	// For more information about this value,
	// see the Display Data Channel Command Interface (DDC/CI) standard.
	//
	// 4.7 Get Timing Report & Timing Message
	//  Bit 7 = 1 Sync.Freq. out of range
	//  Bit 6 = 1 Unstable count
	//  Bit 5-2 Reserved, shall be set to 0
	//  Bit 1 = 1 Positive Horizontal sync polarity
	//  Bit 1 = 0 Negative Horizontal sync polarity
	//  Bit 0 = 1 Positive Vertical sync polarity
	//  Bit 0 = 0 Negative Vertical sync polarity
	TimingStatus byte
}

// GetTimingReport
// Retrieves a monitor's horizontal and vertical synchronization frequencies.
//
// [out] pmtrMonitorTimingReport
// Pointer to an TimingReport structure that receives the timing information.
func (m *PhysicalMonitor) GetTimingReport() (*TimingReport, error) {
	report := new(TimingReport)

	err := m.call(GetTimingReport,
		uintptr(unsafe.Pointer(report)),
	)
	if err != nil {
		return nil, err
	}

	return report, nil
}

// SaveCurrentSettings
// Saves the current monitor settings to the display's nonvolatile storage.
//
// note:
// This low-level function is identical to the high-level function SaveCurrentMonitorSettings.
func (m *PhysicalMonitor) SaveCurrentSettings() error {
	err := m.call(SaveCurrentSettings)
	if err != nil {
		return err
	}

	return nil
}
