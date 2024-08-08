package publisher

import (
	"encoding/json"
	"github.com/ethcero/connected-pv/internal/datacollector"
	"github.com/ethcero/connected-pv/pkg/mqtt"
	"log"
)

type IotPublisher struct {
	SpaceId string
	ThingId string
	Mqtt    *mqtt.Client
}

func NewIotPublisher(spaceId string, thingId string, mqtt *mqtt.Client) *IotPublisher {
	return &IotPublisher{
		SpaceId: spaceId,
		ThingId: thingId,
		Mqtt:    mqtt,
	}
}

func (p *IotPublisher) Start() {
	log.Println("Starting publisher IoT")
	go func() {
		err := p.Mqtt.Connect()
		if err != nil {
			log.Fatalf("Error connecting to MQTT broker: %s", err)
		}
	}()
}

func (p *IotPublisher) Publish(data datacollector.BusMessage) error {
	topic := "spaces/" + p.SpaceId + "/things/" + p.ThingId + "/properties"

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	log.Printf("Publishing data to topic: %s\n", topic)
	log.Printf("Data: %s\n", string(dataBytes))
	return p.Mqtt.Publish(topic, dataBytes)
}
