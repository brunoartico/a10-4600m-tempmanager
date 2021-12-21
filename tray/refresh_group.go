package tray

import (
	"bartico.com/tempmgr/manager"
	"github.com/getlantern/systray"
)

func createRefreshGroup() *systray.MenuItem {
	var rootItem = systray.AddMenuItem("Refresh Rate", "Refresh Rate")
	group := new(MenuGroup)

	group.Add(rootItem.AddSubMenuItemCheckbox("5 seconds refresh", "5 seconds refresh", false), func() {
		manager.SetSleepDuration(5)
	})

	group.Add(rootItem.AddSubMenuItemCheckbox("2 seconds refresh", "2 seconds refresh", true), func() {
		manager.SetSleepDuration(2)
	})

	group.Add(rootItem.AddSubMenuItemCheckbox("1 seconds refresh", "1 seconds refresh", false), func() {
		manager.SetSleepDuration(1)
	})

	return rootItem
}
