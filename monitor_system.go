//go:build windows

package ddcci

import (
	"errors"
	"fmt"
	"slices"
	"syscall"
	"unsafe"
)

const errno0 = syscall.Errno(0)

// The RECT structure defines a rectangle by the coordinates of its upper-left and lower-right corners.
//
// See [microsoft-windows doc]
//
// [microsoft-windows doc]: https://learn.microsoft.com/en-us/windows/win32/api/windef/ns-windef-rect
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// SystemMonitor
//
// See [microsoft-windows doc]
//
// [microsoft-windows doc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-monitorenumproc
type SystemMonitor struct {
	user32 *syscall.DLL
	dxva2  *syscall.DLL

	// A handle to the display monitor. This value will always be non-NULL.
	hMonitor syscall.Handle

	// A handle to a device context.
	//
	// The device context has color attributes that are appropriate for
	// the display monitor identified by hMonitor.
	//
	// The clipping area of the device context is set to the intersection of
	// the visible region of the device context identified by the hdc parameter of EnumDisplayMonitors,
	// the rectangle pointed to by the lprcClip parameter of EnumDisplayMonitors,
	// and the display monitor rectangle.
	//
	// This value is NULL if the hdc parameter of EnumDisplayMonitors was NULL.
	hdc syscall.Handle

	// A pointer to a RECT structure.
	//
	// If hdcMonitor is non-NULL, this rectangle is the intersection of the clipping area
	// of the device context identified by hdcMonitor and the display monitor rectangle.
	// The rectangle coordinates are device-context coordinates.
	//
	// If hdcMonitor is NULL, this rectangle is the display monitor rectangle.
	// The rectangle coordinates are virtual-screen coordinates.
	rect *RECT

	// Application-defined data that EnumDisplayMonitors passes directly to the enumeration function.
	data uintptr
}

// NewSystemMonitors
// The EnumDisplayMonitors function enumerates display monitors (including
// invisible pseudo-monitors associated with the mirroring drivers)
// that intersect a region formed by the intersection of a specified clipping rectangle
// and the visible region of a device context.
//
// EnumDisplayMonitors calls an application-defined MonitorEnumProc callback function once
// for each monitor that is enumerated.
// Note that GetSystemMetrics (SM_CMONITORS) counts only the display monitors.
//
// A value of type MONITORENUMPROC is a pointer to a MonitorEnumProc function.
//
// See [microsoft-windows doc]
//
// [microsoft-windows doc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func NewSystemMonitors() ([]SystemMonitor, error) {
	user32, err := syscall.LoadDLL("user32.dll")
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to load dll(%s)", "user32.dll"))
	}
	proc, err := user32.FindProc(EnumDisplayMonitors)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("fail to load proc(%s)", EnumDisplayMonitors))
	}

	dxva2, err := syscall.LoadDLL("dxva2.dll")
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to load dll(%s)", "dxva2.dll"))
	}

	var monitors []SystemMonitor
	callback := func(hMonitor, hdc syscall.Handle, rect *RECT, data uintptr) uintptr {
		monitors = append(monitors, SystemMonitor{
			user32:   user32,
			dxva2:    dxva2,
			hMonitor: hMonitor,
			hdc:      hdc,
			rect:     rect,
			data:     data,
		})

		// enumerate the next monitor
		return 1
	}

	const handle0 = uintptr(syscall.Handle(0))
	_, _, errno := proc.Call(
		handle0,
		handle0,
		syscall.NewCallback(callback),
		handle0,
	)
	if !errors.Is(errno, errno0) {
		return nil, errors.Join(errno, errors.New("get system monitors"))
	}
	return monitors, nil
}

func (m *SystemMonitor) dxva2Call(procName string, args ...uintptr) error {
	proc, err := m.dxva2.FindProc(procName)
	if err != nil {
		return errors.Join(err, fmt.Errorf("fail to find proc(%s)", procName))
	}

	args = slices.Insert(args, 0, uintptr(m.hMonitor))
	_, _, errno := proc.Call(args...)
	if !errors.Is(errno, errno0) {
		return errors.Join(errno, fmt.Errorf("fail to call proc(%s)", procName))
	}

	return nil
}

// GetNumberOfPhysicalMonitorsFromHMONITOR
// Retrieves the number of physical monitors associated with an HMONITOR monitor handle.
// Call this function before calling GetPhysicalMonitorsFromHMONITOR.
//
// [out] num
// Receives the number of physical monitors associated with the monitor handle.
func (m *SystemMonitor) GetNumberOfPhysicalMonitorsFromHMONITOR() (num int, err error) {
	err = m.dxva2Call(GetNumberOfPhysicalMonitorsFromHMONITOR,
		uintptr(unsafe.Pointer(&num)),
	)
	if err != nil {
		return 0, err
	}

	return num, nil
}
