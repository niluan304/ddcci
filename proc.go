//go:build windows

package ddcci

// https://docs.microsoft.com/en-us/windows/win32/api/_monitor/
const (
	EnumDisplayMonitors = "EnumDisplayMonitors"
	GetMonitorInfoA     = "GetMonitorInfoA"
	GetMonitorInfoW     = "GetMonitorInfoW"
)

// https://learn.microsoft.com/en-us/windows/win32/api/highlevelmonitorconfigurationapi/
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

// https://learn.microsoft.com/en-us/windows/win32/api/highlevelmonitorconfigurationapi/
// Set Monitor With HighLevel
const (
	SetMonitorBrightness = "SetMonitorBrightness"
	SetMonitorContrast   = "SetMonitorContrast"
)

// https://learn.microsoft.com/en-us/windows/win32/api/highlevelmonitorconfigurationapi/
// Set Monitor With HighLevel
//
// Deprecated:
// cannot use preset values because Golang does not support the 'enum' type.
//
// eg:
// compared to *PhysicalMonitor .call(SetMonitorColorTemperature, 4000K),
// it is better: *PhysicalMonitor.SetVCPFeature(vcp.SelectColorPreset.Set4000K)
const (
	SetMonitorColorTemperature    = "SetMonitorColorTemperature"
	SetMonitorDisplayAreaPosition = "SetMonitorDisplayAreaPosition"
	SetMonitorDisplayAreaSize     = "SetMonitorDisplayAreaSize"
	SetMonitorRedGreenOrBlueDrive = "SetMonitorRedGreenOrBlueDrive"
	SetMonitorRedGreenOrBlueGain  = "SetMonitorRedGreenOrBlueGain"
)

// https://learn.microsoft.com/en-us/windows/win32/api/lowlevelmonitorconfigurationapi/
const (
	GetCapabilitiesStringLength             = "GetCapabilitiesStringLength"
	CapabilitiesRequestAndCapabilitiesReply = "CapabilitiesRequestAndCapabilitiesReply"
	GetTimingReport                         = "GetTimingReport"
	SaveCurrentSettings                     = "SaveCurrentSettings"

	GetVCPFeatureAndVCPFeatureReply = "GetVCPFeatureAndVCPFeatureReply"
	SetVCPFeature                   = "SetVCPFeature"
)

// https://learn.microsoft.com/en-us/windows/win32/api/physicalmonitorenumerationapi/
const (
	DestroyPhysicalMonitor                          = "DestroyPhysicalMonitor"
	DestroyPhysicalMonitors                         = "DestroyPhysicalMonitors"
	GetNumberOfPhysicalMonitorsFromHMONITOR         = "GetNumberOfPhysicalMonitorsFromHMONITOR"
	GetNumberOfPhysicalMonitorsFromIDirect3DDevice9 = "GetNumberOfPhysicalMonitorsFromIDirect3DDevice9"
	GetPhysicalMonitorsFromHMONITOR                 = "GetPhysicalMonitorsFromHMONITOR"
	GetPhysicalMonitorsFromIDirect3DDevice9         = "GetPhysicalMonitorsFromIDirect3DDevice9"
)
