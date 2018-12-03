// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
    "github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl/version"
)

//::private::
type TAboutFormFields struct {
}


func (f *TAboutForm) OnFormCreate(sender vcl.IObject) {
	AboutForm.SetCaption(version.OSVersion.ToString())
}

