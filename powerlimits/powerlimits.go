package powerlimits

import (
	"fmt"
	"log"
	"strconv"
)

type CpuAndTurboLimits struct {
	Cpu   int
	Turbo int
}

func (lim *CpuAndTurboLimits) ToString() string {
	toString := fmt.Sprintf("%d%% - Turbo ", lim.Cpu)
	if lim.Turbo == 1 {
		return toString + "On"
	} else {
		return toString + "Off"
	}
}

func (lim *CpuAndTurboLimits) StepUp() bool {
	if lim.Cpu == 100 && lim.Turbo == 0 {
		lim.Turbo = 1
		return true
	} else if lim.Cpu < 100 {
		lim.Cpu += 10
		return true
	}
	return false
}

func (lim *CpuAndTurboLimits) StepDown() bool {
	if lim.Cpu == 100 && lim.Turbo == 1 {
		lim.Turbo = 0
		return true
	} else if lim.Cpu > 60 {
		lim.Cpu -= 10
		return true
	}
	return false
}

var lastLimits CpuAndTurboLimits

func SetPower(limits CpuAndTurboLimits) error {
	if limits == lastLimits {
		log.Println("Already applied.")
		return nil
	}

	schemeId, err := getCurrentScheme()
	if err != nil {
		return err
	}

	err = setCpuLimit(schemeId, strconv.Itoa(limits.Cpu))
	if err != nil {
		return err
	}

	err = setTurboMode(schemeId, strconv.Itoa(limits.Turbo))
	if err != nil {
		return err
	}

	err = applyScheme(schemeId)
	if err == nil {
		lastLimits = limits
	}

	return err
}
