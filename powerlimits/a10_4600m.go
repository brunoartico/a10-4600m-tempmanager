package powerlimits

const cpu_group = "54533251-82be-4824-96c1-47b60b740d00"
const cpu_limit_subgroup = "bc5038f7-23e0-4960-96da-33abaf5935ec"
const gpu_group = "f693fb01-e858-4f00-b20f-f30e12ac06d6"
const gpu_performance_subgroup = "191f65b5-d45c-4a4f-8aae-1ab8bfd980e6"

func getCurrentScheme() (string, error) {
	out, err := outputWithoutWindow("powercfg", "/GETACTIVESCHEME")
	if err != nil {
		return "", err
	}

	return string(out[19:55]), nil
}

func setCpuLimit(schemeId, cpuLimit string) error {
	return runWithoutWindow("powercfg", "/SETACVALUEINDEX", schemeId, cpu_group, cpu_limit_subgroup, cpuLimit)
}

func setTurboMode(schemeId, turboMode string) error {
	return runWithoutWindow("powercfg", "/SETACVALUEINDEX", schemeId, gpu_group, gpu_performance_subgroup, turboMode)
}

func applyScheme(schemeId string) error {
	return runWithoutWindow("powercfg", "/S", schemeId)
}
