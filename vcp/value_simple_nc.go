package vcp

// vcp code with simple non-continuous value
var (
	// NewControlValue
	//
	// Indicates that a display user control (other than power) has been used to change and save (or autosave) a new value.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	// Attributes: Read Write, Non-Continuous (complex)
	//
	// Simple NC values:
	//    0x01: No new control values
	//    0x02: One or more new control values have been saved
	//    0xff: No user controls are present
	NewControlValue newControlValue

	// SoftControls
	//
	// Allows display controls to be used as soft keys
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: No button active
	//    0x01: Button 1 active
	//    0x02: Button 2 active
	//    0x03: Button 3 active
	//    0x04: Button 4 active
	//    0x05: Button 5 active
	//    0x06: Button 6 active
	//    0x07: Button 7 active
	//    0xff: No user controls are present
	SoftControls softControls

	// SelectColorPreset
	//
	// Select a specified color temperature
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes (v2.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.1): Read Write, Non-Continuous (simple)
	// Attributes (v3.0): Read Write, Non-Continuous (complex)
	// Attributes (v2.2): Read Write, Non-Continuous (complex)
	//
	// Simple NC values:
	//    0x01: sRGB
	//    0x02: Display Native
	//    0x03: 4000 K
	//    0x04: 5000 K
	//    0x05: 6500 K
	//    0x06: 7500 K
	//    0x07: 8200 K
	//    0x08: 9300 K
	//    0x09: 10000 K
	//    0x0a: 11500 K
	//    0x0b: User 1
	//    0x0c: User 2
	//    0x0d: User 3
	SelectColorPreset selectColorPreset

	// AutoSetup
	//
	// Perform auto setup function (H/V position, clock, clock phase, A/D converter, etc.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets:
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Auto setup not active
	//    0x01: Performing auto setup
	//    0x02: Enable continuous/periodic auto setup
	AutoSetup autoSetup

	// AutoColorSetup
	//
	// Perform color auto setup function (R/G/B gain and offset, A/D setup, etc.
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: COLOR
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Auto setup not active
	//    0x01: Performing auto setup
	//    0x02: Enable continuous/periodic auto setup
	AutoColorSetup autoColorSetup

	// InputSource
	//
	// Selects active video source
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	// Attributes (v2.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.1): Read Write, Non-Continuous (simple)
	// Attributes (v3.0): Read Write, Table (normal)
	// Attributes (v2.2): Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x01: VGA-1
	//    0x02: VGA-2
	//    0x03: DVI-1
	//    0x04: DVI-2
	//    0x05: Composite video 1
	//    0x06: Composite video 2
	//    0x07: S-Video-1
	//    0x08: S-Video-2
	//    0x09: Tuner-1
	//    0x0a: Tuner-2
	//    0x0b: Tuner-3
	//    0x0c: Component video (YPrPb/YCrCb) 1
	//    0x0d: Component video (YPrPb/YCrCb) 2
	//    0x0e: Component video (YPrPb/YCrCb) 3
	//    0x0f: DisplayPort-1
	//    0x10: DisplayPort-2
	//    0x11: HDMI-1
	//    0x12: HDMI-2
	InputSource inputSource

	// SpeakerSelect
	//
	// Selects a group of speakers
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Audio
	// ddcutil feature subsets: AUDIO
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Front L/R
	//    0x01: Side L/R
	//    0x02: Rear L/R
	//    0x03: Center/Subwoofer
	SpeakerSelect speakerSelect

	// AmbientLightSensor
	//
	// Enable/Disable ambient light sensor
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x01: Disabled
	//    0x02: Enabled
	AmbientLightSensor ambientLightSensor

	// HorizontalMirrorFlip
	//
	// Flip picture horizontally
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Geometry
	// ddcutil feature subsets:
	// Attributes (v2.0): Write Only, Non-Continuous (write-only)
	// Attributes (v2.1): Read Write, Non-Continuous (simple)
	// Attributes (v3.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.2): Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Normal mode
	//    0x01: Mirrored horizontally mode
	HorizontalMirrorFlip horizontalMirrorFlip

	// VerticalMirrorFlip
	//
	// Flip picture vertically
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Geometry
	// ddcutil feature subsets:
	// Attributes (v2.0): Write Only, Non-Continuous (write-only)
	// Attributes (v2.1): Read Write, Non-Continuous (simple)
	// Attributes (v3.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.2): Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Normal mode
	//    0x01: Mirrored vertically mode
	VerticalMirrorFlip verticalMirrorFlip

	// DisplayScaling
	//
	// Control the scaling (input vs output) of the display
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets:
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x01: No scaling
	//    0x02: Max image, no aspect ration distortion
	//    0x03: Max vertical image, no aspect ratio distortion
	//    0x04: Max horizontal image, no aspect ratio distortion
	//    0x05: Max vertical image with aspect ratio distortion
	//    0x06: Max horizontal image with aspect ratio distortion
	//    0x07: Linear expansion (compression) on horizontal axis
	//    0x08: Linear expansion (compression) on h and v axes
	//    0x09: Squeeze mode
	//    0x0a: Non-linear expansion
	DisplayScaling displayScaling

	// Sharpness
	//
	// Selects one of a range of algorithms. Increasing (decreasing) the value must increase (decrease) the edge sharpness of image features.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets:
	// Attributes (v2.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.1): Read Write, Continuous (normal)
	// Attributes (v3.0): Read Write, Continuous (normal)
	// Attributes (v2.2): Read Write, Continuous (normal)
	//
	// Simple NC values:
	//    0x01: Filter function 1
	//    0x02: Filter function 2
	//    0x03: Filter function 3
	//    0x04: Filter function 4
	Sharpness sharpness

	// TVChannelUpDown
	//
	// Increment (1) or decrement (2) television channel
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets: TV
	// Attributes: Write Only, Non-Continuous (write-only)
	//
	// Simple NC values:
	//    0x01: Increment channel
	//    0x02: Decrement channel
	TVChannelUpDown tvChannelUpDown

	// AudioMuteScreenBlank
	//
	// Mute/unmute audio, and (v2.2) screen blank
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets: TV, AUDIO
	// Attributes (v2.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.1): Read Write, Non-Continuous (simple)
	// Attributes (v3.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.2): Read Write, Non-Continuous (complex)
	//
	// Simple NC values:
	//    0x01: Mute the audio
	//    0x02: Unmute the audio
	AudioMuteScreenBlank audioMuteScreenBlank

	// AudioProcessorMode
	//
	// Select audio mode
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Audio
	// ddcutil feature subsets: TV, AUDIO
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Speaker off/Audio not supported
	//    0x01: Mono
	//    0x02: Stereo
	//    0x03: Stereo expanded
	//    0x11: SRS 2.0
	//    0x12: SRS 2.1
	//    0x13: SRS 3.1
	//    0x14: SRS 4.1
	//    0x15: SRS 5.1
	//    0x16: SRS 6.1
	//    0x17: SRS 7.1
	//    0x21: Dolby 2.0
	//    0x22: Dolby 2.1
	//    0x23: Dolby 3.1
	//    0x24: Dolby 4.1
	//    0x25: Dolby 5.1
	//    0x26: Dolby 6.1
	//    0x27: Dolby 7.1
	//    0x31: THX 2.0
	//    0x32: THX 2.1
	//    0x33: THX 3.1
	//    0x34: THX 4.1
	//    0x35: THX 5.1
	//    0x36: THX 6.1
	//    0x37: THX 7.1
	AudioProcessorMode audioProcessorMode

	// WindowControl
	//
	// Enables the brightness and color within a window to be different from the desktop.
	// MCCS versions: 2.0, 2.1
	// MCCS specification groups: Window
	// ddcutil feature subsets: WINDOW
	// Attributes (v2.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.1): Read Write, Non-Continuous (simple)
	// Attributes (v3.0): Deprecated,
	// Attributes (v2.2): Deprecated,
	//
	// Simple NC values:
	//    0x00: No effect
	//    0x01: Off
	//    0x02: On
	WindowControl windowControl

	// AutoSetupOnOff
	//
	// Turn on/off an auto setup function
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets:
	// Attributes: Write Only, Non-Continuous (write-only)
	//
	// Simple NC values:
	//    0x01: Off
	//    0x02: On
	AutoSetupOnOff autoSetupOnOff

	// ChangeTheSelectedWindow
	//
	// Change selected window (as defined by 95h..98h)
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Window
	// ddcutil feature subsets: WINDOW
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Full display image area selected except active windows
	//    0x01: Window 1 selected
	//    0x02: Window 2 selected
	//    0x03: Window 3 selected
	//    0x04: Window 4 selected
	//    0x05: Window 5 selected
	//    0x06: Window 6 selected
	//    0x07: Window 7 selected
	ChangeTheSelectedWindow changeTheSelectedWindow

	// Settings
	//
	// Store/restore the user saved values for the current mode.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Preset
	// ddcutil feature subsets:
	// Attributes: Write Only, Non-Continuous (write-only)
	//
	// Simple NC values:
	//    0x01: Store current settings in the monitor
	//    0x02: Restore factory defaults for current mode
	Settings settings

	// OSDButtonControl
	// Sets and indicates the current operational state of OSD (and buttons in v2.2)
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Control, Miscellaneous
	// ddcutil feature subsets:
	// Attributes (v2.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.1): Read Write, Non-Continuous (simple)
	// Attributes (v3.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.2): Read Write, Non-Continuous (complex)
	//
	// Simple NC values:
	//    0x01: OSD Disabled
	//    0x02: OSD Enabled
	//    0xff: Display cannot supply this information
	OSDButtonControl osdButtonControl

	// OSDLanguage
	// On Screen Display language
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Control, Miscellaneous
	// ddcutil feature subsets:
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Reserved value, must be ignored
	//    0x01: Chinese (traditional, Hantai)
	//    0x02: English
	//    0x03: French
	//    0x04: German
	//    0x05: Italian
	//    0x06: Japanese
	//    0x07: Korean
	//    0x08: Portuguese (Portugal)
	//    0x09: Russian
	//    0x0a: Spanish
	//    0x0b: Swedish
	//    0x0c: Turkish
	//    0x0d: Chinese (simplified / Kantai)
	//    0x0e: Portuguese (Brazil)
	//    0x0f: Arabic
	//    0x10: Bulgarian
	//    0x11: Croatian
	//    0x12: Czech
	//    0x13: Danish
	//    0x14: Dutch
	//    0x15: Estonian
	//    0x16: Finnish
	//    0x17: Greek
	//    0x18: Hebrew
	//    0x19: Hindi
	//    0x1a: Hungarian
	//    0x1b: Latvian
	//    0x1c: Lithuanian
	//    0x1d: Norwegian
	//    0x1e: Polish
	//    0x1f: Romanian
	//    0x20: Serbian
	//    0x21: Slovak
	//    0x22: Slovenian
	//    0x23: Thai
	//    0x24: Ukranian
	//    0x25: Vietnamese
	OSDLanguage osdLanguage
	// OutputSelect
	//
	// Selects the active output
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	// Attributes (v2.0): Read Write, Non-Continuous (simple)
	// Attributes (v2.1): Read Write, Non-Continuous (simple)
	// Attributes (v3.0): Read Write, Table (normal)
	// Attributes (v2.2): Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x01: Analog video (R/G/B) 1
	//    0x02: Analog video (R/G/B) 2
	//    0x03: Digital video (TDMS) 1
	//    0x04: Digital video (TDMS) 22
	//    0x05: Composite video 1
	//    0x06: Composite video 2
	//    0x07: S-Video-1
	//    0x08: S-Video-2
	//    0x09: Tuner-1
	//    0x0a: Tuner-2
	//    0x0b: Tuner-3
	//    0x0c: Component video (YPrPb/YCrCb) 1
	//    0x0d: Component video (YPrPb/YCrCb) 2
	//    0x0e: Component video (YPrPb/YCrCb) 3
	//    0x0f: DisplayPort-1
	//    0x10: DisplayPort-2
	//    0x11: HDMI-1
	//    0x12: HDMI-2
	OutputSelect outputSelect

	// PowerMode
	//
	// DPM and DPMS status
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Control, Miscellaneous
	// ddcutil feature subsets:
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x01: DPM: On,  DPMS: Off
	//    0x02: DPM: Off, DPMS: Standby
	//    0x03: DPM: Off, DPMS: Suspend
	//    0x04: DPM: Off, DPMS: Off
	//    0x05: Write only value to turn off display
	PowerMode powerMode

	// AuxiliaryPowerOutput
	//
	// Controls an auxiliary power output from a display to a host device
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x01: Disable auxiliary power
	//    0x02: Enable Auxiliary power
	AuxiliaryPowerOutput auxiliaryPowerOutput

	// ScanMode
	//
	// Controls scan characteristics (aka format)
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Geometry
	// ddcutil feature subsets: CRT
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Normal operation
	//    0x01: UnderScan
	//    0x02: OverScan
	//    0x03: Widescreen
	ScanMode scanMode

	// ImageMode
	//
	// Controls aspects of the displayed image (TV applications)
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Control
	// ddcutil feature subsets: TV
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: No effect
	//    0x01: Full mode
	//    0x02: Zoom mode
	//    0x03: Squeeze mode
	//    0x04: Variable
	ImageMode imageMode

	// DisplayMode
	//
	// Type of application used on display
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: COLOR
	// Attributes: Read Write, Non-Continuous (simple)
	//
	// Simple NC values:
	//    0x00: Standard/Default mode
	//    0x01: Productivity
	//    0x02: Mixed
	//    0x03: Movie
	//    0x04: User defined
	//    0x05: Games
	//    0x06: Sports
	//    0x07: Professional (all signal processing disabled)
	//    0x08: Standard/Default mode with intermediate power consumption
	//    0x09: Standard/Default mode with low power consumption
	//    0x0a: Demonstration
	//    0xf0: Dynamic contrast
	DisplayMode displayMode
)

func (newControlValue) Code() int         { return _NewControlValue }         // support Read
func (softControls) Code() int            { return _SoftControls }            // support Read
func (selectColorPreset) Code() int       { return _SelectColorPreset }       // support Read
func (autoSetup) Code() int               { return _AutoSetup }               // support Read
func (autoColorSetup) Code() int          { return _AutoColorSetup }          // support Read
func (inputSource) Code() int             { return _InputSource }             // support Read
func (speakerSelect) Code() int           { return _SpeakerSelect }           // support Read
func (ambientLightSensor) Code() int      { return _AmbientLightSensor }      // support Read
func (horizontalMirrorFlip) Code() int    { return _HorizontalMirrorFlip }    // support Read
func (verticalMirrorFlip) Code() int      { return _VerticalMirrorFlip }      // support Read
func (displayScaling) Code() int          { return _DisplayScaling }          // support Read
func (sharpness) Code() int               { return _Sharpness }               // support Read
func (audioMuteScreenBlank) Code() int    { return _AudioMuteScreenBlank }    // support Read
func (audioProcessorMode) Code() int      { return _AudioProcessorMode }      // support Read
func (windowControl) Code() int           { return _WindowControl }           // support Read
func (changeTheSelectedWindow) Code() int { return _ChangeTheSelectedWindow } // support Read
func (osdButtonControl) Code() int        { return _OSDButtonControl }        // support Read
func (osdLanguage) Code() int             { return _OSDLanguage }             // support Read
func (outputSelect) Code() int            { return _OutputSelect }            // support Read
func (powerMode) Code() int               { return _PowerMode }               // support Read
func (auxiliaryPowerOutput) Code() int    { return _AuxiliaryPowerOutput }    // support Read
func (scanMode) Code() int                { return _ScanMode }                // support Read
func (imageMode) Code() int               { return _ImageMode }               // support Read
func (displayMode) Code() int             { return _DisplayMode }             // support Read

// newControlValue
type newControlValue struct{}

func (newControlValue) NoNewControlValues() VCP { return VCP{code: _NewControlValue, value: 0x01} }
func (newControlValue) OneOrMoreNewControlValuesHaveBeenSaved() VCP {
	return VCP{code: _NewControlValue, value: 0x02}
}

func (newControlValue) NoUserControlsArePresent() VCP {
	return VCP{code: _NewControlValue, value: 0xff}
}

// softControls
type softControls struct{}

func (softControls) NoButtonActive() VCP           { return VCP{code: _SoftControls, value: 0x00} }
func (softControls) Button1Active() VCP            { return VCP{code: _SoftControls, value: 0x01} }
func (softControls) Button2Active() VCP            { return VCP{code: _SoftControls, value: 0x02} }
func (softControls) Button3Active() VCP            { return VCP{code: _SoftControls, value: 0x03} }
func (softControls) Button4Active() VCP            { return VCP{code: _SoftControls, value: 0x04} }
func (softControls) Button5Active() VCP            { return VCP{code: _SoftControls, value: 0x05} }
func (softControls) Button6Active() VCP            { return VCP{code: _SoftControls, value: 0x06} }
func (softControls) Button7Active() VCP            { return VCP{code: _SoftControls, value: 0x07} }
func (softControls) NoUserControlsArePresent() VCP { return VCP{code: _SoftControls, value: 0xff} }

// selectColorPreset
type selectColorPreset struct{}

func (selectColorPreset) SRGB() VCP          { return VCP{code: _SelectColorPreset, value: 0x01} }
func (selectColorPreset) DisplayNative() VCP { return VCP{code: _SelectColorPreset, value: 0x02} }
func (selectColorPreset) Set4000K() VCP      { return VCP{code: _SelectColorPreset, value: 0x03} }
func (selectColorPreset) Set5000K() VCP      { return VCP{code: _SelectColorPreset, value: 0x04} }
func (selectColorPreset) Set6500K() VCP      { return VCP{code: _SelectColorPreset, value: 0x05} }
func (selectColorPreset) Set7500K() VCP      { return VCP{code: _SelectColorPreset, value: 0x06} }
func (selectColorPreset) Set8200K() VCP      { return VCP{code: _SelectColorPreset, value: 0x07} }
func (selectColorPreset) Set9300K() VCP      { return VCP{code: _SelectColorPreset, value: 0x08} }
func (selectColorPreset) Set10000K() VCP     { return VCP{code: _SelectColorPreset, value: 0x09} }
func (selectColorPreset) Set11500K() VCP     { return VCP{code: _SelectColorPreset, value: 0x0a} }
func (selectColorPreset) User1() VCP         { return VCP{code: _SelectColorPreset, value: 0x0b} }
func (selectColorPreset) User2() VCP         { return VCP{code: _SelectColorPreset, value: 0x0c} }
func (selectColorPreset) User3() VCP         { return VCP{code: _SelectColorPreset, value: 0x0d} }

// autoSetup
type autoSetup struct{}

func (autoSetup) AutoSetupNotActive() VCP  { return VCP{code: _AutoSetup, value: 0x00} }
func (autoSetup) PerformingAutoSetup() VCP { return VCP{code: _AutoSetup, value: 0x01} }
func (autoSetup) EnableContinuousOrPeriodicAutoSetup() VCP {
	return VCP{code: _AutoSetup, value: 0x02}
}

// autoColorSetup
type autoColorSetup struct{}

func (autoColorSetup) AutoSetupNotActive() VCP  { return VCP{code: _AutoColorSetup, value: 0x00} }
func (autoColorSetup) PerformingAutoSetup() VCP { return VCP{code: _AutoColorSetup, value: 0x01} }
func (autoColorSetup) EnableContinuousOrPeriodicAutoSetup() VCP {
	return VCP{code: _AutoColorSetup, value: 0x02}
}

// inputSource
type inputSource struct{}

func (inputSource) VGA1() VCP                      { return VCP{code: _InputSource, value: 0x01} }
func (inputSource) VGA2() VCP                      { return VCP{code: _InputSource, value: 0x02} }
func (inputSource) DVI1() VCP                      { return VCP{code: _InputSource, value: 0x03} }
func (inputSource) DVI2() VCP                      { return VCP{code: _InputSource, value: 0x04} }
func (inputSource) CompositeVideo1() VCP           { return VCP{code: _InputSource, value: 0x05} }
func (inputSource) CompositeVideo2() VCP           { return VCP{code: _InputSource, value: 0x06} }
func (inputSource) SVideo1() VCP                   { return VCP{code: _InputSource, value: 0x07} }
func (inputSource) SVideo2() VCP                   { return VCP{code: _InputSource, value: 0x08} }
func (inputSource) Tuner1() VCP                    { return VCP{code: _InputSource, value: 0x09} }
func (inputSource) Tuner2() VCP                    { return VCP{code: _InputSource, value: 0x0a} }
func (inputSource) Tuner3() VCP                    { return VCP{code: _InputSource, value: 0x0b} }
func (inputSource) ComponentVideoYPrPbYCrCb1() VCP { return VCP{code: _InputSource, value: 0x0c} }
func (inputSource) ComponentVideoYPrPbYCrCb2() VCP { return VCP{code: _InputSource, value: 0x0d} }
func (inputSource) ComponentVideoYPrPbYCrCb3() VCP { return VCP{code: _InputSource, value: 0x0e} }
func (inputSource) DisplayPort1() VCP              { return VCP{code: _InputSource, value: 0x0f} }
func (inputSource) DisplayPort2() VCP              { return VCP{code: _InputSource, value: 0x10} }
func (inputSource) HDMI1() VCP                     { return VCP{code: _InputSource, value: 0x11} }
func (inputSource) HDMI2() VCP                     { return VCP{code: _InputSource, value: 0x12} }

// speakerSelect
type speakerSelect struct{}

func (speakerSelect) FrontLR() VCP         { return VCP{code: _SpeakerSelect, value: 0x00} }
func (speakerSelect) SideLR() VCP          { return VCP{code: _SpeakerSelect, value: 0x01} }
func (speakerSelect) RearLR() VCP          { return VCP{code: _SpeakerSelect, value: 0x02} }
func (speakerSelect) CenterSubwoofer() VCP { return VCP{code: _SpeakerSelect, value: 0x03} }

// ambientLightSensor
type ambientLightSensor struct{}

func (ambientLightSensor) Disabled() VCP { return VCP{code: _AmbientLightSensor, value: 0x01} }
func (ambientLightSensor) Enabled() VCP  { return VCP{code: _AmbientLightSensor, value: 0x02} }

// horizontalMirrorFlip
type horizontalMirrorFlip struct{}

func (horizontalMirrorFlip) NormalMode() VCP {
	return VCP{code: _HorizontalMirrorFlip, value: 0x00}
}

func (horizontalMirrorFlip) MirroredHorizontallyMode() VCP {
	return VCP{code: _HorizontalMirrorFlip, value: 0x01}
}

// verticalMirrorFlip
type verticalMirrorFlip struct{}

func (verticalMirrorFlip) NormalMode() VCP { return VCP{code: _VerticalMirrorFlip, value: 0x00} }
func (verticalMirrorFlip) MirroredVerticallyMode() VCP {
	return VCP{code: _VerticalMirrorFlip, value: 0x01}
}

// displayScaling
type displayScaling struct{}

func (displayScaling) NoScaling() VCP { return VCP{code: _DisplayScaling, value: 0x01} }
func (displayScaling) MaxImageNoAspectRationDistortion() VCP {
	return VCP{code: _DisplayScaling, value: 0x02}
}

func (displayScaling) MaxVerticalImageNoAspectRatioDistortion() VCP {
	return VCP{code: _DisplayScaling, value: 0x03}
}

func (displayScaling) MaxHorizontalImageNoAspectRatioDistortion() VCP {
	return VCP{code: _DisplayScaling, value: 0x04}
}

func (displayScaling) MaxVerticalImageWithAspectRatioDistortion() VCP {
	return VCP{code: _DisplayScaling, value: 0x05}
}

func (displayScaling) MaxHorizontalImageWithAspectRatioDistortion() VCP {
	return VCP{code: _DisplayScaling, value: 0x06}
}

func (displayScaling) LinearExpansionCompressionOnHorizontalAxis() VCP {
	return VCP{code: _DisplayScaling, value: 0x07}
}

func (displayScaling) LinearExpansionCompressionOnHAndVAxes() VCP {
	return VCP{code: _DisplayScaling, value: 0x08}
}
func (displayScaling) SqueezeMode() VCP        { return VCP{code: _DisplayScaling, value: 0x09} }
func (displayScaling) NonlinearExpansion() VCP { return VCP{code: _DisplayScaling, value: 0x0a} }

// sharpness
type sharpness struct{}

func (sharpness) FilterFunction1() VCP { return VCP{code: _Sharpness, value: 0x01} }
func (sharpness) FilterFunction2() VCP { return VCP{code: _Sharpness, value: 0x02} }
func (sharpness) FilterFunction3() VCP { return VCP{code: _Sharpness, value: 0x03} }
func (sharpness) FilterFunction4() VCP { return VCP{code: _Sharpness, value: 0x04} }

// tvChannelUpDown
type tvChannelUpDown struct{}

func (tvChannelUpDown) IncrementChannel() VCP { return VCP{code: _TVChannelUpDown, value: 0x01} }
func (tvChannelUpDown) DecrementChannel() VCP { return VCP{code: _TVChannelUpDown, value: 0x02} }

// audioMuteScreenBlank
type audioMuteScreenBlank struct{}

func (audioMuteScreenBlank) MuteTheAudio() VCP { return VCP{code: _AudioMuteScreenBlank, value: 0x01} }
func (audioMuteScreenBlank) UnmuteTheAudio() VCP {
	return VCP{code: _AudioMuteScreenBlank, value: 0x02}
}

// audioProcessorMode
type audioProcessorMode struct{}

func (audioProcessorMode) SpeakerOffAudioNotSupported() VCP {
	return VCP{code: _AudioProcessorMode, value: 0x00}
}
func (audioProcessorMode) Mono() VCP   { return VCP{code: _AudioProcessorMode, value: 0x01} }
func (audioProcessorMode) Stereo() VCP { return VCP{code: _AudioProcessorMode, value: 0x02} }
func (audioProcessorMode) StereoExpanded() VCP {
	return VCP{code: _AudioProcessorMode, value: 0x03}
}
func (audioProcessorMode) SRS20() VCP   { return VCP{code: _AudioProcessorMode, value: 0x11} } // SRS2.0
func (audioProcessorMode) SRS21() VCP   { return VCP{code: _AudioProcessorMode, value: 0x12} } // SRS2.1
func (audioProcessorMode) SRS31() VCP   { return VCP{code: _AudioProcessorMode, value: 0x13} } // SRS3.1
func (audioProcessorMode) SRS41() VCP   { return VCP{code: _AudioProcessorMode, value: 0x14} } // SRS4.1
func (audioProcessorMode) SRS51() VCP   { return VCP{code: _AudioProcessorMode, value: 0x15} } // SRS5.1
func (audioProcessorMode) SRS61() VCP   { return VCP{code: _AudioProcessorMode, value: 0x16} } // SRS6.1
func (audioProcessorMode) SRS71() VCP   { return VCP{code: _AudioProcessorMode, value: 0x17} } // SRS7.1
func (audioProcessorMode) Dolby20() VCP { return VCP{code: _AudioProcessorMode, value: 0x21} } // Dolby2.0
func (audioProcessorMode) Dolby21() VCP { return VCP{code: _AudioProcessorMode, value: 0x22} } // Dolby2.1
func (audioProcessorMode) Dolby31() VCP { return VCP{code: _AudioProcessorMode, value: 0x23} } // Dolby3.1
func (audioProcessorMode) Dolby41() VCP { return VCP{code: _AudioProcessorMode, value: 0x24} } // Dolby4.1
func (audioProcessorMode) Dolby51() VCP { return VCP{code: _AudioProcessorMode, value: 0x25} } // Dolby5.1
func (audioProcessorMode) Dolby61() VCP { return VCP{code: _AudioProcessorMode, value: 0x26} } // Dolby6.1
func (audioProcessorMode) Dolby71() VCP { return VCP{code: _AudioProcessorMode, value: 0x27} } // Dolby7.1
func (audioProcessorMode) THX20() VCP   { return VCP{code: _AudioProcessorMode, value: 0x31} } // THX2.0
func (audioProcessorMode) THX21() VCP   { return VCP{code: _AudioProcessorMode, value: 0x32} } // THX2.1
func (audioProcessorMode) THX31() VCP   { return VCP{code: _AudioProcessorMode, value: 0x33} } // THX3.1
func (audioProcessorMode) THX41() VCP   { return VCP{code: _AudioProcessorMode, value: 0x34} } // THX4.1
func (audioProcessorMode) THX51() VCP   { return VCP{code: _AudioProcessorMode, value: 0x35} } // THX5.1
func (audioProcessorMode) THX61() VCP   { return VCP{code: _AudioProcessorMode, value: 0x36} } // THX6.1
func (audioProcessorMode) THX71() VCP   { return VCP{code: _AudioProcessorMode, value: 0x37} } // THX7.1

// windowControl
type windowControl struct{}

func (windowControl) NoEffect() VCP { return VCP{code: _WindowControl, value: 0x00} }
func (windowControl) Off() VCP      { return VCP{code: _WindowControl, value: 0x01} }
func (windowControl) On() VCP       { return VCP{code: _WindowControl, value: 0x02} }

// autoSetupOnOff
type autoSetupOnOff struct{}

func (autoSetupOnOff) Off() VCP { return VCP{code: _AutoSetupOnOff, value: 0x01} }
func (autoSetupOnOff) On() VCP  { return VCP{code: _AutoSetupOnOff, value: 0x02} }

// changeTheSelectedWindow
type changeTheSelectedWindow struct{}

func (changeTheSelectedWindow) FullDisplayImageAreaSelectedExceptActiveWindows() VCP {
	return VCP{code: _ChangeTheSelectedWindow, value: 0x00}
}

func (changeTheSelectedWindow) Window1Selected() VCP {
	return VCP{code: _ChangeTheSelectedWindow, value: 0x01}
}

func (changeTheSelectedWindow) Window2Selected() VCP {
	return VCP{code: _ChangeTheSelectedWindow, value: 0x02}
}

func (changeTheSelectedWindow) Window3Selected() VCP {
	return VCP{code: _ChangeTheSelectedWindow, value: 0x03}
}

func (changeTheSelectedWindow) Window4Selected() VCP {
	return VCP{code: _ChangeTheSelectedWindow, value: 0x04}
}

func (changeTheSelectedWindow) Window5Selected() VCP {
	return VCP{code: _ChangeTheSelectedWindow, value: 0x05}
}

func (changeTheSelectedWindow) Window6Selected() VCP {
	return VCP{code: _ChangeTheSelectedWindow, value: 0x06}
}

func (changeTheSelectedWindow) Window7Selected() VCP {
	return VCP{code: _ChangeTheSelectedWindow, value: 0x07}
}

// settings
type settings struct{}

func (settings) StoreCurrentSettingsInTheMonitor() VCP { return VCP{code: _Settings, value: 0x01} }
func (settings) RestoreFactoryDefaultsForCurrentMode() VCP {
	return VCP{code: _Settings, value: 0x02}
}

// osdButtonControl
type osdButtonControl struct{}

func (osdButtonControl) Disabled() VCP { return VCP{code: _OSDButtonControl, value: 0x01} }
func (osdButtonControl) Enabled() VCP  { return VCP{code: _OSDButtonControl, value: 0x02} }
func (osdButtonControl) DisplayCannotSupplyThisInformation() VCP {
	return VCP{code: _OSDButtonControl, value: 0xff}
}

// osdLanguage
type osdLanguage struct{}

func (osdLanguage) ReservedValueMustBeIgnored() VCP { return VCP{code: _OSDLanguage, value: 0x00} }
func (osdLanguage) ChineseTraditionalHantai() VCP   { return VCP{code: _OSDLanguage, value: 0x01} }
func (osdLanguage) English() VCP                    { return VCP{code: _OSDLanguage, value: 0x02} }
func (osdLanguage) French() VCP                     { return VCP{code: _OSDLanguage, value: 0x03} }
func (osdLanguage) German() VCP                     { return VCP{code: _OSDLanguage, value: 0x04} }
func (osdLanguage) Italian() VCP                    { return VCP{code: _OSDLanguage, value: 0x05} }
func (osdLanguage) Japanese() VCP                   { return VCP{code: _OSDLanguage, value: 0x06} }
func (osdLanguage) Korean() VCP                     { return VCP{code: _OSDLanguage, value: 0x07} }
func (osdLanguage) PortuguesePortugal() VCP         { return VCP{code: _OSDLanguage, value: 0x08} }
func (osdLanguage) Russian() VCP                    { return VCP{code: _OSDLanguage, value: 0x09} }
func (osdLanguage) Spanish() VCP                    { return VCP{code: _OSDLanguage, value: 0x0a} }
func (osdLanguage) Swedish() VCP                    { return VCP{code: _OSDLanguage, value: 0x0b} }
func (osdLanguage) Turkish() VCP                    { return VCP{code: _OSDLanguage, value: 0x0c} }
func (osdLanguage) ChineseSimplifiedKantai() VCP    { return VCP{code: _OSDLanguage, value: 0x0d} }
func (osdLanguage) PortugueseBrazil() VCP           { return VCP{code: _OSDLanguage, value: 0x0e} }
func (osdLanguage) Arabic() VCP                     { return VCP{code: _OSDLanguage, value: 0x0f} }
func (osdLanguage) Bulgarian() VCP                  { return VCP{code: _OSDLanguage, value: 0x10} }
func (osdLanguage) Croatian() VCP                   { return VCP{code: _OSDLanguage, value: 0x11} }
func (osdLanguage) Czech() VCP                      { return VCP{code: _OSDLanguage, value: 0x12} }
func (osdLanguage) Danish() VCP                     { return VCP{code: _OSDLanguage, value: 0x13} }
func (osdLanguage) Dutch() VCP                      { return VCP{code: _OSDLanguage, value: 0x14} }
func (osdLanguage) Estonian() VCP                   { return VCP{code: _OSDLanguage, value: 0x15} }
func (osdLanguage) Finnish() VCP                    { return VCP{code: _OSDLanguage, value: 0x16} }
func (osdLanguage) Greek() VCP                      { return VCP{code: _OSDLanguage, value: 0x17} }
func (osdLanguage) Hebrew() VCP                     { return VCP{code: _OSDLanguage, value: 0x18} }
func (osdLanguage) Hindi() VCP                      { return VCP{code: _OSDLanguage, value: 0x19} }
func (osdLanguage) Hungarian() VCP                  { return VCP{code: _OSDLanguage, value: 0x1a} }
func (osdLanguage) Latvian() VCP                    { return VCP{code: _OSDLanguage, value: 0x1b} }
func (osdLanguage) Lithuanian() VCP                 { return VCP{code: _OSDLanguage, value: 0x1c} }
func (osdLanguage) Norwegian() VCP                  { return VCP{code: _OSDLanguage, value: 0x1d} }
func (osdLanguage) Polish() VCP                     { return VCP{code: _OSDLanguage, value: 0x1e} }
func (osdLanguage) Romanian() VCP                   { return VCP{code: _OSDLanguage, value: 0x1f} }
func (osdLanguage) Serbian() VCP                    { return VCP{code: _OSDLanguage, value: 0x20} }
func (osdLanguage) Slovak() VCP                     { return VCP{code: _OSDLanguage, value: 0x21} }
func (osdLanguage) Slovenian() VCP                  { return VCP{code: _OSDLanguage, value: 0x22} }
func (osdLanguage) Thai() VCP                       { return VCP{code: _OSDLanguage, value: 0x23} }
func (osdLanguage) Ukranian() VCP                   { return VCP{code: _OSDLanguage, value: 0x24} }
func (osdLanguage) Vietnamese() VCP                 { return VCP{code: _OSDLanguage, value: 0x25} }

// outputSelect
type outputSelect struct{}

func (outputSelect) AnalogVideoRGB1() VCP    { return VCP{code: _OutputSelect, value: 0x01} }
func (outputSelect) AnalogVideoRGB2() VCP    { return VCP{code: _OutputSelect, value: 0x02} }
func (outputSelect) DigitalVideoTDMS1() VCP  { return VCP{code: _OutputSelect, value: 0x03} }
func (outputSelect) DigitalVideoTDMS22() VCP { return VCP{code: _OutputSelect, value: 0x04} }
func (outputSelect) CompositeVideo1() VCP    { return VCP{code: _OutputSelect, value: 0x05} }
func (outputSelect) CompositeVideo2() VCP    { return VCP{code: _OutputSelect, value: 0x06} }
func (outputSelect) SVideo1() VCP            { return VCP{code: _OutputSelect, value: 0x07} }
func (outputSelect) SVideo2() VCP            { return VCP{code: _OutputSelect, value: 0x08} }
func (outputSelect) Tuner1() VCP             { return VCP{code: _OutputSelect, value: 0x09} }
func (outputSelect) Tuner2() VCP             { return VCP{code: _OutputSelect, value: 0x0a} }
func (outputSelect) Tuner3() VCP             { return VCP{code: _OutputSelect, value: 0x0b} }
func (outputSelect) ComponentVideoYPrPbYCrCb1() VCP {
	return VCP{code: _OutputSelect, value: 0x0c}
}

func (outputSelect) ComponentVideoYPrPbYCrCb2() VCP {
	return VCP{code: _OutputSelect, value: 0x0d}
}

func (outputSelect) ComponentVideoYPrPbYCrCb3() VCP {
	return VCP{code: _OutputSelect, value: 0x0e}
}
func (outputSelect) DisplayPort1() VCP { return VCP{code: _OutputSelect, value: 0x0f} }
func (outputSelect) DisplayPort2() VCP { return VCP{code: _OutputSelect, value: 0x10} }
func (outputSelect) HDMI1() VCP        { return VCP{code: _OutputSelect, value: 0x11} }
func (outputSelect) HDMI2() VCP        { return VCP{code: _OutputSelect, value: 0x12} }

// powerMode
type powerMode struct{}

func (powerMode) On() VCP      { return VCP{code: _PowerMode, value: 0x01} }
func (powerMode) Standby() VCP { return VCP{code: _PowerMode, value: 0x02} }
func (powerMode) Suspend() VCP { return VCP{code: _PowerMode, value: 0x03} }
func (powerMode) Off() VCP     { return VCP{code: _PowerMode, value: 0x04} }
func (powerMode) HardOff() VCP { return VCP{code: _PowerMode, value: 0x05} }

// auxiliaryPowerOutput
type auxiliaryPowerOutput struct{}

func (auxiliaryPowerOutput) DisableAuxiliaryPower() VCP {
	return VCP{code: _AuxiliaryPowerOutput, value: 0x01}
}

func (auxiliaryPowerOutput) EnableAuxiliaryPower() VCP {
	return VCP{code: _AuxiliaryPowerOutput, value: 0x02}
}

// scanMode
type scanMode struct{}

func (scanMode) NormalOperation() VCP { return VCP{code: _ScanMode, value: 0x00} }
func (scanMode) UnderScan() VCP       { return VCP{code: _ScanMode, value: 0x01} }
func (scanMode) OverScan() VCP        { return VCP{code: _ScanMode, value: 0x02} }
func (scanMode) Widescreen() VCP      { return VCP{code: _ScanMode, value: 0x03} }

// imageMode
type imageMode struct{}

func (imageMode) NoEffect() VCP    { return VCP{code: _ImageMode, value: 0x00} }
func (imageMode) FullMode() VCP    { return VCP{code: _ImageMode, value: 0x01} }
func (imageMode) ZoomMode() VCP    { return VCP{code: _ImageMode, value: 0x02} }
func (imageMode) SqueezeMode() VCP { return VCP{code: _ImageMode, value: 0x03} }
func (imageMode) Variable() VCP    { return VCP{code: _ImageMode, value: 0x04} }

// displayMode
type displayMode struct{}

func (displayMode) StandardOrDefaultMode() VCP { return VCP{code: _DisplayMode, value: 0x00} }
func (displayMode) Productivity() VCP          { return VCP{code: _DisplayMode, value: 0x01} }
func (displayMode) Mixed() VCP                 { return VCP{code: _DisplayMode, value: 0x02} }
func (displayMode) Movie() VCP                 { return VCP{code: _DisplayMode, value: 0x03} }
func (displayMode) UserDefined() VCP           { return VCP{code: _DisplayMode, value: 0x04} }
func (displayMode) Games() VCP                 { return VCP{code: _DisplayMode, value: 0x05} }
func (displayMode) Sports() VCP                { return VCP{code: _DisplayMode, value: 0x06} }
func (displayMode) ProfessionalAllSignalProcessingDisabled() VCP {
	return VCP{code: _DisplayMode, value: 0x07}
}

func (displayMode) StandardOrDefaultModeWithIntermediatePowerConsumption() VCP {
	return VCP{code: _DisplayMode, value: 0x08}
}

func (displayMode) StandardOrDefaultModeWithLowPowerConsumption() VCP {
	return VCP{code: _DisplayMode, value: 0x09}
}
func (displayMode) Demonstration() VCP   { return VCP{code: _DisplayMode, value: 0x0a} }
func (displayMode) DynamicContrast() VCP { return VCP{code: _DisplayMode, value: 0xf0} }
