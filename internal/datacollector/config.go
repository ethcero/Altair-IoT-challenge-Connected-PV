package datacollector

import (
	"context"
	"os"
	"strconv"
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
	Model    string
	Address  string
	Interval int
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
			Model:    os.Getenv("COLLECTOR_MODEL"),
			Address:  os.Getenv("COLLECTOR_ADDRESS"),
			Interval: castToInt(getEnvOrDefault("COLLECTOR_INTERVAL", "5")),
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

func castToInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic("Error casting to int: " + value)
	}
	return i
}

func getEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
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
