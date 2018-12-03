// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TMainForm struct {
    *vcl.TForm
    GroupBox1      *vcl.TGroupBox
    ServerListView *vcl.TListView
    GroupBox2      *vcl.TGroupBox
    OutpuMemo      *vcl.TMemo
    MainMenu1      *vcl.TMainMenu
    FileMenu       *vcl.TMenuItem
    NewItemMenu    *vcl.TMenuItem
    ModifyItemMenu *vcl.TMenuItem
    ExitMenu       *vcl.TMenuItem
    SettingMenu    *vcl.TMenuItem
    HelpMenu       *vcl.TMenuItem
    AboutMenu      *vcl.TMenuItem
    TrayIcon1      *vcl.TTrayIcon
    PopupMenu1     *vcl.TPopupMenu
    PopStartMenu   *vcl.TMenuItem
    PopStopMenu    *vcl.TMenuItem
    PopExitMenu    *vcl.TMenuItem
    PopupMenu2     *vcl.TPopupMenu
    ResetFormMenu  *vcl.TMenuItem
    ExitPop2Menu   *vcl.TMenuItem

    //::private::
    TMainFormFields
}

var MainForm *TMainForm




// 以字节形式加载
// vcl.Application.CreateForm(mainFormBytes, &MainForm)

var mainFormBytes = []byte("\x54\x50\x46\x30\x09\x54\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x08\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x03\xCB\x01\x06\x48\x65\x69\x67\x68\x74\x03\xE3\x01\x03\x54\x6F\x70\x03\x23\x01\x05\x57\x69\x64\x74\x68\x03\xAF\x02\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x0A\x62\x69\x4D\x69\x6E\x69\x6D\x69\x7A\x65\x00\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x08\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xCA\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xAF\x02\x0D\x44\x65\x73\x69\x67\x6E\x54\x69\x6D\x65\x50\x50\x49\x02\x78\x04\x4D\x65\x6E\x75\x07\x09\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x31\x08\x4F\x6E\x43\x72\x65\x61\x74\x65\x07\x0A\x46\x6F\x72\x6D\x43\x72\x65\x61\x74\x65\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x07\x31\x2E\x38\x2E\x34\x2E\x30\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x31\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\xE8\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xAB\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\xE6\x9C\x8D\xE5\x8A\xA1\xE5\x99\xA8\xE5\x88\x97\xE8\xA1\xA8\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xCF\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xA7\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x09\x54\x4C\x69\x73\x74\x56\x69\x65\x77\x0E\x53\x65\x72\x76\x65\x72\x4C\x69\x73\x74\x56\x69\x65\x77\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xD0\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xA8\x02\x07\x43\x6F\x6C\x75\x6D\x6E\x73\x0E\x01\x07\x56\x69\x73\x69\x62\x6C\x65\x08\x05\x57\x69\x64\x74\x68\x02\x00\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x90\x8D\xE7\xA7\xB0\x05\x57\x69\x64\x74\x68\x02\x78\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE4\xB8\x8A\xE7\xBA\xA7\xE5\x9C\xB0\xE5\x9D\x80\x05\x57\x69\x64\x74\x68\x03\xC8\x00\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE4\xB8\x8A\xE7\xBA\xA7\xE7\xAB\xAF\xE5\x8F\xA3\x05\x57\x69\x64\x74\x68\x02\x50\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE4\xB8\x8A\xE7\xBA\xA7\xE5\x8D\x8F\xE8\xAE\xAE\x05\x57\x69\x64\x74\x68\x02\x50\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x54\x4C\x53\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x9C\xAC\xE5\x9C\xB0\xE5\x8D\x8F\xE8\xAE\xAE\x05\x57\x69\x64\x74\x68\x02\x50\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x9C\xAC\xE5\x9C\xB0\xE7\xAB\xAF\xE5\x8F\xA3\x05\x57\x69\x64\x74\x68\x02\x50\x00\x00\x0D\x49\x74\x65\x6D\x73\x2E\x4C\x61\x7A\x44\x61\x74\x61\x0A\x20\x00\x00\x00\x20\x00\x00\x00\x01\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\x00\x00\x00\x00\x04\x00\x00\x00\x74\x65\x73\x74\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x09\x52\x6F\x77\x53\x65\x6C\x65\x63\x74\x09\x0A\x53\x63\x72\x6F\x6C\x6C\x42\x61\x72\x73\x07\x0E\x73\x73\x41\x75\x74\x6F\x56\x65\x72\x74\x69\x63\x61\x6C\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x09\x56\x69\x65\x77\x53\x74\x79\x6C\x65\x07\x08\x76\x73\x52\x65\x70\x6F\x72\x74\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x32\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\xDF\x00\x03\x54\x6F\x70\x03\xE8\x00\x05\x57\x69\x64\x74\x68\x03\xAB\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xBE\x93\xE5\x87\xBA\xE7\xAA\x97\xE5\x8F\xA3\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xC6\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xA7\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x05\x54\x4D\x65\x6D\x6F\x09\x4F\x75\x74\x70\x75\x4D\x65\x6D\x6F\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xC8\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xA8\x02\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x0A\x53\x63\x72\x6F\x6C\x6C\x42\x61\x72\x73\x07\x06\x73\x73\x42\x6F\x74\x68\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x09\x54\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x09\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x31\x04\x6C\x65\x66\x74\x03\x80\x01\x03\x74\x6F\x70\x03\x10\x01\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x08\x46\x69\x6C\x65\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x96\x87\xE4\xBB\xB6\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0B\x4E\x65\x77\x49\x74\x65\x6D\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x96\xB0\xE5\xBB\xBA\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x10\x4E\x65\x77\x49\x74\x65\x6D\x4D\x65\x6E\x75\x43\x6C\x69\x63\x6B\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0E\x4D\x6F\x64\x69\x66\x79\x49\x74\x65\x6D\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE4\xBF\xAE\xE6\x94\xB9\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x13\x4D\x6F\x64\x69\x66\x79\x49\x74\x65\x6D\x4D\x65\x6E\x75\x43\x6C\x69\x63\x6B\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x08\x45\x78\x69\x74\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE9\x80\x80\xE5\x87\xBA\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0D\x45\x78\x69\x74\x4D\x65\x6E\x75\x43\x6C\x69\x63\x6B\x00\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0B\x53\x65\x74\x74\x69\x6E\x67\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xAE\xBE\xE7\xBD\xAE\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x08\x48\x65\x6C\x70\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\xB8\xAE\xE5\x8A\xA9\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x09\x41\x62\x6F\x75\x74\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x85\xB3\xE4\xBA\x8E\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0E\x41\x62\x6F\x75\x74\x4D\x65\x6E\x75\x43\x6C\x69\x63\x6B\x00\x00\x00\x00\x09\x54\x54\x72\x61\x79\x49\x63\x6F\x6E\x09\x54\x72\x61\x79\x49\x63\x6F\x6E\x31\x09\x50\x6F\x70\x55\x70\x4D\x65\x6E\x75\x07\x0A\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x32\x07\x56\x69\x73\x69\x62\x6C\x65\x09\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0E\x54\x72\x61\x79\x49\x63\x6F\x6E\x31\x43\x6C\x69\x63\x6B\x04\x6C\x65\x66\x74\x03\xD8\x01\x03\x74\x6F\x70\x03\x10\x01\x00\x00\x0A\x54\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x0A\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x31\x04\x6C\x65\x66\x74\x03\x20\x01\x03\x74\x6F\x70\x03\x10\x01\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0C\x50\x6F\x70\x53\x74\x61\x72\x74\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x90\xAF\xE5\x8A\xA8\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x11\x50\x6F\x70\x53\x74\x61\x72\x74\x4D\x65\x6E\x75\x43\x6C\x69\x63\x6B\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0B\x50\x6F\x70\x53\x74\x6F\x70\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x81\x9C\xE6\xAD\xA2\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x10\x50\x6F\x70\x53\x74\x6F\x70\x4D\x65\x6E\x75\x43\x6C\x69\x63\x6B\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0B\x50\x6F\x70\x45\x78\x69\x74\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE9\x80\x80\xE5\x87\xBA\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x10\x50\x6F\x70\x45\x78\x69\x74\x4D\x65\x6E\x75\x43\x6C\x69\x63\x6B\x00\x00\x00\x0A\x54\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x0A\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x32\x04\x6C\x65\x66\x74\x03\x38\x02\x03\x74\x6F\x70\x03\x10\x01\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0D\x52\x65\x73\x65\x74\x46\x6F\x72\x6D\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xBF\x98\xE5\x8E\x9F\xE7\xAA\x97\xE5\x8F\xA3\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x12\x52\x65\x73\x65\x74\x46\x6F\x72\x6D\x4D\x65\x6E\x75\x43\x6C\x69\x63\x6B\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0C\x45\x78\x69\x74\x50\x6F\x70\x32\x4D\x65\x6E\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE9\x80\x80\xE5\x87\xBA\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x11\x45\x78\x69\x74\x50\x6F\x70\x32\x4D\x65\x6E\x75\x43\x6C\x69\x63\x6B\x00\x00\x00\x00")