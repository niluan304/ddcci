//go:build windows

package ddcci_test

import (
	"fmt"
	"time"

	"github.com/niluan304/ddcci"
	"github.com/niluan304/ddcci/vcp"
)

func ExampleNewPhysicalMonitor() {
	monitors, err := ddcci.NewSystemMonitors()
	if err != nil {
		panic(err)
	}

	m, err := ddcci.NewPhysicalMonitor(&monitors[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(m != nil)

	// Output:
	// true
}

func ExamplePhysicalMonitor0() {
	m := ddcci.PhysicalMonitor0()
	fmt.Println(m != nil)

	// Output:
	// true
}

func ExamplePhysicalMonitor_GetVCPFeatureAndVCPFeatureReply() {
	m := ddcci.PhysicalMonitor0()

	current, maxValue, err := m.GetVCPFeatureAndVCPFeatureReply(vcp.Brightness)
	if err != nil {
		panic(err)
	}

	fmt.Println(maxValue)
	fmt.Println(0 <= current && current <= 100)

	// Output:
	// 100
	// true
}

func ExamplePhysicalMonitor_SetVCPFeature() {
	m := ddcci.PhysicalMonitor0()

	origin, maxValue, err := m.GetVCPFeatureAndVCPFeatureReply(vcp.Brightness)
	if err != nil {
		panic(err)
	}
	defer func() {
		time.Sleep(time.Second)
		m.SetVCPFeature(vcp.Brightness.Value(origin))
	}()

	m.SetVCPFeature(vcp.Brightness.Value(50))
	current, _, _ := m.GetVCPFeatureAndVCPFeatureReply(vcp.Brightness)
	fmt.Println(maxValue)
	fmt.Println(current)

	// Output:
	// 100
	// 50
}

func ExamplePhysicalMonitor_GetBrightness() {
	m := ddcci.PhysicalMonitor0()

	minValue, current, maxValue, err := m.GetBrightness()
	if err != nil {
		panic(err)
	}

	fmt.Println(minValue)
	fmt.Println(maxValue)
	fmt.Println(0 <= current && current <= 100)

	// Output:
	// 0
	// 100
	// true
}

func ExamplePhysicalMonitor_GetContrast() {
	m := ddcci.PhysicalMonitor0()

	minValue, current, maxValue, err := m.GetContrast()
	if err != nil {
		panic(err)
	}

	fmt.Println(minValue)
	fmt.Println(maxValue)
	fmt.Println(0 <= current && current <= 100)

	// Output:
	// 0
	// 100
	// true
}

func ExamplePhysicalMonitor_SetBrightness() {
	m := ddcci.PhysicalMonitor0()

	minValue, origin, maxValue, err := m.GetBrightness()
	if err != nil {
		panic(err)
	}
	defer func() {
		time.Sleep(time.Second)
		m.SetVCPFeature(vcp.Brightness.Value(origin))
	}()

	m.SetBrightness(50)
	_, current, _, _ := m.GetBrightness()

	fmt.Println(minValue)
	fmt.Println(maxValue)
	fmt.Println(current)

	// output:
	// 0
	// 100
	// 50
}

func ExamplePhysicalMonitor_SetContrast() {
	m := ddcci.PhysicalMonitor0()

	minValue, origin, maxValue, err := m.GetContrast()
	if err != nil {
		panic(err)
	}
	defer func() {
		time.Sleep(time.Second)
		m.SetVCPFeature(vcp.Contrast.Value(origin))
	}()

	m.SetContrast(50)
	_, current, _, _ := m.GetContrast()

	fmt.Println(minValue)
	fmt.Println(maxValue)
	fmt.Println(current)

	// output:
	// 0
	// 100
	// 50
}
