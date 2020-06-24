// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
    "github.com/ying32/govcl/vcl"
)

//::private::
type TMainFormFields struct {
}


func (f *TMainForm) OnSettingsBtnClick(sender vcl.IObject) {
	SettingForm.ShowModal()
}


func (f *TMainForm) OnUpdateBtnClick(sender vcl.IObject) {

}


func (f *TMainForm) OnHelpBtnClick(sender vcl.IObject) {

}


func (f *TMainForm) OnAddServerClick(sender vcl.IObject) {
    NewProxyServerForm.ShowModal()
}


func (f *TMainForm) OnSetActivityServerClick(sender vcl.IObject) {

}


func (f *TMainForm) OnPingServerClick(sender vcl.IObject) {

}


func (f *TMainForm) OnShowMainWindowClick(sender vcl.IObject) {

}


func (f *TMainForm) OnExitMainWindowClick(sender vcl.IObject) {

}

