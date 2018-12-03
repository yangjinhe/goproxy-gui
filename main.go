// 由res2go自动生成。
package main

import (
	"context"
	"github.com/Unknwon/goconfig"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"io/ioutil"
	"os"
	"runtime"
)

var cancelGoroutine context.CancelFunc
var ctx context.Context

func main() {
	vcl.Application.SetFormScaled(true)
	vcl.Application.Initialize()
    vcl.Application.CreateForm(mainFormBytes, &MainForm)
    vcl.Application.CreateForm(aboutFormBytes, &AboutForm)
    vcl.Application.CreateForm(newProxyServerFormBytes, &NewProxyServerForm)
	initListView()

	icon := vcl.NewIcon()
	defer icon.Free()
	icon.LoadFromResourceName(rtl.MainInstance(), "MAINICON")

	MainForm.TrayIcon1.SetIcon(icon)

	// 捕捉最小化
	vcl.Application.SetOnMinimize(func(sender vcl.IObject) {
		MainForm.Hide() // 主窗口最小化掉
	})
	vcl.Application.Run()
}

func initListView() {
	b, err := PathExists("config.ini")
	if err != nil {
		panic("错误")
	}
	if !b {
		WriteWithIoutil("config.ini", "[common]")
		cfg, _ := goconfig.LoadConfigFile("config.ini")
		cfg.SetValue("common", "version", "v1.0")
	} else {
		MainForm.ServerListView.Clear()
		cfg, _ := goconfig.LoadConfigFile("config.ini")
		list := cfg.GetSectionList()
		MainForm.ServerListView.Items().BeginUpdate()
		for idx := range list {
			//fmt.Println(list[idx])
			if list[idx] == "common" {
				continue
			}
			lstitem := MainForm.ServerListView.Items().Add()
			lstitem.SetCaption(list[idx])
			lstitem.SubItems().Add(cfg.MustValue(list[idx], "name"))
			lstitem.SubItems().Add(cfg.MustValue(list[idx], "SrcIp"))
			lstitem.SubItems().Add(cfg.MustValue(list[idx], "SrcPort"))
			lstitem.SubItems().Add(cfg.MustValue(list[idx], "SrcType"))
			lstitem.SubItems().Add(cfg.MustValue(list[idx], "IsTls"))
			lstitem.SubItems().Add(cfg.MustValue(list[idx], "LocalType"))
			lstitem.SubItems().Add(cfg.MustValue(list[idx], "LocalPort"))
		}
		MainForm.ServerListView.Items().EndUpdate()
	}
	if err != nil {
		panic("错误")
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		//fmt.Println("写入文件成功:",content)
	}
}

func resetCtx() {
	if cancelGoroutine != nil {
		cancelGoroutine()
	}
	cleanEnv()
	ctx = context.Background()
	ctx, cancelGoroutine = context.WithCancel(ctx)
}

func cleanEnv() bool {
	switch runtime.GOOS {
	case "darwin":
		return true
	case "windows":
		execCommand(nil, "taskkill.exe", nil,"/f", "/im", "proxy.exe")
		execCommand(nil, "taskkill.exe", nil,"/f", "/im", "proxy-noconsole.exe")
		return true
	case "linux":
		execCommand(nil, "killall", nil,"proxy")
	}
	return true
}
