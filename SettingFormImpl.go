// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/Unknwon/goconfig"
	"github.com/ying32/govcl/vcl"
	"strconv"
)

//::private::
type TSettingFormFields struct {
}

func (f *TSettingForm) OnConfirmBtnClick(sender vcl.IObject) {
	if SettingForm.AutoStartChk.Checked() {
		setAutoRun()
	} else {
		unSetAutoRun()
	}
	addr := SettingForm.PortEdit.Text()
	if SettingForm.SysProxyChk.Checked() {
		setSysProxy(addr)
	} else {
		unsetSysProxy(addr)
	}
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		panic("错误")
	}
	cfg.SetValue("common", "autoRun", strconv.FormatBool(SettingForm.AutoStartChk.Checked()))
	cfg.SetValue("common", "sysProxy", strconv.FormatBool(SettingForm.SysProxyChk.Checked()))
	cfg.SetValue("common", "sysProxyAddr", addr)
}

func (f *TSettingForm) OnCancelBtnClick(sender vcl.IObject) {
	SettingForm.Close()
}
