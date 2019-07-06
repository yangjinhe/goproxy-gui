unit Unit2;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ComCtrls,
  Menus, ActnList, StdActns;

type

  { TNewProxyServerForm }

  TNewProxyServerForm = class(TForm)
    Label15: TLabel;
    ProxyTypeComboBox: TComboBox;
    SelectCrtAction: TAction;
    SelectKeyAction: TAction;
    SelectPrivateKeyAction: TAction;
    ActionList1: TActionList;
    SelectCrtBtn: TButton;
    SelectKeyBtn: TButton;
    SelectPrivateKeyBtn: TButton;
    ConfirmBtn: TButton;
    CancelBtn: TButton;
    GroupBox4: TGroupBox;
    ListView1: TListView;
    AddMapping: TMenuItem;
    ModifyMapping: TMenuItem;
    DelMapping: TMenuItem;
    SelectCrtDialog: TOpenDialog;
    SelectKeyDialog: TOpenDialog;
    SelectPrivateKeyDialog: TOpenDialog;
    PopupMenu1: TPopupMenu;
    LocalProtocolComboBox: TComboBox;
    DnsInput: TEdit;
    LocalPort: TEdit;
    Label13: TLabel;
    Label14: TLabel;
    LocalIp: TEdit;
    Label12: TLabel;
    TTLInput: TEdit;
    GroupBox2: TGroupBox;
    GroupBox3: TGroupBox;
    KcpPwdInput: TEdit;
    Label10: TLabel;
    Label11: TLabel;
    KcpPwdLabel: TLabel;
    ServerProtocolComboBox: TComboBox;
    IpInput: TEdit;
    PortInput: TEdit;
    CrtInput: TEdit;
    KeyInput: TEdit;
    SshUserNameInput: TEdit;
    SshPwdInput: TEdit;
    SshKeyInput: TEdit;
    GroupBox1: TGroupBox;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    CrtLabel: TLabel;
    KeyLabel: TLabel;
    SshUserNameLabel: TLabel;
    SshPwdLabel: TLabel;
    SshKeyLabel: TLabel;
    procedure AddMappingClick(Sender: TObject);
    procedure DelMappingClick(Sender: TObject);
    procedure LocalProtocolComboBoxChange(Sender: TObject);
    procedure ModifyMappingClick(Sender: TObject);
    procedure ProxyTypeComboBoxChange(Sender: TObject);
    procedure SelectCrtActionExecute(Sender: TObject);
    procedure SelectCrtBtnClick(Sender: TObject);
    procedure CancelBtnClick(Sender: TObject);
    procedure ConfirmBtnClick(Sender: TObject);
    procedure SelectKeyActionExecute(Sender: TObject);
    procedure SelectPrivateKeyActionExecute(Sender: TObject);
    procedure ServerProtocolComboBoxChange(Sender: TObject);
  private

  public

  end;

var
  NewProxyServerForm: TNewProxyServerForm;

implementation

{$R *.lfm}

{ TNewProxyServerForm }

procedure TNewProxyServerForm.ConfirmBtnClick(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.SelectKeyActionExecute(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.SelectPrivateKeyActionExecute(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.ServerProtocolComboBoxChange(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.CancelBtnClick(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.SelectCrtBtnClick(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.SelectCrtActionExecute(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.ProxyTypeComboBoxChange(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.LocalProtocolComboBoxChange(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.ModifyMappingClick(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.AddMappingClick(Sender: TObject);
begin

end;

procedure TNewProxyServerForm.DelMappingClick(Sender: TObject);
begin

end;

end.

