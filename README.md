# Temperature Manager for A10-4600M

A10-4600M with Radeon HD7660G APU does not have a good built-in throttling mechanism, combine it with poor cooling solution and you have a laptop that reboots frequently due to overheating.

This tool attempts to elimitate this behavior by periodically checking the APU temperature and limitting the CPU's max frequency and turbo boosting when it hits a threshold. The frequency and turbo are re-established when temperature falls back to safe range.

## Supporting software

* APU temperature is fetched using values exposed by [OpenHardwareMonitor](https://openhardwaremonitor.org/) via [WMI](https://docs.microsoft.com/en-us/windows/win32/wmisdk/wmi-start-page)
* APU max frequency and turbo boost are set to the current Window's powerprofile via [powercfg CLI](https://docs.microsoft.com/en-us/windows-hardware/design/device-experiences/powercfg-command-line-options)

## TO-DO
- [ ] Set Max stepping
- [ ] Disarm manager on error
- [ ] Rework manager logic to be completely agnostic to processor powerlimits, supporting processor profiles
- [ ] Unit Tests
- [ ] Introduce configuration for hardcoded values

## Build
To build without console window, use:
go build -ldflags "-H=windowsgui"