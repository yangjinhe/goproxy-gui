// 由res2go自动生成。
package main

import (
	"fmt"
	"github.com/ProtonMail/go-autostart"
	"github.com/Unknwon/goconfig"
	"github.com/getlantern/sysproxy"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var currentServerName string
var currentEditServerName string
var mainConfigName = "config.ini"

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

	if rtl.LcLLoaded() {
		if runtime.GOOS != "windows" {
			icon := vcl.NewIcon()
			icon.LoadFromFile(rtl.ExtractFilePath(vcl.Application.ExeName()) + "/1.ico")
			MainForm.TrayIcon1.SetIcon(icon)
			icon.Free()
		} else {
			MainForm.TrayIcon1.SetIcon(vcl.Application.Icon())
		}
	}

	MainForm.TrayIcon1.SetHint(MainForm.Caption())
	MainForm.TrayIcon1.SetVisible(true)
	// 捕捉最小化

	vcl.Application.SetOnMinimize(func(sender vcl.IObject) {
		if runtime.GOOS == "windows" {
			MainForm.Hide() // 主窗口最小化掉
		}
	})
	loadSysSetting()
	vcl.Application.Run()
}

func initListView() {
	b, err := PathExists(mainConfigName)
	if err != nil {
		panic("错误")
	}
	if !b {
		WriteWithIoutil(mainConfigName, "[common]")
		cfg, _ := goconfig.LoadConfigFile(mainConfigName)
		cfg.SetValue("common", "version", "v1.0")
	} else {
		MainForm.ServerListView.Clear()
		cfg, _ := goconfig.LoadConfigFile(mainConfigName)
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
			if cfg.MustBool(list[idx], "Status", false) {
				lstitem.SubItems().Add("运行中")
			} else {
				lstitem.SubItems().Add("未启动")
			}

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

func loadSysSetting() {
	cfg, err := goconfig.LoadConfigFile(mainConfigName)
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

func setAutoRun() bool {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(dir)
	split := strings.Split(os.Args[0], string(os.PathSeparator))
	exeName := ""
	if len(split) > 0 {
		fmt.Println(split[len(split)-1])
		exeName = split[len(split)-1]
	} else {
		exeName = os.Args[0]
	}
	fmt.Println(exeName)

	var exec = dir + string(os.PathSeparator) + exeName
	app := &autostart.App{
		Name:        "goproxy-gui",
		DisplayName: "goproxy的图形界面",
		Exec:        []string{exec},
	}

	if app.IsEnabled() {
		log.Println("App is already enabled, removing it...")
	} else {
		log.Println("Enabling app...")

		if err := app.Enable(); err != nil {
			log.Fatal(err)
			return false
		}
	}
	return true
}

func unSetAutoRun() bool {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(dir)
	split := strings.Split(os.Args[0], string(os.PathSeparator))
	exeName := ""
	if len(split) > 0 {
		fmt.Println(split[len(split)-1])
		exeName = split[len(split)-1]
	} else {
		exeName = os.Args[0]
	}
	fmt.Println(exeName)

	var exec = dir + string(os.PathSeparator) + exeName
	app := &autostart.App{
		Name:        "goproxy-gui",
		DisplayName: "goproxy的图形界面",
		Exec:        []string{exec},
	}

	if app.IsEnabled() {
		log.Println("App is already enabled, removing it...")

		if err := app.Disable(); err != nil {
			log.Fatal(err)
			return false
		}
	}
	return true
}

func setSysProxy(addr string) bool {
	helperFullPath := "sysproxy-cmd"
	iconFullPath, _ := filepath.Abs("./icon.png")
	log.Println("Using icon at %v", iconFullPath)
	err := sysproxy.EnsureHelperToolPresent(helperFullPath, "Input your password and save the world!", iconFullPath)
	if err != nil {
		fmt.Printf("Error EnsureHelperToolPresent: %s\n", err)
		return false
	}
	_, err1 := sysproxy.On(addr)
	if err1 != nil {
		fmt.Printf("Error set proxy: %s\n", err)
		return false
	}
	return true
}

func unsetSysProxy(addr string) bool {
	helperFullPath := "sysproxy-cmd"
	iconFullPath, _ := filepath.Abs("./icon.png")
	//log.Println("Using icon at %v", iconFullPath)
	err := sysproxy.EnsureHelperToolPresent(helperFullPath, "Input your password and save the world!", iconFullPath)
	if err != nil {
		fmt.Printf("Error EnsureHelperToolPresent: %s\n", err)
		return false
	}
	off, err1 := sysproxy.On(addr)
	if err1 != nil {
		fmt.Printf("Error set proxy: %s\n", err1)
		return false
	}
	err2 := off()
	if err2 != nil {
		fmt.Printf("Error unset proxy: %s\n", err2)
		return false
	}
	return true
}
