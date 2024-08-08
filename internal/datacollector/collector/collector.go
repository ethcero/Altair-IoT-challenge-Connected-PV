package collector

import (
	"fmt"
	"github.com/ethcero/connected-pv/internal/datacollector"
)

type Collector interface {
	gatherPowerData() (datacollector.PowerData, error)
	gatherDeviceData() (datacollector.DeviceData, error)
}

func CollectAndDispatch(c Collector, bus chan datacollector.BusMessage) {
	powerData, err := c.gatherPowerData()
	if err != nil {
		println("Error gathering power data")
	}

	deviceData, err := c.gatherDeviceData()
	if err != nil {
		println("Error gathering device data")
	}

	fmt.Println("Data collected")
	fmt.Println(powerData)
	fmt.Println(deviceData)
	bus <- datacollector.BusMessage{
		PowerData:  powerData,
		DeviceData: deviceData,
	}

}

func NewCollector(model string, address string) Collector {
	switch model {
	case "fronius":
		return NewFroniusInverter(address)
	default:
		return nil
	}
}
