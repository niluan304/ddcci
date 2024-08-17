## DDCCI

`ddcci` is a go library that calls the interface provided by the operating system and interacts with the display through the `DDC/CI` display communication standard protocol.
- Get the description, driver information, and current location of the monitor
- Get/Set the brightness, contrast, color temperature, or other adjustable parameters of the monitor
- Modify the power mode of the monitor (requires support from the monitor itself)

## Install
```shell
go get github.com/niluan304/ddcci
```

## Example
```go
package main

import (
	"fmt"

	"github.com/niluan304/ddcci"
)

func main() {
	monitors, err := ddcci.NewSystemMonitors()
	if err != nil {
		panic(err)
	}

	for _, monitor := range monitors {
		m, err := ddcci.NewPhysicalMonitor(&monitor)
		if err != nil {
			panic(err)
		}

		minValue, current, maxValue, err := m.GetBrightness()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Brightness: minValue %d, current: %d, maxValue %d\n", minValue, current, maxValue)
	}
}

```

## FAQ
### What operating systems does DDCCI support?
- Now only Windows. Welcome PR to support other OS.

## Reference
- [ddcutil Documentation](https://www.ddcutil.com/vcpinfo_output/)
- [Monitor Configuration | Microsoft Learn](https://learn.microsoft.com/en-us/windows/win32/api/_monitor/)

## Credits
- [goland](https://www.jetbrains.com/go/)