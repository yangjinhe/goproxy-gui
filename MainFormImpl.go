// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/snail007/goproxy/sdk/android-ios"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"strconv"
	"strings"
)

//::private::
type TMainFormFields struct {
}

type ServerInfo struct {
	localType   string
	localIp     string
	localPort   string
	srcType     string
	srcIp       string
	srcPort     string
	isTls       string
	crtFile     string
	keyFile     string
	authAddr    string
	authUserPwd string
	name        string
	uuid        string
}

func (f *TMainForm) OnNewItemMenuClick(sender vcl.IObject) {
	NewProxyServerForm.ShowModal()
}

func (f *TMainForm) OnModifyItemMenuClick(sender vcl.IObject) {

}

func (f *TMainForm) OnSettingMenuClick(sender vcl.IObject) {
	SettingForm.ShowModal()
}

func (f *TMainForm) OnExitMenuClick(sender vcl.IObject) {
	MainForm.Close()
}

func (f *TMainForm) OnAboutMenuClick(sender vcl.IObject) {
	AboutForm.ShowModal()
}

func (f *TMainForm) OnPopStartMenuClick(sender vcl.IObject) {
	currentServerName = MainForm.ServerListView.Selected().Caption()
	serverInfo := readConfig(currentServerName)
	configContent := buildConfig(serverInfo, " ")
	resetCtx()
	MainForm.OutpuMemo.Lines().Clear()
	var loggerCallback proxy.LogCallback
	loggerCallback = new(MyLogCallback)
	errStr := proxy.StartWithLog(currentServerName, configContent, loggerCallback)
	if errStr != "" {
		fmt.Println("启动失败：" + errStr)
	}
}

func readConfig(serverName string) ServerInfo {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		panic("错误")
	}
	var serverInfo ServerInfo
	serverInfo.uuid = serverName
	serverInfo.name = cfg.MustValue(serverName, "name")
	serverInfo.localType = cfg.MustValue(serverName, "LocalType")
	serverInfo.localIp = cfg.MustValue(serverName, "LocalIp")
	serverInfo.localPort = cfg.MustValue(serverName, "LocalPort")
	serverInfo.srcType = cfg.MustValue(serverName, "SrcType")
	serverInfo.srcIp = cfg.MustValue(serverName, "SrcIp")
	serverInfo.srcPort = cfg.MustValue(serverName, "SrcPort")
	serverInfo.isTls = cfg.MustValue(serverName, "IsTls")
	serverInfo.crtFile = cfg.MustValue(serverName, "CrtFile")
	serverInfo.keyFile = cfg.MustValue(serverName, "KeyFile")
	serverInfo.authAddr = cfg.MustValue(serverName, "AuthAddr")
	serverInfo.authUserPwd = cfg.MustValue(serverName, "AuthUserPwd")

	currentLocalIpPort = serverInfo.localIp + ":" + serverInfo.localPort
	return serverInfo
}

func buildConfig(serverInfo ServerInfo, sep string) string {
	var configContent = ""
	configContent += strings.ToLower(serverInfo.localType) + sep
	//configContent += "--debug\r\n"
	configContent += "--local=" + serverInfo.localIp + ":" + serverInfo.localPort + sep
	configContent += "--parent=" + serverInfo.srcIp + ":" + serverInfo.srcPort + sep
	configContent += "--parent-service-type=" + strings.ToLower(serverInfo.srcType) + sep
	b, _ := strconv.ParseBool(serverInfo.isTls)
	if b {
		configContent += "--parent-type=tls" + sep
		configContent += "--cert=" + serverInfo.crtFile + sep
		configContent += "--key=" + serverInfo.keyFile + sep
	} else {
		configContent += "--parent-type=tcp" + sep
	}
	if len(serverInfo.authAddr) > 0 {
		configContent += "--auth-url=" + serverInfo.authAddr + sep
	}
	if len(serverInfo.authUserPwd) > 0 {
		split := strings.Split(serverInfo.authUserPwd, ",")
		for idx := range split {
			if len(split[idx]) > 0 {
				configContent += "--auth=" + split[idx] + sep
			}
		}
	}
	return configContent
}

type MyLogCallback struct {
}

func (myLogCallback MyLogCallback) Write(line string) {
	MainForm.OutpuMemo.Lines().Append(line)
}

func (f *TMainForm) OnPopStopMenuClick(sender vcl.IObject) {
	resetCtx()
}

func (f *TMainForm) OnPopExitMenuClick(sender vcl.IObject) {
	MainForm.Close()
}

func (f *TMainForm) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	b := vcl.MessageDlg("是否退出?", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes
	if b {
		resetCtx()
	}
	*canClose = b
}

func (f *TMainForm) OnServerListViewMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if MainForm.ServerListView.Selected().IsValid() {
		if button == types.MbRight {
			MainForm.PopupMenu1.Popup(vcl.Mouse.CursorPos().X, vcl.Mouse.CursorPos().Y)
		}
	}
}

func (f *TMainForm) OnServerListViewDblClick(sender vcl.IObject) {
	if MainForm.ServerListView.Selected().IsValid() {
		currentEditServerName = MainForm.ServerListView.Selected().Caption()

		serverInfo := readConfig(currentEditServerName)
		fmt.Println(serverInfo.localIp)
		NewProxyServerForm.ConfigNameEdit.SetText(serverInfo.name)
		NewProxyServerForm.LocalTypeCombox.SetText(serverInfo.localType)
		NewProxyServerForm.LocalIpEdit.SetText(serverInfo.localIp)
		NewProxyServerForm.LocalPortEdit.SetText(serverInfo.localPort)
		NewProxyServerForm.SrcTypeComboBox.SetText(serverInfo.srcType)
		NewProxyServerForm.SrcIpEdit.SetText(serverInfo.srcIp)
		NewProxyServerForm.SrcPortEdit.SetText(serverInfo.srcPort)
		b, _ := strconv.ParseBool(serverInfo.isTls)
		NewProxyServerForm.TlsChk.SetChecked(b)
		NewProxyServerForm.CrtEdit.SetText(serverInfo.crtFile)
		NewProxyServerForm.KeyEdit.SetText(serverInfo.keyFile)
		NewProxyServerForm.AuthAddrEdit.SetText(serverInfo.authAddr)
		NewProxyServerForm.AuthMemo.SetText(serverInfo.authUserPwd)
	}
	NewProxyServerForm.ShowModal()

}

func (f *TMainForm) OnExitPop2MenuClick(sender vcl.IObject) {
	MainForm.Close()
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

}

func (f *TMainForm) OnResetFormMenuClick(sender vcl.IObject) {
	MainForm.Show()
}

func (f *TMainForm) OnTrayIcon1DblClick(sender vcl.IObject) {
	MainForm.Show()
}

func (f *TMainForm) OnTrayIcon1Click(sender vcl.IObject) {

}

func (f *TMainForm) OnPopDeleteMenuClick(sender vcl.IObject) {
	if MainForm.ServerListView.Selected().IsValid() {
		b := vcl.MessageDlg("确定删除吗?", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes
		if b {
			cfg, err := goconfig.LoadConfigFile("config.ini")
			if err != nil {
				panic("错误")
			}
			fmt.Println(MainForm.ServerListView.Selected().Caption())
			deleteSection := cfg.DeleteSection(MainForm.ServerListView.Selected().Caption())
			fmt.Println(deleteSection)
			initListView()
		}
	}
}
