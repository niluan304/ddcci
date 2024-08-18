package vcp

type Coder interface{ Code() int }

type VCP struct {
	code  int
	value int
}

func NewVCP(code, value int) VCP { return VCP{code: code, value: value} }

func (v VCP) Code() int  { return v.code }
func (v VCP) Value() int { return v.value }

// [ddcutil vcpinfo] provides a quick list of all features defined in the Monitor
//
// [ddcutil vcpinfo]: https://www.ddcutil.com/vcpinfo_output/
const (
	_Degauss                                  = 0x01
	_NewControlValue                          = 0x02
	_SoftControls                             = 0x03
	_RestoreFactoryDefaults                   = 0x04
	_RestoreFactoryBrightnessContrastDefaults = 0x05
	_RestoreFactoryGeometryDefaults           = 0x06
	_RestoreColorDefaults                     = 0x08
	_RestoreFactoryTVDefaults                 = 0x0A
	_ColorTemperatureIncrement                = 0x0B
	_ColorTemperatureRequest                  = 0x0C
	_Clock                                    = 0x0E
	_Brightness                               = 0x10
	_FleshToneEnhancement                     = 0x11
	_Contrast                                 = 0x12
	_BacklightControl                         = 0x13
	_SelectColorPreset                        = 0x14
	_VideoGainRed                             = 0x16
	_UserColorVisionCompensation              = 0x17
	_VideoGainGreen                           = 0x18
	_VideoGainBlue                            = 0x1A
	_Focus                                    = 0x1C
	_AutoSetup                                = 0x1E
	_AutoColorSetup                           = 0x1F
	_HorizontalPositionPhase                  = 0x20
	_HorizontalSize                           = 0x22
	_HorizontalPincushion                     = 0x24
	_HorizontalPincushionBalance              = 0x26
	_HorizontalConvergenceRB                  = 0x28
	_HorizontalConvergenceMG                  = 0x29
	_HorizontalLinearity                      = 0x2A
	_HorizontalLinearityBalance               = 0x2C
	_GrayScaleExpansion                       = 0x2E
	_VerticalPositionPhase                    = 0x30
	_VerticalSize                             = 0x32
	_VerticalPincushion                       = 0x34
	_VerticalPincushionBalance                = 0x36
	_VerticalConvergenceRB                    = 0x38
	_VerticalConvergenceMG                    = 0x39
	_VerticalLinearity                        = 0x3A
	_VerticalLinearityBalance                 = 0x3C
	_ClockPhase                               = 0x3E
	_HorizontalParallelogram                  = 0x40
	_VerticalParallelogram                    = 0x41
	_HorizontalKeystone                       = 0x42
	_VerticalKeystone                         = 0x43
	_Rotation                                 = 0x44
	_TopCornerFlare                           = 0x46
	_TopCornerHook                            = 0x48
	_BottomCornerFlare                        = 0x4A
	_BottomCornerHook                         = 0x4C
	_ActiveControl                            = 0x52
	_PerformancePreservation                  = 0x54
	_HorizontalMoire                          = 0x56
	_VerticalMoire                            = 0x58
	_SixAxisSaturationRed                     = 0x59
	_SixAxisSaturationYellow                  = 0x5A
	_SixAxisSaturationGreen                   = 0x5B
	_SixAxisSaturationCyan                    = 0x5C
	_SixAxisSaturationBlue                    = 0x5D
	_SixAxisSaturationMagenta                 = 0x5E
	_InputSource                              = 0x60
	_AudioSpeakerVolume                       = 0x62
	_SpeakerSelect                            = 0x63
	_AudioMicrophoneVolume                    = 0x64
	_AmbientLightSensor                       = 0x66
	_BacklightLevelWhite                      = 0x6B
	_VideoBlackLevelRed                       = 0x6C
	_BacklightLevelRed                        = 0x6D
	_VideoBlackLevelGreen                     = 0x6E
	_BacklightLevelGreen                      = 0x6F
	_VideoBlackLevelBlue                      = 0x70
	_BacklightLevelBlue                       = 0x71
	_Gamma                                    = 0x72
	_LUTSize                                  = 0x73
	_SinglePointLUTOperation                  = 0x74
	_BlockLUTOperation                        = 0x75
	_RemoteProcedureCall                      = 0x76
	_DisplayIdentificationOperation           = 0x78
	_AdjustFocalPlane                         = 0x7A
	_AdjustZoom                               = 0x7C
	_Trapezoid                                = 0x7E
	_Keystone                                 = 0x80
	_HorizontalMirrorFlip                     = 0x82
	_VerticalMirrorFlip                       = 0x84
	_DisplayScaling                           = 0x86
	_Sharpness                                = 0x87
	_VelocityScanModulation                   = 0x88
	_ColorSaturation                          = 0x8A
	_TVChannelUpDown                          = 0x8B
	_TVSharpness                              = 0x8C
	_AudioMuteScreenBlank                     = 0x8D
	_TVContrast                               = 0x8E
	_AudioTreble                              = 0x8F
	_Hue                                      = 0x90
	_AudioBass                                = 0x91
	_TVBlackLevelLuminesence                  = 0x92
	_AudioBalanceLR                           = 0x93
	_AudioProcessorMode                       = 0x94
	_WindowPositionTLX                        = 0x95
	_WindowPositionTLY                        = 0x96
	_WindowPositionBRX                        = 0x97
	_WindowPositionBRY                        = 0x98
	_WindowControl                            = 0x99
	_WindowBackground                         = 0x9A
	_SixAxisHueControlRed                     = 0x9B
	_SixAxisHueControlYellow                  = 0x9C
	_SixAxisHueControlGreen                   = 0x9D
	_SixAxisHueControlCyan                    = 0x9E
	_SixAxisHueControlBlue                    = 0x9F
	_SixAxisHueControlMagenta                 = 0xA0
	_AutoSetupOnOff                           = 0xA2
	_WindowMaskControl                        = 0xA4
	_ChangeTheSelectedWindow                  = 0xA5
	_ScreenOrientation                        = 0xAA
	_HorizontalFrequency                      = 0xAC
	_VerticalFrequency                        = 0xAE
	_Settings                                 = 0xB0
	_FlatPanelSubPixelLayout                  = 0xB2
	_SourceTimingMode                         = 0xB4
	_DisplayTechnologyType                    = 0xB6
	_MonitorStatus                            = 0xB7
	_PacketCount                              = 0xB8
	_MonitorXOrigin                           = 0xB9
	_MonitorYOrigin                           = 0xBA
	_HeaderErrorCount                         = 0xBB
	_BodyCRCErrorCount                        = 0xBC
	_ClientID                                 = 0xBD
	_LinkControl                              = 0xBE
	_DisplayUsageTime                         = 0xC0
	_DisplayDescriptorLength                  = 0xC2
	_TransmitDisplayDescriptor                = 0xC3
	_EnableDisplayOfDisplayDescriptor         = 0xC4
	_ApplicationEnableKey                     = 0xC6
	_DisplayControllerType                    = 0xC8
	_DisplayFirmwareLevel                     = 0xC9
	_OSDButtonControl                         = 0xCA
	_OSDLanguage                              = 0xCC
	_StatusIndicators                         = 0xCD
	_AuxiliaryDisplaySize                     = 0xCE
	_AuxiliaryDisplayData                     = 0xCF
	_OutputSelect                             = 0xD0
	_AssetTag                                 = 0xD2
	_StereoVideoMode                          = 0xD4
	_PowerMode                                = 0xD6
	_AuxiliaryPowerOutput                     = 0xD7
	_ScanMode                                 = 0xDA
	_ImageMode                                = 0xDB
	_DisplayMode                              = 0xDC
	_ScratchPad                               = 0xDE
	_Version                                  = 0xDF
)
