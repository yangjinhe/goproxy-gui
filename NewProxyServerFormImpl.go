// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/satori/go.uuid"
	"github.com/ying32/govcl/vcl"
	"strconv"
	"strings"
)

//::private::
type TNewProxyServerFormFields struct {
}

func (f *TNewProxyServerForm) OnFormCreate(sender vcl.IObject) {

}

func (f *TNewProxyServerForm) OnSelectCrtActionExecute(sender vcl.IObject) {
	NewProxyServerForm.SelectCrtDialog.SetFilter("证书文件(*.CRT)|*.CRT|全部文件(*.*)|*.*")
	NewProxyServerForm.SelectCrtDialog.SetDefaultExt("*.CRT")
	if NewProxyServerForm.SelectCrtDialog.Execute() {
		strFileName := NewProxyServerForm.SelectCrtDialog.FileName()
		if strFileName != "" {
			NewProxyServerForm.CrtEdit.SetText(strFileName)
		}
	}
}

func (f *TNewProxyServerForm) OnSelectKeyActionExecute(sender vcl.IObject) {
	NewProxyServerForm.SelectKeyDialog.SetFilter("密钥文件(*.KEY)|*.KEY|全部文件(*.*)|*.*")
	NewProxyServerForm.SelectKeyDialog.SetDefaultExt("*.KEY")
	if NewProxyServerForm.SelectKeyDialog.Execute() {
		strFileName := NewProxyServerForm.SelectKeyDialog.FileName()
		if strFileName != "" {
			NewProxyServerForm.KeyEdit.SetText(strFileName)
		}
	}
}

func (f *TNewProxyServerForm) OnEnableTlsActionExecute(sender vcl.IObject) {
	NewProxyServerForm.SelectCrtBtn.SetEnabled(NewProxyServerForm.TlsChk.Checked())
	NewProxyServerForm.SelectCrtBtn.SetEnabled(NewProxyServerForm.TlsChk.Checked())
}

func (f *TNewProxyServerForm) OnOkButtonClick(sender vcl.IObject) {
	configName := NewProxyServerForm.ConfigNameEdit.Text()
	if configName == "" {
		vcl.ShowMessage("配置文件名称必填")
		return
	}
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		panic("错误")
	}
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)
	key := u1.String()
	cfg.SetValue(key, "name", configName)
	cfg.SetValue(key, "LocalType", NewProxyServerForm.LocalTypeCombox.Text())
	cfg.SetValue(key, "LocalIp", NewProxyServerForm.LocalIpEdit.Text())
	cfg.SetValue(key, "LocalPort", NewProxyServerForm.LocalPortEdit.Text())
	cfg.SetValue(key, "SrcType", NewProxyServerForm.SrcTypeComboBox.Text())
	cfg.SetValue(key, "SrcIp", NewProxyServerForm.SrcIpEdit.Text())
	cfg.SetValue(key, "SrcPort", NewProxyServerForm.SrcPortEdit.Text())
	cfg.SetValue(key, "IsTls", strconv.FormatBool(NewProxyServerForm.TlsChk.Checked()))
	cfg.SetValue(key, "CrtFile", NewProxyServerForm.CrtEdit.Text())
	cfg.SetValue(key, "KeyFile", NewProxyServerForm.KeyEdit.Text())
	cfg.SetValue(key, "AuthAddr", NewProxyServerForm.AuthAddrEdit.Text())
	ss := strings.Fields(NewProxyServerForm.AuthMemo.Text())
	AuthUserPwd := ""
	for idx := range ss {
		AuthUserPwd += ss[idx] + ","
	}
	cfg.SetValue(key, "AuthUserPwd", AuthUserPwd)
	err = goconfig.SaveConfigFile(cfg, "config.ini")
	initListView()
	NewProxyServerForm.Close()
}

func (f *TNewProxyServerForm) OnCancelButtonClick(sender vcl.IObject) {

}


