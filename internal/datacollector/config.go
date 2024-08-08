package datacollector

import (
	"context"
	"os"
)

type Config struct {
	Context         context.Context
	IotConfig       IoTconfig
	CollectorConfig CollectorConfig
	PublisherConfig PublisherConfig
}

type IoTconfig struct {
	SpaceID string
	ThingID string
}

type CollectorConfig struct {
	Model   string
	Address string
}

type PublisherConfig struct {
	Connector           string
	MqttConnectorConfig MqttConnectorConfig
}

type MqttConnectorConfig struct {
	Broker   string
	Username string
	Password string
}

const (
	CollectorModelFronius  = "fronius"
	PublisherConnectorMQTT = "mqtt"
)

func NewConfig() Config {
	c := Config{
		Context: context.Background(),
		IotConfig: IoTconfig{
			SpaceID: os.Getenv("IOT_SPACE_ID"),
			ThingID: os.Getenv("IOT_THING_ID"),
		},
		CollectorConfig: CollectorConfig{
			Model:   os.Getenv("COLLECTOR_MODEL"),
			Address: os.Getenv("COLLECTOR_ADDRESS"),
		},
		PublisherConfig: PublisherConfig{
			Connector: os.Getenv("PUBLISHER_CONNECTOR"),
			MqttConnectorConfig: MqttConnectorConfig{
				Broker:   os.Getenv("PUBLISHER_MQTT_BROKER"),
				Username: os.Getenv("PUBLISHER_MQTT_USERNAME"),
				Password: os.Getenv("PUBLISHER_MQTT_PASSWORD"),
			},
		},
	}
	checkConfig(c)
	return c
}

func checkConfig(config Config) {
	if config.IotConfig.SpaceID == "" {
		panic("IOT_SPACE_ID is required")
	}
	if config.IotConfig.ThingID == "" {
		panic("IOT_THING_ID is required")
	}
	if config.CollectorConfig.Model == "" {
		panic("COLLECTOR_MODEL is required")
	}
	if config.CollectorConfig.Address == "" {
		panic("COLLECTOR_ADDRESS is required")
	}
	if config.PublisherConfig.Connector == "" {
		panic("PUBLISHER_CONNECTOR is required")
	}
	switch config.PublisherConfig.Connector {
	case PublisherConnectorMQTT:
		if config.PublisherConfig.MqttConnectorConfig.Broker == "" {
			panic("PUBLISHER_MQTT_BROKER is required")
		}
		if config.PublisherConfig.MqttConnectorConfig.Username == "" {
			panic("PUBLISHER_MQTT_USERNAME is required")
		}
		if config.PublisherConfig.MqttConnectorConfig.Password == "" {
			panic("PUBLISHER_MQTT_PASSWORD is required")
		}
	}

}
