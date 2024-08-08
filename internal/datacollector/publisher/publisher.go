package publisher

import (
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

func NewPublisher(config datacollector.Config) Publisher {

	switch config.Config["PUBLISHER_CONNECTOR"] {
	case "mqtt":
		u, _ := url.Parse(config.Config["PUBLISHER_MQTT_BROKER"])
		clientConfig := mqtt.ClientConfig{
			Ctx:      config.Context,
			Brokers:  []*url.URL{u},
			Username: config.Config["PUBLISHER_MQTT_USERNAME"],
			Password: config.Config["PUBLISHER_MQTT_PASSWORD"],
		}
		return NewIotPublisher(config.Config["IOT_SPACE_ID"], config.Config["IOT_THING_ID"], mqtt.NewMqttConnector(clientConfig))
	default:
		return nil
	}
}
