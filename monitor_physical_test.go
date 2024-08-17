//go:build windows

package ddcci_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/niluan304/ddcci"
	"github.com/niluan304/ddcci/vcp"
)

func Test_PhysicalMonitor(t *testing.T) {
	type args struct {
		m *ddcci.PhysicalMonitor
		f func(m *ddcci.PhysicalMonitor) error
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "PowerMode Off",
			args: args{
				m: ddcci.PhysicalMonitor0(),
				f: func(m *ddcci.PhysicalMonitor) error {
					err := m.SetVCPFeature(vcp.PowerMode.Off())
					if err != nil {
						return err
					}

					time.Sleep(2 * time.Second)

					err = m.SetVCPFeature(vcp.PowerMode.On())
					if err != nil {
						return err
					}

					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.args.m

			if err := tt.args.f(m); (err != nil) != tt.wantErr {
				t.Errorf("PhysicalMonitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})

		time.Sleep(5 * time.Second)
	}
}

func TestPhysicalMonitor_GetVCPFeatureAndVCPFeatureReply(t *testing.T) {
	type args struct {
		m     *ddcci.PhysicalMonitor
		coder vcp.Coder
	}
	tests := []struct {
		name         string
		args         args
		wantMaxValue int
		wantErr      bool
	}{
		{
			name: "Brightness",
			args: args{
				m:     ddcci.PhysicalMonitor0(),
				coder: vcp.Brightness,
			},
			wantMaxValue: 100,
			wantErr:      false,
		},
		{
			name: "Contrast",
			args: args{
				m:     ddcci.PhysicalMonitor0(),
				coder: vcp.Contrast,
			},
			wantMaxValue: 100,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.args.m

			gotCurrent, gotMaxValue, err := m.GetVCPFeatureAndVCPFeatureReply(tt.args.coder)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVCPFeatureAndVCPFeatureReply() error = %v, wantErr %v", err, tt.wantErr)
			}

			t.Logf("GetVCPFeatureAndVCPFeatureReply() gotCurrent = %v", gotCurrent)

			if gotMaxValue != tt.wantMaxValue {
				t.Errorf("GetVCPFeatureAndVCPFeatureReply() gotMaxValue = %v, want %v", gotMaxValue, tt.wantMaxValue)
			}
		})
	}
}

func TestPhysicalMonitor_SetBrightness(t *testing.T) {
	m := ddcci.PhysicalMonitor0()
	minValue, origin, maxValue, err := m.GetBrightness()
	assert.Nil(t, err)

	t.Logf("minValue = %v, origin = %v, maxValue = %v", minValue, origin, maxValue)
	defer t.Run(fmt.Sprintf("Recover Origin Brightness %d", origin), func(t *testing.T) {
		err := m.SetVCPFeature(vcp.Brightness.Value(origin))
		assert.Nil(t, err)
	})

	type args struct {
		f func(m *ddcci.PhysicalMonitor, value int) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Set Brightness With HighLevelApi",
			args: args{
				f: func(m *ddcci.PhysicalMonitor, value int) error {
					return m.SetBrightness(value)
				},
			},
			wantErr: false,
		},
		{
			name: "Set Brightness With SetVcpFeature",
			args: args{
				f: func(m *ddcci.PhysicalMonitor, value int) error {
					return m.SetVCPFeature(vcp.Brightness.Value(value))
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for value := minValue; value <= maxValue; value += (maxValue - minValue) / 10 {
				t.Run(strconv.Itoa(value), func(t *testing.T) {
					err := tt.args.f(m, value)
					assert.Nil(t, err)
				})

				time.Sleep(time.Second)
			}
		})
	}
}

func TestPhysicalMonitor_SetContrast(t *testing.T) {
	m := ddcci.PhysicalMonitor0()
	minValue, origin, maxValue, err := m.GetContrast()
	assert.Nil(t, err)

	t.Logf("minValue = %v, origin = %v, maxValue = %v", minValue, origin, maxValue)
	defer t.Run(fmt.Sprintf("Recover Origin Contrast %d", origin), func(t *testing.T) {
		err := m.SetVCPFeature(vcp.Contrast.Value(origin))
		assert.Nil(t, err)
	})

	type args struct {
		f func(m *ddcci.PhysicalMonitor, value int) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Set Contrast With HighLevelApi",
			args: args{
				f: func(m *ddcci.PhysicalMonitor, value int) error {
					return m.SetContrast(value)
				},
			},
			wantErr: false,
		},
		{
			name: "Set Contrast With SetVcpFeature",
			args: args{
				f: func(m *ddcci.PhysicalMonitor, value int) error {
					return m.SetVCPFeature(vcp.Contrast.Value(value))
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for value := minValue; value <= maxValue; value += (maxValue - minValue) / 10 {
				t.Run(strconv.Itoa(value), func(t *testing.T) {
					err := tt.args.f(m, value)
					assert.Nil(t, err)
				})

				time.Sleep(time.Second)
			}
		})
	}
}

func TestPhysicalMonitor_RestoreMonitorFactoryColorDefaults(t *testing.T) {
	t.Run("With HighLevel Api", func(t *testing.T) {
		m := ddcci.PhysicalMonitor0()
		err := m.RestoreFactoryColorDefaults()
		assert.Nil(t, err)
	})

	time.Sleep(time.Second)

	t.Run("With SetVCPFeature", func(t *testing.T) {
		m := ddcci.PhysicalMonitor0()
		err := m.SetVCPFeature(vcp.RestoreColorDefaults.Value(0))
		assert.Nil(t, err)
	})
}

func TestPhysicalMonitor_RestoreMonitorFactoryDefaults(t *testing.T) {
	t.Run("With HighLevel Api", func(t *testing.T) {
		m := ddcci.PhysicalMonitor0()
		err := m.RestoreFactoryDefaults()
		assert.Nil(t, err)
	})

	time.Sleep(time.Second)

	t.Run("With SetVCPFeature", func(t *testing.T) {
		m := ddcci.PhysicalMonitor0()
		err := m.SetVCPFeature(vcp.RestoreFactoryDefaults.Value(0))
		assert.Nil(t, err)
	})
}

func TestPhysicalMonitor_AudioSpeakerVolume(t *testing.T) {
	m := ddcci.PhysicalMonitor0()

	origin, maxValue, err := m.GetVCPFeatureAndVCPFeatureReply(vcp.AudioSpeakerVolume)
	assert.Nil(t, err)
	defer t.Run(fmt.Sprintf("Recover AudioSpeakerVolume %d", origin), func(t *testing.T) {
		err := m.SetVCPFeature(vcp.AudioSpeakerVolume.Value(origin))
		assert.Nil(t, err)
	})
	assert.Equal(t, 100, maxValue)

	for value := 0; value <= 100; value += 10 {
		t.Run(strconv.Itoa(value), func(t *testing.T) {
			err := m.SetVCPFeature(vcp.AudioSpeakerVolume.Value(origin))
			assert.Nil(t, err)

			current, maxValue, err := m.GetVCPFeatureAndVCPFeatureReply(vcp.AudioSpeakerVolume)
			assert.Nil(t, err)
			assert.Equal(t, value, current)
			assert.Equal(t, 100, maxValue)
		})
	}
}

func TestPhysicalMonitor_AudioMicrophoneVolume(t *testing.T) {
	m := ddcci.PhysicalMonitor0()

	origin, maxValue, err := m.GetVCPFeatureAndVCPFeatureReply(vcp.AudioMicrophoneVolume)
	assert.Nil(t, err)
	defer t.Run(fmt.Sprintf("Recover AudioMicrophoneVolume %d", origin), func(t *testing.T) {
		err := m.SetVCPFeature(vcp.AudioMicrophoneVolume.Value(origin))
		assert.Nil(t, err)
	})
	assert.Equal(t, 100, maxValue)

	for value := 0; value <= 100; value += 10 {
		t.Run(strconv.Itoa(value), func(t *testing.T) {
			err = m.SetVCPFeature(vcp.AudioMicrophoneVolume.Value(origin))
			assert.Nil(t, err)

			current, maxValue, err := m.GetVCPFeatureAndVCPFeatureReply(vcp.AudioMicrophoneVolume)
			assert.Nil(t, err)
			assert.Equal(t, value, current)
			assert.Equal(t, 100, maxValue)
		})
	}
}

func TestPhysicalMonitor_GetTimingReport(t *testing.T) {
	m := ddcci.PhysicalMonitor0()

	report, err := m.GetTimingReport()
	assert.Nil(t, err)
	assert.NotNil(t, report)

	t.Logf("report = %+v", report)

	// todo data exception

	{
		current, maxValue, err := m.GetVCPFeatureAndVCPFeatureReply(vcp.HorizontalFrequency)
		assert.Nil(t, err)

		t.Logf("current = %v, maxValue = %v", current, maxValue)
	}
	{
		current, maxValue, err := m.GetVCPFeatureAndVCPFeatureReply(vcp.VerticalFrequency)
		assert.Nil(t, err)

		t.Logf("current = %v, maxValue = %v", current, maxValue)
	}
}

func TestPhysicalMonitor_VCPVersion(t *testing.T) {
	//
	//
	// m  := ddcci.PhysicalMonitor0()
	//
	// version := ""
	//
	// err := m.GetVCPFeature(vcp.Version, &version)
	// assert.Nil(t, err)
	//
	// t.Logf("version = %+v", version)
}

func TestPhysicalMonitor_CapabilitiesRequestAndCapabilitiesReply(t *testing.T) {
	m := ddcci.PhysicalMonitor0()
	data, err := m.CapabilitiesRequestAndCapabilitiesReply()
	assert.Nil(t, err)
	assert.NotNil(t, data)

	t.Logf("capabilities = %s", data)
}

func TestPhysicalMonitor_DegaussMonitor(t *testing.T) {
	m := ddcci.PhysicalMonitor0()
	err := m.Degauss()
	assert.Nil(t, err)
}

func TestPhysicalMonitor_GetCapabilitiesStringLength(t *testing.T) {
	m := ddcci.PhysicalMonitor0()
	size, err := m.GetCapabilitiesStringLength()
	assert.Nil(t, err)

	t.Logf("size = %d", size)
}

func TestPhysicalMonitor_GetCapabilities(t *testing.T) {
	m := ddcci.PhysicalMonitor0()
	capabilities, temperatures, err := m.GetCapabilities()
	assert.Nil(t, err)

	t.Logf("capabilities = %v", capabilities)
	t.Logf("temperatures = %v", temperatures)
}

func TestPhysicalMonitor_GetColorTemperature(t *testing.T) {
	m := ddcci.PhysicalMonitor0()
	temperature, err := m.GetColorTemperature()
	assert.Nil(t, err)

	t.Logf("temperature = %v", temperature)
}

func TestPhysicalMonitor_GetDisplayAreaPosition(t *testing.T) {
	positionTypes := []ddcci.PositionType{ddcci.HorizontalPosition, ddcci.VerticalPosition}
	for _, positionType := range positionTypes {
		t.Run(fmt.Sprintf("positionType: %v", positionType), func(t *testing.T) {
			m := ddcci.PhysicalMonitor0()
			minValue, current, maxValue, err := m.GetDisplayAreaPosition(positionType)
			assert.Nil(t, err)

			t.Logf("minValue = %v, current = %v, maxValue = %v", minValue, current, maxValue)
		})
	}
}

func TestPhysicalMonitor_GetDisplayAreaSize(t *testing.T) {
	sizeTypes := []ddcci.SizeType{ddcci.Width, ddcci.Height}
	for _, sizeType := range sizeTypes {
		t.Run(fmt.Sprintf("sizeType: %v", sizeType), func(t *testing.T) {
			m := ddcci.PhysicalMonitor0()
			minValue, current, maxValue, err := m.GetDisplayAreaSize(sizeType)
			assert.Nil(t, err)

			t.Logf("minValue = %v, current = %v, maxValue = %v", minValue, current, maxValue)
		})
	}
}

func TestPhysicalMonitor_GetRedGreenOrBlueDrive(t *testing.T) {
	driveTypes := []ddcci.DriveType{ddcci.RedDrive, ddcci.GreenDrive, ddcci.BlueDrive}
	for _, driveType := range driveTypes {
		t.Run(fmt.Sprintf("driveType: %v", driveType), func(t *testing.T) {
			m := ddcci.PhysicalMonitor0()
			minValue, current, maxValue, err := m.GetRedGreenOrBlueDrive(driveType)
			assert.Nil(t, err)

			t.Logf("minValue = %v, current = %v, maxValue = %v", minValue, current, maxValue)
		})
	}
}

func TestPhysicalMonitor_GetRedGreenOrBlueGain(t *testing.T) {
	driveTypes := []ddcci.DriveType{ddcci.RedDrive, ddcci.GreenDrive, ddcci.BlueDrive}
	for _, driveType := range driveTypes {
		t.Run(fmt.Sprintf("driveType: %v", driveType), func(t *testing.T) {
			m := ddcci.PhysicalMonitor0()
			minValue, current, maxValue, err := m.GetRedGreenOrBlueGain(driveType)
			assert.Nil(t, err)

			t.Logf("minValue = %v, current = %v, maxValue = %v", minValue, current, maxValue)
		})
	}
}

func TestPhysicalMonitor_GetTechnologyType(t *testing.T) {
	m := ddcci.PhysicalMonitor0()
	value, err := m.GetTechnologyType()
	assert.Nil(t, err)

	t.Logf("value = %v", value)
}
