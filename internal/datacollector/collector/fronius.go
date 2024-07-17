package collector

import (
	"encoding/json"
	"github.com/ethcero/connected-pv/pkg/api"
)

type FroniusCollector struct {
	address string
}

func NewFroniusInverter(address string) *FroniusCollector {
	return &FroniusCollector{
		address: address,
	}
}

func (i *FroniusCollector) gatherPowerData() (PowerData, error) {

	endpoint := i.address + "/solar_api/v1/GetPowerFlowRealtimeData.fcgi"

	resp, err := api.Get(api.Request{
		Url: endpoint,
	})
	if err != nil {
		return PowerData{}, err
	}

	var froniusPowerData FroniusPowerData
	err = json.Unmarshal(resp.Body, &froniusPowerData)
	if err != nil {
		return PowerData{}, err
	}

	powerData := PowerData{
		PGrid: froniusPowerData.Body.Data.Site.PGrid,
		PLoad: froniusPowerData.Body.Data.Site.PLoad,
		PPV:   froniusPowerData.Body.Data.Site.PPV,
	}
	return powerData, nil
}

func (i *FroniusCollector) gatherDeviceData() (DeviceData, error) {
	endpoint := i.address + "/solar_api/v1/GetInverterRealtimeData.cgi?Scope=Device&DeviceId=1&DataCollection=CommonInverterData"

	resp, err := api.Get(api.Request{
		Url: endpoint,
	})
	if err != nil {
		return DeviceData{}, err
	}

	var froniusDeviceData FroniusDeviceData
	err = json.Unmarshal(resp.Body, &froniusDeviceData)
	if err != nil {
		return DeviceData{}, err
	}

	deviceData := DeviceData{
		status: DeviceDataStatus{
			ErrorCode:  froniusDeviceData.Body.Data.DeviceStatus.ErrorCode,
			StatusCode: froniusDeviceData.Body.Data.DeviceStatus.StatusCode,
		},
	}
	return deviceData, nil
}

type FroniusPowerData struct {
	Body struct {
		Data struct {
			Inverters struct {
				One struct {
					DT     int
					P      int
					EDay   float64 `json:"E_Day"`
					ETotal float64 `json:"E_Total"`
					EYear  float64 `json:"E_Year"`
				}
			}
			Site struct {
				EDay               float64     `json:"E_Day"`
				ETotal             float64     `json:"E_Total"`
				EYear              float64     `json:"E_Year"`
				MeterLocation      string      `json:"Meter_Location"`
				Mode               string      `json:"Mode"`
				PAkku              interface{} `json:"P_Akku"`
				PGrid              float64     `json:"P_Grid"`
				PLoad              float64     `json:"P_Load"`
				PPV                float64     `json:"P_PV"`
				RelAutonomy        float64     `json:"rel_Autonomy"`
				RelSelfConsumption float64     `json:"rel_SelfConsumption"`
			}
		}
	}
}

type FroniusDeviceData struct {
	Body struct {
		Data struct {
			DeviceStatus struct {
				ErrorCode              int
				LEDColor               int
				LEDState               int
				MgmtTimerRemainingTime int
				StateToReset           bool
				StatusCode             int
			} `json:"DeviceStatus"`
		}
	}
}

/*"Body" : {
"Data" : {
"DAY_ENERGY" : {
"Unit" : "Wh",
"Value" : 21534
},
"DeviceStatus" : {
"ErrorCode" : 0,
"LEDColor" : 2,
"LEDState" : 0,
"MgmtTimerRemainingTime" : -1,
"StateToReset" : false,
"StatusCode" : 7
},
"FAC" : {
"Unit" : "Hz",
"Value" : 49.960000000000001
},
"IAC" : {
"Unit" : "A",
"Value" : 0.56999999999999995
},
"IDC" : {
"Unit" : "A",
"Value" : 1.1799999999999999
},
"PAC" : {
"Unit" : "W",
"Value" : 121
},
"TOTAL_ENERGY" : {
"Unit" : "Wh",
"Value" : 19405940
},
"UAC" : {
"Unit" : "V",
"Value" : 239.30000000000001
},
"UDC" : {
"Unit" : "V",
"Value" : 125.09999999999999
},
"YEAR_ENERGY" : {
"Unit" : "Wh",
"Value" : 3758216.5
}
}
},
"Head" : {
"RequestArguments" : {
"DataCollection" : "CommonInverterData",
"DeviceClass" : "Inverter",
"DeviceId" : "1",
"Scope" : "Device"
},*/
