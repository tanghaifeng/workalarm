package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	ui.Main(func() {
		drinkWater := ui.NewLabel(`1: 要么喝水`)

		toilet := ui.NewLabel(`2: 要么上厕所`)

		run := ui.NewLabel(`3: 要么动起来`)

		// 生成：垂直容器
		box := ui.NewHorizontalBox()

		box.SetPadded(true)

		box.Append(drinkWater, true)
		box.Append(toilet, true)
		box.Append(run, true)

		window := ui.NewWindow(`身体是革命的本钱，你自己悠着点`, 600, 200, true)
		window.SetChild(box)

		window.SetMargined(true)

		window.OnClosing(func(*ui.Window) bool {
			// 窗体关闭
			ui.Quit()
			return true
		})
		// 窗体显示
		window.Show()
	})
}
