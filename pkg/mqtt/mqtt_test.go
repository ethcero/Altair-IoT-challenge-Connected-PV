package mqtt

import (
	"context"
	"crypto/tls"
	"net/url"
	"testing"

	"github.com/eclipse/paho.golang/paho"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockConnectionManager is a mock implementation of the ConnectionManager interface
type MockConnectionManager struct {
	mock.Mock
}

func (m *MockConnectionManager) AwaitConnection(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockConnectionManager) Publish(ctx context.Context, pub *paho.Publish) (*paho.PublishResponse, error) {
	args := m.Called(ctx, pub)
	return args.Get(0).(*paho.PublishResponse), args.Error(1)
}

func TestNewMqttConnector(t *testing.T) {
	brokerURL, _ := url.Parse("tcp://localhost:1883")
	clientConfig := ClientConfig{
		Ctx:       context.Background(),
		ClientId:  "testClient",
		Brokers:   []*url.URL{brokerURL},
		TlsConfig: &tls.Config{},
		Topics:    []string{"test/topic"},
		Username:  "user",
		Password:  "pass",
		Handler:   nil,
	}

	client := NewMqttConnector(clientConfig)

	assert.NotNil(t, client)
	assert.Equal(t, clientConfig.Ctx, client.Context)
	assert.Equal(t, "user", client.Username)
	assert.Equal(t, []byte("pass"), client.Password)
}

func TestClient_Connect(t *testing.T) {
	brokerURL, _ := url.Parse("tcp://localhost:1883")
	clientConfig := ClientConfig{
		Ctx:       context.Background(),
		ClientId:  "testClient",
		Brokers:   []*url.URL{brokerURL},
		TlsConfig: &tls.Config{},
		Topics:    []string{"test/topic"},
		Username:  "user",
		Password:  "pass",
		Handler:   nil,
	}

	client := NewMqttConnector(clientConfig)
	mockCM := new(MockConnectionManager)
	client.cm = mockCM

	mockCM.On("AwaitConnection", mock.Anything).Return(nil)
	mockCM.On("Done").Return(make(chan struct{}))

	err := client.Connect()
	assert.NoError(t, err)
}

func TestClient_Publish(t *testing.T) {
	client := &Client{
		Context: context.Background(),
		cm:      new(MockConnectionManager),
	}

	mockCM := client.cm.(*MockConnectionManager)
	mockCM.On("AwaitConnection", mock.Anything).Return(nil)
	mockCM.On("Publish", mock.Anything, mock.Anything).Return(&paho.PublishResponse{ReasonCode: 0}, nil)

	err := client.Publish("test/topic", []byte("test message"))
	assert.NoError(t, err)
}

func TestClient_AwaitConnection(t *testing.T) {
	client := &Client{
		Context: context.Background(),
		cm:      new(MockConnectionManager),
	}

	mockCM := client.cm.(*MockConnectionManager)
	mockCM.On("AwaitConnection", mock.Anything).Return(nil)

	err := client.AwaitConnection()
	assert.NoError(t, err)
}
