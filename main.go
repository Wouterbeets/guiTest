package main

import (
	"encoding/json"
	"fmt"
	"github.com/andlabs/ui"
)

var window ui.Window

type ConfStruct struct {
	R6_user   string
	R6_pass   string
	R6_appkey string
}

var conf ConfStruct

func main() {
	go ui.Do(func() {
		R6_user := ui.NewTextField()
		R6_pass := ui.NewTextField()
		R6_appkey := ui.NewTextField()
		save := ui.NewButton("Save")
		userLabel := ui.NewLabel("radian6 username")
		passLabel := ui.NewLabel("password")
		appKey := ui.NewLabel("appkey")
		stack := ui.NewVerticalStack(
			userLabel,
			R6_user,
			passLabel,
			R6_pass,
			appKey,
			R6_appkey,
			save,
		)
		window = ui.NewWindow("Hello", 800, 400, stack)
		save.OnClicked(func() {
			conf.R6_user = R6_user.Text()
			conf.R6_pass = R6_pass.Text()
			conf.R6_appkey = R6_appkey.Text()
			js, err := json.Marshal(&conf)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(js))
		})
		window.OnClosing(func() bool {
			ui.Stop()
			return true
		})
		window.Show()
	})
	err := ui.Go()
	if err != nil {
		panic(err)
	}
}
