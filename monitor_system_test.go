package ddcci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemMonitor_GetNumberOfPhysicalMonitorsFromHMONITOR(t *testing.T) {
	a := assert.New(t)

	monitors, err := NewSystemMonitors()
	a.Nil(err)

	for i, monitor := range monitors {
		num, err := monitor.GetNumberOfPhysicalMonitorsFromHMONITOR()
		a.Nil(err)

		a.Equal(i+1, num)
	}
}
