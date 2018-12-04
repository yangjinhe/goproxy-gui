// 由res2go自动生成。
package main

import (
	"fmt"
	"github.com/ProtonMail/go-autostart"
	"github.com/Unknwon/goconfig"
	"github.com/getlantern/sysproxy"
	"github.com/snail007/goproxy/sdk/android-ios"
	"github.com/ying32/govcl/vcl"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var currentServerName string
var currentEditServerName string
var currentLocalIpPort string

func main() {
	vcl.Application.SetFormScaled(true)
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(mainFormBytes, &MainForm)
	vcl.Application.CreateForm(aboutFormBytes, &AboutForm)
	vcl.Application.CreateForm(newProxyServerFormBytes, &NewProxyServerForm)
	vcl.Application.CreateForm(settingFormBytes, &SettingForm)
	initListView()
	MainForm.SetCaption("ProxyGo-GUI")

	MainForm.TrayIcon1.SetIcon(vcl.Application.Icon())
	MainForm.TrayIcon1.SetHint(MainForm.Caption())
	MainForm.TrayIcon1.SetVisible(true)
	// 捕捉最小化
	vcl.Application.SetOnMinimize(func(sender vcl.IObject) {
		MainForm.Hide() // 主窗口最小化掉
	})
	loadSysSetting()
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
	//cleanEnv()
	if currentServerName != "" {
		go proxy.Stop(currentServerName)
	}
}

func loadSysSetting() {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		panic("错误")
	}
	autoRun := cfg.MustBool("common", "autoRun", false)
	sysProxy := cfg.MustBool("common", "sysProxy", false)
	sysProxyAddr := cfg.MustValue("common", "sysProxyAddr", "")
	fmt.Println(autoRun)
	fmt.Println(sysProxy)
	if autoRun {
		setAutoRun()
	} else {
		unSetAutoRun()
	}
	if sysProxy {
		setSysProxy(sysProxyAddr)
	}

}

func setAutoRun() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	app := &autostart.App{
		Name:        "goproxy-gui",
		DisplayName: "goproxy的图形界面",
		Exec:        []string{strings.Join([]string{dir, os.Args[0]}, string(os.PathSeparator))},
	}

	if app.IsEnabled() {
		log.Println("App is already enabled, removing it...")
	} else {
		log.Println("Enabling app...")

		if err := app.Enable(); err != nil {
			log.Fatal(err)
		}
	}
}

func unSetAutoRun() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	app := &autostart.App{
		Name:        "goproxy-gui",
		DisplayName: "goproxy的图形界面",
		Exec:        []string{strings.Join([]string{dir, os.Args[0]}, string(os.PathSeparator))},
	}

	if app.IsEnabled() {
		log.Println("App is already enabled, removing it...")

		if err := app.Disable(); err != nil {
			log.Fatal(err)
		}
	}
}

func setSysProxy(addr string) {
	helperFullPath := "sysproxy-cmd"
	iconFullPath, _ := filepath.Abs("./icon.png")
	log.Println("Using icon at %v", iconFullPath)
	err := sysproxy.EnsureHelperToolPresent(helperFullPath, "Input your password and save the world!", iconFullPath)
	if err != nil {
		fmt.Printf("Error EnsureHelperToolPresent: %s\n", err)
		return
	}
	_, _ = sysproxy.On(addr)
	if err != nil {
		fmt.Printf("Error set proxy: %s\n", err)
		return
	}
}

func unsetSysProxy(addr string) {
	helperFullPath := "sysproxy-cmd"
	iconFullPath, _ := filepath.Abs("./icon.png")
	//log.Println("Using icon at %v", iconFullPath)
	err := sysproxy.EnsureHelperToolPresent(helperFullPath, "Input your password and save the world!", iconFullPath)
	if err != nil {
		fmt.Printf("Error EnsureHelperToolPresent: %s\n", err)
		return
	}
	off, err := sysproxy.On(addr)
	if err != nil {
		fmt.Printf("Error set proxy: %s\n", err)
		return
	}
	err1 := off()
	if err1 != nil {
		fmt.Printf("Error unset proxy: %s\n", err1)
		return
	}
}
