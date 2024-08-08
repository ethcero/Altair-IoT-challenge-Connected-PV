package publisher

import (
	"context"
	"github.com/ethcero/connected-pv/internal/datacollector"
	"github.com/ethcero/connected-pv/pkg/mqtt"
	"log"
	"net/url"
)

type Publisher interface {
	Start()
	Publish(data datacollector.BusMessage) error
}

func HandlePublish(p Publisher, bus chan datacollector.BusMessage) {

	go func() {
		for {
			select {
			case data := <-bus:
				err := p.Publish(data)
				if err != nil {
					log.Printf("Error publishing data: %s\n", err)
				}
			}
		}
	}()
}

func NewPublisher(ctx context.Context, publisherConfig datacollector.PublisherConfig, iotConfig datacollector.IoTconfig) Publisher {

	switch publisherConfig.Connector {
	case datacollector.PublisherConnectorMQTT:
		mqttConfig := publisherConfig.MqttConnectorConfig
		u, _ := url.Parse(mqttConfig.Broker)
		clientConfig := mqtt.ClientConfig{
			Ctx:      ctx,
			Brokers:  []*url.URL{u},
			Username: mqttConfig.Username,
			Password: mqttConfig.Password,
		}
		return NewIotPublisher(iotConfig.SpaceID, iotConfig.ThingID, mqtt.NewMqttConnector(clientConfig))
	default:
		return nil
	}
}
