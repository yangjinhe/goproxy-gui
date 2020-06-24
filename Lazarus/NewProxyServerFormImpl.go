// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
    "github.com/ying32/govcl/vcl"
)

//::private::
type TNewProxyServerFormFields struct {
}


func (f *TNewProxyServerForm) OnServerProtocolComboBoxChange(sender vcl.IObject) {
    text := NewProxyServerForm.ServerProtocolComboBox.Text()
    switch text {
    case "tls":
        NewProxyServerForm.PageControl1.SetVisible(true)
        NewProxyServerForm.PageControl1.SetActivePageIndex(0)
        break;
    case "ssh":
        NewProxyServerForm.PageControl1.SetVisible(true)
        NewProxyServerForm.PageControl1.SetActivePageIndex(1)
        break;
    case "kcp":
        NewProxyServerForm.PageControl1.SetVisible(true)
        NewProxyServerForm.PageControl1.SetActivePageIndex(2)
        break;
    default:
        NewProxyServerForm.PageControl1.SetVisible(false)
    }
}


func (f *TNewProxyServerForm) OnLocalProtocolComboBoxChange(sender vcl.IObject) {

}


func (f *TNewProxyServerForm) OnConfirmBtnClick(sender vcl.IObject) {
    proxyType := NewProxyServerForm.ProxyTypeComboBox.Text()
    remoteIp := NewProxyServerForm.IpInput.Text()
    remotePort := NewProxyServerForm.PortInput.Text()
    remoteProtocol := NewProxyServerForm.ServerProtocolComboBox.Text()

    crtLocation := NewProxyServerForm.CrtInput.Text()
    keyLocation := NewProxyServerForm.KeyInput.Text()

    sshPrivateKey := NewProxyServerForm.SshKeyInput.Text()
    sshUserName := NewProxyServerForm.SshUserNameInput.Text()
    sshPwd := NewProxyServerForm.SshPwdInput.Text()

    kcpPwd := NewProxyServerForm.KcpPwdInput.Text()

    localProtocol := NewProxyServerForm.LocalProtocolComboBox.Text()
    localIp := NewProxyServerForm.LocalIp.Text()
    localPort := NewProxyServerForm.LocalPort.Text()

    dns := NewProxyServerForm.DnsInput.Text()
    dnsTtl := NewProxyServerForm.TTLInput.Text()

    p := ProxyInfo{0,
        "",
        proxyType, remoteIp,remotePort, remoteProtocol,
        localIp, localPort, localProtocol, crtLocation, keyLocation,
        sshUserName, sshPwd, sshPrivateKey, kcpPwd, dns, dnsTtl}
    addNewProxyInfo(p)
}


func (f *TNewProxyServerForm) OnCancelBtnClick(sender vcl.IObject) {

}


func (f *TNewProxyServerForm) OnProxyTypeComboBoxChange(sender vcl.IObject) {
    text := NewProxyServerForm.ProxyTypeComboBox.Text()
    if text == "穿透" {
        NewProxyServerForm.GroupBox2.SetVisible(false)
        NewProxyServerForm.GroupBox3.SetVisible(false)
        NewProxyServerForm.GroupBox4.SetVisible(true)
    } else {
        NewProxyServerForm.GroupBox2.SetVisible(true)
        NewProxyServerForm.GroupBox3.SetVisible(true)
        NewProxyServerForm.GroupBox4.SetVisible(false)
    }
}

func (f *TNewProxyServerForm) OnAddMappingClick(sender vcl.IObject) {
    AddMappingForm.ShowModal()
}


func (f *TNewProxyServerForm) OnModifyMappingClick(sender vcl.IObject) {

}


func (f *TNewProxyServerForm) OnDelMappingClick(sender vcl.IObject) {

}


func (f *TNewProxyServerForm) OnSelectCrtActionExecute(sender vcl.IObject) {
    NewProxyServerForm.SelectCrtDialog.SetFilter("证书文件(*.CRT;*.CER)|*.CRT;*.CER|全部文件(*.*)|*.*")
    NewProxyServerForm.SelectCrtDialog.SetDefaultExt("*.CRT")
    if NewProxyServerForm.SelectCrtDialog.Execute() {
        strFileName := NewProxyServerForm.SelectCrtDialog.FileName()
        if strFileName != "" {
            NewProxyServerForm.CrtInput.SetText(strFileName)
        }
    }
}


func (f *TNewProxyServerForm) OnSelectKeyActionExecute(sender vcl.IObject) {
    NewProxyServerForm.SelectKeyDialog.SetFilter("密钥文件(*.KEY)|*.KEY|全部文件(*.*)|*.*")
    NewProxyServerForm.SelectKeyDialog.SetDefaultExt("*.KEY")
    if NewProxyServerForm.SelectKeyDialog.Execute() {
        strFileName := NewProxyServerForm.SelectKeyDialog.FileName()
        if strFileName != "" {
            NewProxyServerForm.KeyInput.SetText(strFileName)
        }
    }
}


func (f *TNewProxyServerForm) OnSelectPrivateKeyActionExecute(sender vcl.IObject) {
    NewProxyServerForm.SelectKeyDialog.SetFilter("密钥文件(*.PEM)|*.PEM|全部文件(*.*)|*.*")
    NewProxyServerForm.SelectKeyDialog.SetDefaultExt("*.PEM")
    if NewProxyServerForm.SelectKeyDialog.Execute() {
        strFileName := NewProxyServerForm.SelectKeyDialog.FileName()
        if strFileName != "" {
            NewProxyServerForm.SshKeyInput.SetText(strFileName)
        }
    }
}


func (f *TNewProxyServerForm) OnSelectCrtBtnClick(sender vcl.IObject) {

}

