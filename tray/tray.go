package tray

import (
	"time"

	"bartico.com/tempmgr/icon"
	"bartico.com/tempmgr/manager"
	"github.com/getlantern/systray"
)

var refreshRateGroup *systray.MenuItem
var maxTempGroup *systray.MenuItem
var mStopManager *systray.MenuItem
var mStartManager *systray.MenuItem
var mStepUp *systray.MenuItem
var mStepDown *systray.MenuItem
var mLimits *systray.MenuItem
var mTurbo *systray.MenuItem

func addGroupControls() {
	maxTempGroup = createTempGroup()
	refreshRateGroup = createRefreshGroup()
}

func addManagerControls() {
	mStopManager = systray.AddMenuItem("Stop Manager", "Stop Manager")
	mStartManager = systray.AddMenuItem("Start Manager", "Start Manager")

	go func() {
		for {
			<-mStopManager.ClickedCh
			maxTempGroup.Disable()
			refreshRateGroup.Disable()
			mStepDown.Enable()
			mStepUp.Enable()
			mTurbo.Enable()
			mStopManager.Hide()
			mStartManager.Show()
			manager.Stop()
		}
	}()

	go func() {
		for {
			<-mStartManager.ClickedCh
			maxTempGroup.Enable()
			refreshRateGroup.Enable()
			mStepDown.Disable()
			mStepUp.Disable()
			mTurbo.Disable()
			mStopManager.Show()
			mStartManager.Hide()
			manager.Start()
		}
	}()

	maxTempGroup.Disable()
	refreshRateGroup.Disable()
	mStopManager.Hide()
}

func addManualControls() {
	mTurbo = systray.AddMenuItemCheckbox("Turbo Mode", "Switch Turbo Mode", false)

	go func() {
		for {
			select {
				case <-mTurbo.ClickedCh:
					if (mTurbo.Checked()) {
						manager.SwitchTurboOff()
						mTurbo.Uncheck()
					} else {
						manager.SwitchTurboOn()
						mTurbo.Check()
					}
				case <-time.After(3 * time.Second):
					if (manager.GetLimits().Turbo == 1) {
						mTurbo.Check()
					} else {
						mTurbo.Uncheck()
					}
			}
		}
	}()

	mStepUp = systray.AddMenuItem("Step Up", "Step Up")

	go func() {
		for {
			<-mStepUp.ClickedCh
			manager.StepUp()
		}
	}()

	mStepDown = systray.AddMenuItem("Step Down", "Step Down")

	go func() {
		for {
			<-mStepDown.ClickedCh
			manager.StepDown()
		}
	}()

	mTurbo.Enable()
	mStepDown.Enable()
	mStepUp.Enable()
}

func addLimitInfo() {
	mLimits = systray.AddMenuItem("Unknown Limits", "Current Stepping")
	mLimits.Disable()
	go func() {
		for {
			currentLimits := manager.GetLimits()
			if currentLimits != nil {
				mLimits.SetTitle(manager.GetLimits().ToString())
			}
			time.Sleep(time.Second)
		}
	}()
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Temperature Manager")

	addLimitInfo()
	systray.AddSeparator()
	addGroupControls()
	addManagerControls()
	systray.AddSeparator()
	addManualControls()
	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Quit")
	go func() {
		<-mQuit.ClickedCh
		manager.Stop()
		systray.Quit()
	}()
}

func Start() {
	systray.Run(onReady, nil)
}
