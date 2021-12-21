package tray

import (
	"bartico.com/a10-4600m-tempmanager/manager"
	"github.com/getlantern/systray"
)

func createTempGroup() *systray.MenuItem {
	var rootItem = systray.AddMenuItem("Max Temperature", "Max Temperature")

	group := new(MenuGroup)
	group.Add(rootItem.AddSubMenuItemCheckbox("95°C Max", "95°C Max", false), func() {
		manager.SetMaxTemp(95)
	})

	group.Add(rootItem.AddSubMenuItemCheckbox("90°C Max", "90°C Max", true), func() {
		manager.SetMaxTemp(90)
	})

	group.Add(rootItem.AddSubMenuItemCheckbox("85°C Max", "85°C Max", false), func() {
		manager.SetMaxTemp(85)
	})

	group.Add(rootItem.AddSubMenuItemCheckbox("80°C Max", "80°C Max", false), func() {
		manager.SetMaxTemp(80)
	})

	group.Add(rootItem.AddSubMenuItemCheckbox("75°C Max", "75°C Max", false), func() {
		manager.SetMaxTemp(75)
	})

	return rootItem
}
