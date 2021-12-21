package sensors

import (
	"github.com/StackExchange/wmi"
)

type WmiSensor struct {
	Identifier string
	Value      float32
}

const CPU_TEMP_SENSOR_ID = "/amdcpu/0/temperature/0"
const GPU_LOAD_SENSOR_ID = "/atigpu/0/load/0"

func ReadAll() ([]WmiSensor, error) {
	var dst []WmiSensor
	err := wmi.QueryNamespace("SELECT Identifier, Value FROM Sensor WHERE Identifier='"+CPU_TEMP_SENSOR_ID+"' OR Identifier='"+GPU_LOAD_SENSOR_ID+"'", &dst, "ROOT\\OpenHardwareMonitor")
	return dst, err
}

func ReadCPUTemp() ([]WmiSensor, error) {
	var dst []WmiSensor
	err := wmi.QueryNamespace("SELECT Identifier, Value FROM Sensor WHERE Identifier='"+CPU_TEMP_SENSOR_ID+"'", &dst, "ROOT\\OpenHardwareMonitor")
	return dst, err
}
