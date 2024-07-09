
## Interesting data to collect

- `/solar_api/v1/GetPowerFlowRealtimeData.fcgi`: Solar power flow data
  ```python
    # mandatory field
    #this value is null if no meter is enabled ( + from grid , - to grid )
    number P_Grid ;
    # mandatory field
    #this value is null if no meter is enabled ( + generator , - consumer )
    number P_Load ;
    # mandatory field
    #this value is null if inverter is not running ( + production ( default ) )
    # On GEN24 and SymoHybrid : reports production on DC side (PV generator ).
    # On SnapInverter : is ident to power generated on AC side (ac power output ).
    number P_PV;

    # mandatory field
    # available since Fronius Hybrid version 1.3.1 -1
    # available since Fronius Non Hybrid version 3.7.1 -2
    # current relative self consumption in %, null if no smart meter is connected
    number rel_SelfConsumption ;
    # mandatory field
    # available since Fronius Hybrid version 1.3.1 -1
    # available since Fronius Non Hybrid version 3.7.1 -2
    # current relative autonomy in %, null if no smart meter is connected
    number rel_Autonomy ;
    # optional field
    # implemented since Fronius Non Hybrid version 3.4.1 -7
    # this value is always null on GEN24 /Tauro
    # AC Energy [Wh] this day , null if no inverter is connected
    number E_Day;
    # optional field
    # implemented since Fronius Non Hybrid version 3.4.1 -7
    # this value is always null on GEN24 /Tauro
    # AC Energy [Wh] this year , null if no inverter is connected
    number E_Year ;
    # optional field
    # implemented since Fronius Non Hybrid version 3.4.1 -7
    # implemented since Fronius GEN24 /Tauro version 1.14 and null before
    # updated only every 5 minutes on GEN24/ Tauro .
    # AC Energy [Wh] ever since , null if no inverter is connected
    number E_Total ;
    ```
    example:
    ```json
    {
    "Body" : {
        "Data" : {
            "Inverters" : {
                "1" : {
                "DT" : 78,
                "E_Day" : 5514,
                "E_Total" : 19207280,
                "E_Year" : 3559556.75,
                "P" : 2492
                }
            },
            "Site" : {
                "E_Day" : 5514,
                "E_Total" : 19207280,
                "E_Year" : 3559556.75,
                "Meter_Location" : "grid",
                "Mode" : "meter",
                "P_Akku" : null,
                "P_Grid" : -1025.2,
                "P_Load" : -1466.8,
                "P_PV" : 2492,
                "rel_Autonomy" : 100,
                "rel_SelfConsumption" : 58.860353130016051
            },
            "Version" : "12"
        }
    },
    "Head" : {
        "RequestArguments" : {},
        "Status" : {
            "Code" : 0,
            "Reason" : "",
            "UserMessage" : ""
        },
        "Timestamp" : "2024-07-09T12:09:19+02:00"
    }
    }
    ```


- `/solar_api/v1/GetInverterRealtimeData.cgi?Scope=Device&DeviceId=1&DataCollection=CommonInverterData`: Inverter realtime data
  ```json
    {
    "Body" : {
        "Data" : {
            "DAY_ENERGY" : {
                "Unit" : "Wh",
                "Value" : 5664
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
                "Value" : 49.979999999999997
            },
            "IAC" : {
                "Unit" : "A",
                "Value" : 9.7799999999999994
            },
            "IDC" : {
                "Unit" : "A",
                "Value" : 19.800000000000001
            },
            "PAC" : {
                "Unit" : "W",
                "Value" : 2534
            },
            "TOTAL_ENERGY" : {
                "Unit" : "Wh",
                "Value" : 19207420
            },
            "UAC" : {
                "Unit" : "V",
                "Value" : 252.30000000000001
            },
            "UDC" : {
                "Unit" : "V",
                "Value" : 134.69999999999999
            },
            "YEAR_ENERGY" : {
                "Unit" : "Wh",
                "Value" : 3559706
            }
        }
    },
    "Head" : {
        "RequestArguments" : {
            "DataCollection" : "CommonInverterData",
            "DeviceClass" : "Inverter",
            "DeviceId" : "1",
            "Scope" : "Device"
        },
        "Status" : {
            "Code" : 0,
            "Reason" : "",
            "UserMessage" : ""
        },
        "Timestamp" : "2024-07-09T12:12:52+02:00"
    }
    }
  ```