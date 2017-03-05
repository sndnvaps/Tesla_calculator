package main

import (
	"fmt"
	"os"
	"log"
	"github.com/visualfc/goqt/ui"
	"runtime"
)

func main() {
	ui.RunEx(os.Args, func() {
		w, err := NewMainForm()
		if err != nil {
			log.Fatalln(err)
		}
		w.Show()
	})
}

func version() {
	info := fmt.Sprintf("GoQt Version %v \nQt Version %v\ngo verison %v %v/%v",
		ui.Version(),
		ui.QtVersion(),
		runtime.Version(), runtime.GOOS, runtime.GOARCH)

	lable := ui.NewLabel()
	lable.SetText(info)

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(lable)

	widget := ui.NewWidget()
	widget.SetLayout(hbox)
	widget.Show()
}

