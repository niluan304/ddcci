package vcp

var (
	// WindowPosition 允许用户或程序指定图像区域在屏幕上的位置，通常用于调整或定位显示内容。
	WindowPosition windowPosition

	// SixAxisSaturation 调节显示器六轴饱和度
	SixAxisSaturation saturation
	// SixAxisHueControl 控制显示器色彩六轴色调
	SixAxisHueControl hueControl

	// Horizontal 水平调节
	Horizontal horizontal
	// Vertical 垂直调节
	Vertical vertical
)

// ======================= windowPosition =======================

// 允许用户或程序指定图像区域在屏幕上的位置，通常用于调整或定位显示内容。
type windowPosition struct{}

// TLX
//
// Window Position(TL_X)
// Top left X pixel of an area of the image
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry, Window
// ddcutil feature subsets: WINDOW
// Attributes: Read Write, Continuous (normal)
func (w windowPosition) TLX(value int) VCP { return VCP{code: _WindowPositionTLX, value: value} }

// TLY
//
// Window Position(TL_Y)
// Top left Y pixel of an area of the image
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry, Window
// ddcutil feature subsets: WINDOW
// Attributes: Read Write, Continuous (normal)
func (w windowPosition) TLY(value int) VCP { return VCP{code: _WindowPositionTLY, value: value} }

// BRX
//
// Window Position(BR_X)
// Bottom right X pixel of an area of the image
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry, Window
// ddcutil feature subsets: WINDOW
// Attributes: Read Write, Continuous (normal)
func (w windowPosition) BRX(value int) VCP { return VCP{code: _WindowPositionBRX, value: value} }

// BRY
//
// Window Position(BR_Y)
// Bottom right Y pixel of an area of the image
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry, Window
// ddcutil feature subsets: WINDOW
// Attributes: Read Write, Continuous (normal)
func (w windowPosition) BRY(value int) VCP { return VCP{code: _WindowPositionBRY, value: value} }

// ======================= saturation =======================
// 调节显示器六轴饱和度
type saturation struct{}

func (s saturation) Red(value int) VCP     { return VCP{code: _SixAxisSaturationRed, value: value} }
func (s saturation) Yellow(value int) VCP  { return VCP{code: _SixAxisSaturationYellow, value: value} }
func (s saturation) Green(value int) VCP   { return VCP{code: _SixAxisSaturationGreen, value: value} }
func (s saturation) Cyan(value int) VCP    { return VCP{code: _SixAxisSaturationCyan, value: value} }
func (s saturation) Blue(value int) VCP    { return VCP{code: _SixAxisSaturationBlue, value: value} }
func (s saturation) Magenta(value int) VCP { return VCP{code: _SixAxisSaturationMagenta, value: value} }

// ======================= hueControl =======================

// 控制显示器色彩六轴色调
type hueControl struct{}

func (c hueControl) Red(value int) VCP     { return VCP{code: _SixAxisHueControlRed, value: value} }
func (c hueControl) Yellow(value int) VCP  { return VCP{code: _SixAxisHueControlYellow, value: value} }
func (c hueControl) Green(value int) VCP   { return VCP{code: _SixAxisHueControlGreen, value: value} }
func (c hueControl) Cyan(value int) VCP    { return VCP{code: _SixAxisHueControlCyan, value: value} }
func (c hueControl) Blue(value int) VCP    { return VCP{code: _SixAxisHueControlBlue, value: value} }
func (c hueControl) Magenta(value int) VCP { return VCP{code: _SixAxisHueControlMagenta, value: value} }

// ======================= horizontal =======================

// 水平调节
type horizontal struct{}

// PositionPhase
//
// Increasing (decreasing) this value moves the image toward the right (left) of the display.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) PositionPhase(value int) VCP {
	return VCP{code: _HorizontalPositionPhase, value: value}
}

// Size
//
// Increase/decrease the width of the image.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) Size(value int) VCP { return VCP{code: _HorizontalSize, value: value} }

// Pincushion
//
// Increasing (decreasing) this value causes the right and left sides of the image to become more (less) convex.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) Pincushion(value int) VCP { return VCP{code: _HorizontalPincushion, value: value} }

// PincushionBalance
//
// Increasing (decreasing) this value moves the center section of the image toward the right (left) side of the display.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) PincushionBalance(value int) VCP {
	return VCP{code: _HorizontalPincushionBalance, value: value}
}

// ConvergenceRB
//
// Increasing (decreasing) this value shifts the red pixels to the right (left) and the blue pixels left (right) across the image with respect to the green pixels.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) ConvergenceRB(value int) VCP {
	return VCP{code: _HorizontalConvergenceRB, value: value}
}

// ConvergenceMG
//
// Increasing (decreasing) this value shifts the magenta pixels to the right (left) and the green pixels left (right) across the image with respect to the magenta (sic) pixels.
// MCCS versions: 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) ConvergenceMG(value int) VCP {
	return VCP{code: _HorizontalConvergenceMG, value: value}
}

// Linearity
//
// Increase/decrease the density of pixels in the image center.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) Linearity(value int) VCP { return VCP{code: _HorizontalLinearity, value: value} }

// LinearityBalance
//
// Increasing (decreasing) this value shifts the density of pixels from the left (right) side to the right (left) side of the image.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) LinearityBalance(value int) VCP {
	return VCP{code: _HorizontalLinearityBalance, value: value}
}

// Parallelogram
//
// Increasing (decreasing) this value shifts the top section of the image to the right (left) with respect to the bottom section of the image.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) Parallelogram(value int) VCP {
	return VCP{code: _HorizontalParallelogram, value: value}
}

// Keystone
//
// Increasing (decreasing) this value will increase (decrease) the ratio between the horizontal size at the top of the image and the horizontal size at the bottom of the image.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (h horizontal) Keystone(value int) VCP { return VCP{code: _HorizontalKeystone, value: value} }

// Moire
//
// Increase/decrease horizontal moire cancellation.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Image
// ddcutil feature subsets:
// Attributes:
func (h horizontal) Moire(value int) VCP { return VCP{code: _HorizontalMoire, value: value} }

// ======================= vertical =======================

// 垂直调节
type vertical struct{}

// PositionPhase
//
// Increasing (decreasing) this value moves the image toward the top (bottom) edge of the display.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) PositionPhase(value int) VCP {
	return VCP{code: _VerticalPositionPhase, value: value}
}

// Size
//
// Increase/decreasing the height of the image.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) Size(value int) VCP { return VCP{code: _VerticalSize, value: value} }

// Pincushion
//
// Increasing (decreasing) this value will cause the top and bottom edges of the image to become more (less) convex.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) Pincushion(value int) VCP { return VCP{code: _VerticalPincushion, value: value} }

// PincushionBalance
//
// Increasing (decreasing) this value will move the center section of the image toward the top (bottom) edge of the display.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) PincushionBalance(value int) VCP {
	return VCP{code: _VerticalPincushionBalance, value: value}
}

// ConvergenceRB
//
// Increasing (decreasing) this value shifts the red pixels up (down) across the image and the blue pixels down (up) across the image with respect to the green pixels.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) ConvergenceRB(value int) VCP {
	return VCP{code: _VerticalConvergenceRB, value: value}
}

// ConvergenceMG
// Increasing (decreasing) this value shifts the magenta pixels up (down) across the image and the green pixels down (up) across the image with respect to the magenta (sic) pixels.
// MCCS versions: 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) ConvergenceMG(value int) VCP {
	return VCP{code: _VerticalConvergenceMG, value: value}
}

// Linearity
//
// Increase/decease the density of scan lines in the image center.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) Linearity(value int) VCP { return VCP{code: _VerticalLinearity, value: value} }

// LinearityBalance
//
// Increase/decrease the density of scan lines in the image center.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) LinearityBalance(value int) VCP {
	return VCP{code: _VerticalLinearityBalance, value: value}
}

// Parallelogram
//
// Increasing (decreasing) this value shifts the top section of the image to the right (left) with respect to the bottom section of the image. (sic)
// MCCS versions: 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) Parallelogram(value int) VCP {
	return VCP{code: _VerticalParallelogram, value: value}
}

// Keystone
//
// Increasing (decreasing) this value will increase (decrease) the ratio between the vertical size at the left of the image and the vertical size at the right of the image.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Geometry
// ddcutil feature subsets: CRT
// Attributes:
func (v vertical) Keystone(value int) VCP { return VCP{code: _VerticalKeystone, value: value} }

// Moire
//
// Increase/decrease vertical moire cancellation.
// MCCS versions: 2.0, 2.1, 3.0, 2.2
// MCCS specification groups: Image
// ddcutil feature subsets:
// Attributes:
func (v vertical) Moire(value int) VCP { return VCP{code: _VerticalMoire, value: value} }
