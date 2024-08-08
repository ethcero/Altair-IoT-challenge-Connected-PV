package mqtt

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"log"
	"net/url"
	"time"
)

type ConnectionManager interface {
	AwaitConnection(ctx context.Context) error
	Publish(ctx context.Context, p *paho.Publish) (*paho.PublishResponse, error)
}

type Client struct {
	clientConfig autopaho.ClientConfig
	Context      context.Context
	cm           ConnectionManager
	Username     string
	Password     []byte
}

type ClientConfig struct {
	Ctx       context.Context
	ClientId  string
	Brokers   []*url.URL
	TlsConfig *tls.Config
	Topics    []string
	Username  string
	Password  string
	Handler   func(m *paho.Publish)
}

func NewMqttConnector(clientConfig ClientConfig) *Client {
	pahoClientConfig := autopaho.ClientConfig{
		BrokerUrls:        clientConfig.Brokers,
		TlsCfg:            clientConfig.TlsConfig,
		KeepAlive:         120,
		ConnectRetryDelay: 4 * time.Second,
		ConnectTimeout:    10 * time.Second,
		WebSocketCfg:      nil,
		OnConnectionUp:    nil,
		OnConnectError:    newOnErrorCallback(clientConfig.ClientId),
		Debug:             nil,
		PahoDebug:         nil,
		ClientConfig:      newClientConfig(clientConfig.ClientId),
	}

	return &Client{
		pahoClientConfig,
		clientConfig.Ctx,
		nil,
		clientConfig.Username,
		[]byte(clientConfig.Password),
	}

}

func newClientConfig(clientId string) paho.ClientConfig {
	return paho.ClientConfig{
		ClientID: clientId,
		OnServerDisconnect: func(d *paho.Disconnect) {
			log.Println("OnServerDisconnect: " + fmt.Sprintf("%+v", d))
		},
		OnClientError: func(err error) {
			log.Fatal("OnClientError: " + err.Error())
		},
	}
}

func (c *Client) Publish(topic string, data []byte) error {

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if errConnection := c.cm.AwaitConnection(ctxWithTimeout); errConnection != nil {
		log.Println("publish await conn error: " + errConnection.Error())
		return errConnection
	}

	pubResponse, errPub := c.cm.Publish(c.Context, &paho.Publish{
		QoS:     1,
		Topic:   topic,
		Payload: data,
	})

	if errPub != nil {
		log.Println("mqtt publish error: " + errPub.Error())
		return errPub
	}

	if pubResponse != nil {
		if pubResponse.ReasonCode != 0 && pubResponse.ReasonCode != 16 {
			// MQTT Protocol codes: 0 = Success 16 = Server received message but there are no subscribers
			msgError := fmt.Sprintf("pub reason-code: %d", pubResponse.ReasonCode)
			return errors.New(msgError)
		}
	}

	return nil
}
func (c *Client) Connect() error {
	var err error

	c.clientConfig.ConnectUsername = c.Username
	c.clientConfig.ConnectPassword = c.Password
	if c.cm == nil {
		c.cm, err = autopaho.NewConnection(c.Context, c.clientConfig)
		if err != nil {
			log.Println("New Connection MQTT error: " + err.Error())
			return err
		}
	}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	errAwait := c.cm.AwaitConnection(ctxWithTimeout)
	if errAwait != nil {
		log.Println("AwaitConnection error: " + errAwait.Error())
		return errAwait
	}

	log.Println("Connect MQTT OK: " + c.Username)

	return nil
}

func (c *Client) AwaitConnection() error {
	return c.cm.AwaitConnection(c.Context)
}

func newOnErrorCallback(clientId string) func(err error) {
	return func(err error) {
		errorMessage := "NewOnErrorCallback '" + clientId + "' error: " + err.Error()
		log.Println(errorMessage)
	}
}
