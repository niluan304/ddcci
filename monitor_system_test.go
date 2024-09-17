//go:build windows

package ddcci_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/niluan304/ddcci"
)

func TestSystemMonitor_GetNumberOfPhysicalMonitorsFromHMONITOR(t *testing.T) {
	monitors, err := ddcci.NewSystemMonitors()
	require.NoError(t, err)

	for i, monitor := range monitors {
		num, err := monitor.GetNumberOfPhysicalMonitorsFromHMONITOR()
		require.NoError(t, err)

		assert.Equal(t, i+1, num)
	}
}
