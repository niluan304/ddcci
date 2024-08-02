package vcp

// vcp code with readwrite
type readwrite int

func (code readwrite) Code() int           { return int(code) }
func (code readwrite) Value(value int) VCP { return VCP{code: int(code), value: value} }

// Attributes: Read Write, Continuous
const (
	// ColorTemperatureRequest
	// Specifies a color temperature (degrees Kelvin)
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: COLOR
	// Attributes:
	ColorTemperatureRequest readwrite = _ColorTemperatureRequest

	// Clock
	// Increase/decrease the sampling clock frequency.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets:
	// Attributes:
	Clock readwrite = _Clock

	// Brightness
	// Increase/decrease the brightness of the image.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	Brightness readwrite = _Brightness

	// Contrast
	// Increase/decrease the contrast of the image.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	Contrast readwrite = _Contrast

	// BacklightControl
	// Increase/decrease the specified backlight control value
	// MCCS versions: 2.1, 3.0
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes (v2.1): Read Write, Continuous (complex)
	// Attributes (v3.0):
	BacklightControl readwrite = _BacklightControl

	// VideoGainRed
	// Increase/decrease the luminesence of red pixels
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	VideoGainRed readwrite = _VideoGainRed

	// UserColorVisionCompensation
	// Increase/decrease the degree of compensation
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: COLOR
	// Attributes:
	UserColorVisionCompensation readwrite = _UserColorVisionCompensation

	// VideoGainGreen
	// Increase/decrease the luminesence of green pixels
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	VideoGainGreen readwrite = _VideoGainGreen

	// VideoGainBlue
	// Increase/decrease the luminesence of blue pixels
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	VideoGainBlue readwrite = _VideoGainBlue

	// Focus
	// Increase/decrease the focus of the image
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets:
	// Attributes:
	Focus readwrite = _Focus

	// GrayScaleExpansion
	// Gray Scale Expansion
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: COLOR
	// Attributes: Read Write, Non-Continuous (complex)
	GrayScaleExpansion readwrite = _GrayScaleExpansion

	// ClockPhase
	// Increase/decrease the sampling clock phase shift
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Miscellaneous
	// ddcutil feature subsets:
	// Attributes:
	ClockPhase readwrite = _ClockPhase

	// Rotation
	// Increasing (decreasing) this value rotates the image (counter) clockwise around the center point of the screen.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Geometry
	// ddcutil feature subsets: CRT
	// Attributes:
	Rotation readwrite = _Rotation

	// TopCornerFlare
	// Increase/decrease the distance between the left and right sides at the top of the image.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Geometry
	// ddcutil feature subsets: CRT
	// Attributes:
	TopCornerFlare readwrite = _TopCornerFlare

	// TopCornerHook
	// Increasing (decreasing) this value moves the top of the image to the right (left).
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Geometry
	// ddcutil feature subsets: CRT
	// Attributes:
	TopCornerHook readwrite = _TopCornerHook

	// BottomCornerFlare
	// Increase/decrease the distance between the left and right sides at the bottom of the image.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Geometry
	// ddcutil feature subsets: CRT
	// Attributes:
	BottomCornerFlare readwrite = _BottomCornerFlare

	// BottomCornerHook
	// Increasing (decreasing) this value moves the bottom end of the image to the right (left).
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Geometry
	// ddcutil feature subsets: CRT
	// Attributes:
	BottomCornerHook readwrite = _BottomCornerHook

	// AudioSpeakerVolume
	// Adjusts speaker volume
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Audio
	// ddcutil feature subsets: AUDIO
	// Attributes (v2.0): Read Write, Continuous (normal)
	// Attributes (v2.1):
	AudioSpeakerVolume readwrite = _AudioSpeakerVolume

	// AudioMicrophoneVolume
	// Increase/decrease microphone gain
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Audio
	// ddcutil feature subsets: AUDIO
	// Attributes:
	AudioMicrophoneVolume readwrite = _AudioMicrophoneVolume

	// BacklightLevelWhite
	// Increase/decrease the white backlight level
	// MCCS versions: 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	BacklightLevelWhite readwrite = _BacklightLevelWhite

	// VideoBlackLevelRed
	// Increase/decrease the black level of red pixels
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	VideoBlackLevelRed readwrite = _VideoBlackLevelRed

	// BacklightLevelRed
	// Increase/decrease the red backlight level
	// MCCS versions: 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	BacklightLevelRed readwrite = _BacklightLevelRed

	// VideoBlackLevelGreen
	// Increase/decrease the black level of green pixels
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	VideoBlackLevelGreen readwrite = _VideoBlackLevelGreen

	// BacklightLevelGreen
	// Increase/decrease the green backlight level
	// MCCS versions: 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	BacklightLevelGreen readwrite = _BacklightLevelGreen

	// VideoBlackLevelBlue
	// Increase/decrease the black level of blue pixels
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	VideoBlackLevelBlue readwrite = _VideoBlackLevelBlue

	// BacklightLevelBlue
	// Increase/decrease the blue backlight level
	// MCCS versions: 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: PROFILE, COLOR
	// Attributes:
	BacklightLevelBlue readwrite = _BacklightLevelBlue

	// AdjustFocalPlane
	// Increase/decrease the distance to the focal plane of the image
	// MCCS versions: 2.0, 2.1
	// MCCS specification groups: Image
	// ddcutil feature subsets:
	// Attributes (v2.0): Read Write, Continuous (normal)
	// Attributes (v2.1):
	AdjustFocalPlane readwrite = _AdjustFocalPlane

	// AdjustZoom
	// Increase/decrease the distance to the zoom function of the projection lens (optics)
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets:
	// Attributes:
	AdjustZoom readwrite = _AdjustZoom

	// Trapezoid
	// Increase/decrease the trapezoid distortion in the image
	// MCCS versions: 2.0, 2.1
	// MCCS specification groups: Geometry
	// ddcutil feature subsets: CRT
	// Attributes (v2.0): Read Write, Continuous (normal)
	// Attributes (v2.1):
	Trapezoid readwrite = _Trapezoid

	// Keystone
	// Increase/decrease the keystone distortion in the image.
	// MCCS versions: 2.0
	// MCCS specification groups: Geometry
	// ddcutil feature subsets: CRT
	// Attributes (v2.0):
	Keystone readwrite = _Keystone

	// VelocityScanModulation
	// Increase (decrease) the velocity modulation of the horizontal scan as a function of the change in luminescence level
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Miscellaneous
	// ddcutil feature subsets: CRT
	// Attributes:
	VelocityScanModulation readwrite = _VelocityScanModulation

	// ColorSaturation
	// Increase/decrease the amplitude of the color difference components of the video signal
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: COLOR, TV
	// Attributes:
	ColorSaturation readwrite = _ColorSaturation

	// TVSharpness
	// Increase/decrease the amplitude of the high frequency components  of the video signal
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Miscellaneous
	// ddcutil feature subsets: TV
	// Attributes:
	TVSharpness readwrite = _TVSharpness

	// TVContrast
	// Increase/decrease the ratio between blacks and whites in the image
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets: TV
	// Attributes:
	TVContrast readwrite = _TVContrast

	// AudioTreble
	// Emphasize/de-emphasize high frequency audio
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Audio
	// ddcutil feature subsets: AUDIO
	// Attributes (v2.0): Read Write, Continuous (normal)
	// Attributes (v2.1):
	AudioTreble readwrite = _AudioTreble

	// Hue
	// Increase/decrease the wavelength of the color component of the video signal. AKA tint.  Applies to currently active interface
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets: COLOR, TV
	// Attributes:
	Hue readwrite = _Hue

	// AudioBass
	// Emphasize/de-emphasize low frequency audio
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Audio
	// ddcutil feature subsets: AUDIO
	// Attributes (v2.0): Read Write, Continuous (normal)
	// Attributes (v2.1):
	AudioBass readwrite = _AudioBass

	// TVBlackLevelLuminesence
	// Increase/decrease the black level of the video
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets: TV
	// Attributes:
	TVBlackLevelLuminesence readwrite = _TVBlackLevelLuminesence

	// AudioBalanceLR
	// Controls left/right audio balance
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Audio
	// ddcutil feature subsets: AUDIO
	// Attributes (v2.0): Read Write, Continuous (normal)
	// Attributes (v2.1):
	AudioBalanceLR readwrite = _AudioBalanceLR

	// WindowBackground
	// Changes the contrast ratio between the area of the window and the rest of the desktop
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Window
	// ddcutil feature subsets: WINDOW
	// Attributes:
	WindowBackground readwrite = _WindowBackground

	// PacketCount
	// Counter for DPVL packets received
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: DPVL
	// ddcutil feature subsets: DPVL
	// Attributes:
	PacketCount readwrite = _PacketCount

	// MonitorXOrigin
	// X origin of the monitor in the vertical screen
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: DPVL
	// ddcutil feature subsets: DPVL
	// Attributes:
	MonitorXOrigin readwrite = _MonitorXOrigin

	// MonitorYOrigin
	// Y origin of the monitor in the vertical screen
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: DPVL
	// ddcutil feature subsets: DPVL
	// Attributes:
	MonitorYOrigin readwrite = _MonitorYOrigin

	// HeaderErrorCount
	// Error counter for the DPVL header
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: DPVL
	// ddcutil feature subsets: DPVL
	// Attributes:
	HeaderErrorCount readwrite = _HeaderErrorCount

	// BodyCRCErrorCount
	// CRC error counter for the DPVL body
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: DPVL
	// ddcutil feature subsets: DPVL
	// Attributes:
	BodyCRCErrorCount readwrite = _BodyCRCErrorCount

	// ClientID
	// Assigned identification number for the monitor
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: DPVL
	// ddcutil feature subsets: DPVL
	// Attributes:
	ClientID readwrite = _ClientID
)

const (
	// FleshToneEnhancement
	// Select contrast enhancement algorithm respecting flesh tone region
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: COLOR
	FleshToneEnhancement readwrite = _FleshToneEnhancement

	// PerformancePreservation
	// Controls features aimed at preserving display performance
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	PerformancePreservation readwrite = _PerformancePreservation

	// Gamma
	// Select relative or absolute gamma
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: COLOR
	Gamma readwrite = _Gamma

	// SinglePointLUTOperation
	// Writes a single point within the display's LUT, reads a single point from the LUT
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Miscellaneous
	// ddcutil feature subsets: LUT
	SinglePointLUTOperation readwrite = _SinglePointLUTOperation

	// BlockLUTOperation
	// Load (read) multiple values into (from) the display's LUT
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Miscellaneous
	// ddcutil feature subsets: LUT
	BlockLUTOperation readwrite = _BlockLUTOperation

	// WindowMaskControl
	// Turn selected window operation on/off, window mask
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image
	// ddcutil feature subsets: WINDOW
	// Attributes (v2.0): Read Write, Non-Continuous (complex)
	// Attributes (v2.1): Read Write, Non-Continuous (complex)
	// Attributes (v3.0): Read Write, Table (normal)
	WindowMaskControl readwrite = _WindowMaskControl

	// SourceTimingMode
	// Indicates timing mode being sent by host
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Control
	// ddcutil feature subsets:
	// Attributes (v2.1): Read Write, Non-Continuous (complex)
	// Attributes (v3.0): Read Write, Table (normal)
	SourceTimingMode readwrite = _SourceTimingMode

	// LinkControl
	// Indicates status of the DVI link
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: DPVL
	// ddcutil feature subsets: DPVL
	LinkControl readwrite = _LinkControl

	// TransmitDisplayDescriptor
	// Reads (writes) a display descriptor from (to) non-volatile storage in the display.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	TransmitDisplayDescriptor readwrite = _TransmitDisplayDescriptor

	// EnableDisplayOfDisplayDescriptor
	// If enabled, the display descriptor shall be displayed when no video is being received.
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	EnableDisplayOfDisplayDescriptor readwrite = _EnableDisplayOfDisplayDescriptor

	// StatusIndicators
	// Control up to 16 LED (or similar) indicators to indicate system status
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	StatusIndicators readwrite = _StatusIndicators

	// AssetTag
	// Read an Asset Tag to/from the display
	// MCCS versions: 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	AssetTag readwrite = _AssetTag

	// StereoVideoMode
	// Stereo video mode
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Image, Miscellaneous
	StereoVideoMode readwrite = _StereoVideoMode

	// ScratchPad
	// Operation mode (2.0) or scratch pad (3.0/2.2)
	// MCCS versions: 2.0, 2.1, 3.0, 2.2
	// MCCS specification groups: Miscellaneous
	// ddcutil feature subsets:
	// Attributes (v2.0): Write Only, Non-Continuous (write-only)
	// Attributes (v2.1): Read Write, Non-Continuous (complex)
	// Attributes (v3.0): Read Write, Non-Continuous (complex)
	ScratchPad readwrite = _ScratchPad
)
