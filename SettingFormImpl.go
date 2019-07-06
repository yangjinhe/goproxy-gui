// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TSettingFormFields struct {
}

func (f *TSettingForm) OnConfirmBtnClick(sender vcl.IObject) {
	var isAutoRunFlag = true
	var isSysProxyFlag = true
	if SettingForm.AutoStartChk.Checked() {
		isAutoRunFlag = setAutoRun()
	} else {
		isAutoRunFlag = unSetAutoRun()
	}
	addr := SettingForm.PortEdit.Text()
	if SettingForm.SysProxyChk.Checked() {
		isSysProxyFlag = setSysProxy(addr)
	} else {
		isSysProxyFlag = unsetSysProxy(addr)
	}
	/*cfg, _ := goconfig.LoadConfigFile(mainConfigName)
	cfg.SetValue("common", "autoRun", strconv.FormatBool(SettingForm.AutoStartChk.Checked()))
	cfg.SetValue("common", "sysProxy", strconv.FormatBool(SettingForm.SysProxyChk.Checked()))
	cfg.SetValue("common", "sysProxyAddr", addr)
	_ = goconfig.SaveConfigFile(cfg, mainConfigName)*/
	if isAutoRunFlag && isSysProxyFlag {
		vcl.MessageDlg("设置成功", types.MtInformation, types.MbOK)
		SettingForm.Close()
	} else {
		vcl.MessageDlg("设置失败，请尝试以管理员身份运行", types.MtError, types.MbOK)
	}
}

func (f *TSettingForm) OnCancelBtnClick(sender vcl.IObject) {
	SettingForm.Close()
}
