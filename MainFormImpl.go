// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/Unknwon/goconfig"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"runtime"
	"strconv"
	"strings"
)

//::private::
type TMainFormFields struct {
}

func (f *TMainForm) OnNewItemMenuClick(sender vcl.IObject) {
	NewProxyServerForm.ShowModal()
}

func (f *TMainForm) OnModifyItemMenuClick(sender vcl.IObject) {

}

func (f *TMainForm) OnExitMenuClick(sender vcl.IObject) {
	MainForm.Close()
}

func (f *TMainForm) OnAboutMenuClick(sender vcl.IObject) {
	AboutForm.ShowModal()
}

func (f *TMainForm) OnPopStartMenuClick(sender vcl.IObject) {
	serverName := MainForm.ServerListView.Selected().Caption()

	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		panic("错误")
	}
	localType := cfg.MustValue(serverName, "LocalType")
	localIp := cfg.MustValue(serverName, "LocalIp")
	localPort := cfg.MustValue(serverName, "LocalPort")
	srcType := cfg.MustValue(serverName, "SrcType")
	srcIp := cfg.MustValue(serverName, "SrcIp")
	srcPort := cfg.MustValue(serverName, "SrcPort")
	isTls := cfg.MustValue(serverName, "IsTls")
	crtFile := cfg.MustValue(serverName, "CrtFile")
	keyFile := cfg.MustValue(serverName, "KeyFile")
	authAddr := cfg.MustValue(serverName, "AuthAddr")
	authUserPwd := cfg.MustValue(serverName, "AuthUserPwd")

	var configContent = ""
	configContent += strings.ToLower(localType) + "\r\n"
	//configContent += "--debug\r\n"
	configContent += "--local=" + localIp + ":" + localPort + "\r\n"
	configContent += "--parent=" + srcIp + ":" + srcPort + "\r\n"
	configContent += "--parent-service-type=" + strings.ToLower(srcType) + "\r\n"
	b, err := strconv.ParseBool(isTls)
	if b {
		configContent += "--parent-type=tls" + "\r\n"
		configContent += "--cert=" + crtFile + "\r\n"
		configContent += "--key=" + keyFile + "\r\n"
	} else {
		configContent += "--parent-type=tcp" + "\r\n"
	}
	if len(authAddr) > 0 {
		configContent += "--auth-url=" + authAddr + "\r\n"
	}
	if len(authUserPwd) > 0 {
		split := strings.Split(authUserPwd, ",")
		for idx := range split {
			if len(split[idx]) > 0 {
				configContent += "--auth=" + split[idx] + "\r\n"
			}
		}
	}
	WriteWithIoutil("runconfig.txt", configContent)
	resetCtx()
	var cmd = ""
	if runtime.GOOS == "windows" {
		cmd = ".\\proxy.exe"
	} else {
		cmd = "./proxy"
	}
	MainForm.OutpuMemo.Lines().Clear()
	go execCommand(ctx, cmd, MainForm.OutpuMemo, "@runconfig.txt")
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
