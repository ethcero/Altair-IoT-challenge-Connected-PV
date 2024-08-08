package collector

import (
	"github.com/ethcero/connected-pv/internal/datacollector"
	"log"
)

type Collector interface {
	gatherPowerData() (datacollector.PowerData, error)
	gatherDeviceData() (datacollector.DeviceData, error)
}

func CollectAndDispatch(c Collector, bus chan datacollector.BusMessage) {
	powerData, err := c.gatherPowerData()
	if err != nil {
		log.Println("Error gathering power data")
	}

	deviceData, err := c.gatherDeviceData()
	if err != nil {
		log.Println("Error gathering device data")
	}

	bus <- datacollector.BusMessage{
		PowerData:  powerData,
		DeviceData: deviceData,
	}

}

func NewCollector(config datacollector.CollectorConfig) Collector {
	switch config.Model {
	case datacollector.CollectorModelFronius:
		return NewFroniusInverter(config.Address)
	default:
		return nil
	}
}
