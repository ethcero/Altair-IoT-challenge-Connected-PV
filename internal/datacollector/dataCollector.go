package datacollector

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
	Status DeviceDataStatus
}

type BusMessage struct {
	PowerData  PowerData
	DeviceData DeviceData
}
