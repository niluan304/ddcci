//go:build windows

package ddcci

// See [microsoft-windows doc]
//
// [microsoft-windows doc]: https://docs.microsoft.com/en-us/windows/win32/api/_monitor/
const (
	EnumDisplayMonitors = "EnumDisplayMonitors"
	GetMonitorInfoA     = "GetMonitorInfoA"
	GetMonitorInfoW     = "GetMonitorInfoW"
)

// See [microsoft-windows doc]
//
// [microsoft-windows doc]: https://learn.microsoft.com/en-us/windows/win32/api/highlevelmonitorconfigurationapi/
const (
	GetMonitorBrightness          = "GetMonitorBrightness"
	GetMonitorCapabilities        = "GetMonitorCapabilities"
	GetMonitorColorTemperature    = "GetMonitorColorTemperature"
	GetMonitorContrast            = "GetMonitorContrast"
	GetMonitorDisplayAreaPosition = "GetMonitorDisplayAreaPosition"
	GetMonitorDisplayAreaSize     = "GetMonitorDisplayAreaSize"
	GetMonitorRedGreenOrBlueDrive = "GetMonitorRedGreenOrBlueDrive"
	GetMonitorRedGreenOrBlueGain  = "GetMonitorRedGreenOrBlueGain"
	GetMonitorTechnologyType      = "GetMonitorTechnologyType"

	DegaussMonitor                     = "DegaussMonitor"
	RestoreMonitorFactoryColorDefaults = "RestoreMonitorFactoryColorDefaults"
	RestoreMonitorFactoryDefaults      = "RestoreMonitorFactoryDefaults"
	SaveCurrentMonitorSettings         = "SaveCurrentMonitorSettings"
)

// See [microsoft-windows doc]
// Set Monitor With HighLevel
//
// [microsoft-windows doc]: https://learn.microsoft.com/en-us/windows/win32/api/highlevelmonitorconfigurationapi/
const (
	SetMonitorBrightness = "SetMonitorBrightness"
	SetMonitorContrast   = "SetMonitorContrast"
)

// Set Monitor With HighLevel
//
// See [microsoft-windows doc]
//
// Deprecated:
// cannot use preset values because Golang does not support the 'enum' type.
//
// eg:
// compared to *PhysicalMonitor .call(SetMonitorColorTemperature, 4000K),
// it is better: *PhysicalMonitor.SetVCPFeature(vcp.SelectColorPreset.Set4000K)
//
// [microsoft-windows doc]: https://learn.microsoft.com/en-us/windows/win32/api/highlevelmonitorconfigurationapi/
const (
	SetMonitorColorTemperature    = "SetMonitorColorTemperature"
	SetMonitorDisplayAreaPosition = "SetMonitorDisplayAreaPosition"
	SetMonitorDisplayAreaSize     = "SetMonitorDisplayAreaSize"
	SetMonitorRedGreenOrBlueDrive = "SetMonitorRedGreenOrBlueDrive"
	SetMonitorRedGreenOrBlueGain  = "SetMonitorRedGreenOrBlueGain"
)

// See [microsoft-windows doc]
//
// [microsoft-windows doc]: https://learn.microsoft.com/en-us/windows/win32/api/lowlevelmonitorconfigurationapi/
const (
	GetCapabilitiesStringLength             = "GetCapabilitiesStringLength"
	CapabilitiesRequestAndCapabilitiesReply = "CapabilitiesRequestAndCapabilitiesReply"
	GetTimingReport                         = "GetTimingReport"
	SaveCurrentSettings                     = "SaveCurrentSettings"

	GetVCPFeatureAndVCPFeatureReply = "GetVCPFeatureAndVCPFeatureReply"
	SetVCPFeature                   = "SetVCPFeature"
)

// See [microsoft-windows doc]
//
// [microsoft-windows doc]: https://learn.microsoft.com/en-us/windows/win32/api/physicalmonitorenumerationapi/
const (
	DestroyPhysicalMonitor                          = "DestroyPhysicalMonitor"
	DestroyPhysicalMonitors                         = "DestroyPhysicalMonitors"
	GetNumberOfPhysicalMonitorsFromHMONITOR         = "GetNumberOfPhysicalMonitorsFromHMONITOR"
	GetNumberOfPhysicalMonitorsFromIDirect3DDevice9 = "GetNumberOfPhysicalMonitorsFromIDirect3DDevice9"
	GetPhysicalMonitorsFromHMONITOR                 = "GetPhysicalMonitorsFromHMONITOR"
	GetPhysicalMonitorsFromIDirect3DDevice9         = "GetPhysicalMonitorsFromIDirect3DDevice9"
)
