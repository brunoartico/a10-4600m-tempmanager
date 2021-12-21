package manager

import (
	"log"
	"time"

	"bartico.com/a10-4600m-tempmanager/powerlimits"
	"bartico.com/a10-4600m-tempmanager/sensors"
)

var maxTemp float32 = 90
var minTemp float32 = 80
var currentLimit *powerlimits.CpuAndTurboLimits
var running = false
var sleepSeconds time.Duration = 2
var useSpacedStepUp = true
var skipNextStepUp = false

func init() {
	currentLimit = new(powerlimits.CpuAndTurboLimits)
	currentLimit.Cpu = 60
	currentLimit.Turbo = 0
	powerlimits.SetPower(*currentLimit)
}

func steppingManagerLoop() {
	for running == true {
		sensorReadings, err := sensors.ReadCPUTemp()
		if err != nil {
			log.Fatal(err)
		}

		if len(sensorReadings) != 1 {
			log.Fatal("Invalid sensor readings:", sensorReadings)
		}

		if cpuTemp := sensorReadings[0]; cpuTemp.Identifier == sensors.CPU_TEMP_SENSOR_ID {
			if cpuTemp.Value >= maxTemp && currentLimit.StepDown() {
				powerlimits.SetPower(*currentLimit)
			} else if cpuTemp.Value < minTemp && currentLimit.StepUp() {
				if useSpacedStepUp && skipNextStepUp {
					currentLimit.StepDown()
					skipNextStepUp = false
				} else {
					powerlimits.SetPower(*currentLimit)
					skipNextStepUp = true
				}
			}
		} else {
			log.Fatal("Sensor Reading is not cpu temp!")
		}

		log.Println(minTemp, maxTemp, currentLimit)
		time.Sleep(sleepSeconds * time.Second)
	}
}

func Start() {
	if !running {
		go steppingManagerLoop()
	}
	running = true
}

func Stop() {
	running = false
}

func StepUp() {
	currentLimit.StepUp()
	powerlimits.SetPower(*currentLimit)
}

func StepDown() {
	currentLimit.StepDown()
	powerlimits.SetPower(*currentLimit)
}

func SwitchTurboOff() {
	currentLimit.Turbo = 0
	powerlimits.SetPower(*currentLimit)
}

func SwitchTurboOn() {
	currentLimit.Turbo = 1
	powerlimits.SetPower(*currentLimit)
}

func GetLimits() *powerlimits.CpuAndTurboLimits {
	return currentLimit
}

func SetSleepDuration(sleep time.Duration) {
	sleepSeconds = sleep
}

func SetMaxTemp(temp float32) {
	maxTemp = temp
	minTemp = temp - 10
}
