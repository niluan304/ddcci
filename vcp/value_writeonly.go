package vcp

// vcp code with write only
type writeonly int

func (code writeonly) Value(value int) VCP { return VCP{code: int(code), value: value} }

const (
	// Degauss
	// Causes a CRT to perform a degauss cycle
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets: CRT
	Degauss writeonly = _Degauss

	// RestoreFactoryDefaults
	// Restore all factory presets including brightness/contrast, geometry, color, and TV defaults.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Preset
	// ddcutil feature subsets: COLOR
	RestoreFactoryDefaults writeonly = _RestoreFactoryDefaults

	// RestoreFactoryBrightnessContrastDefaults
	// Restore factory defaults for brightness and contrast
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Preset
	// ddcutil feature subsets: COLOR
	RestoreFactoryBrightnessContrastDefaults writeonly = _RestoreFactoryBrightnessContrastDefaults

	// RestoreFactoryGeometryDefaults
	// Restore factory defaults for geometry adjustments
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Preset
	// ddcutil feature subsets:
	RestoreFactoryGeometryDefaults writeonly = _RestoreFactoryGeometryDefaults

	// RestoreColorDefaults
	// Restore factory defaults for color settings.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Preset
	// ddcutil feature subsets: COLOR
	RestoreColorDefaults writeonly = _RestoreColorDefaults

	// RestoreFactoryTVDefaults
	// Restore factory defaults for TV functions.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Preset
	// ddcutil feature subsets: TV
	RestoreFactoryTVDefaults writeonly = _RestoreFactoryTVDefaults

	// RemoteProcedureCall
	// Initiates a routine resident in the display
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets: LUT
	RemoteProcedureCall writeonly = _RemoteProcedureCall

	// AuxiliaryDisplayData
	// Sets contents of auxiliary display device
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	AuxiliaryDisplayData writeonly = _AuxiliaryDisplayData
)
