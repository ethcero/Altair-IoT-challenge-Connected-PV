package collector

import "fmt"

type PowerData struct {
	PGrid float64
	PLoad float64
	PPV   float64
}

type DeviceDataStatus struct {
	ErrorCode  int
	StatusCode int
}
type DeviceData struct {
	status DeviceDataStatus
}

type Collector interface {
	gatherPowerData() (PowerData, error)
	gatherDeviceData() (DeviceData, error)
}

func CollectAndDispatch(c Collector) {
	powerData, err := c.gatherPowerData()
	if err != nil {
		println("Error gathering power data")
	}

	deviceData, err := c.gatherDeviceData()
	if err != nil {
		println("Error gathering device data")
	}

	fmt.Println(powerData)
	fmt.Println(deviceData)
}

func NewCollector(model string, address string) Collector {
	switch model {
	case "fronius":
		return NewFroniusInverter(address)
	default:
		return nil
	}
}
