unit Unit2;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ComCtrls,
  Menus, ActnList, StdActns;

type

  { TNewProxyServerForm }

  TNewProxyServerForm = class(TForm)
    CrtInput: TEdit;
    CrtLabel: TLabel;
    ServerNameInput: TEdit;
    KcpPwdInput: TEdit;
    KcpPwdLabel: TLabel;
    KeyInput: TEdit;
    KeyLabel: TLabel;
    Label15: TLabel;
    Label4: TLabel;
    PageControl1: TPageControl;
    ProxyTypeComboBox: TComboBox;
    SelectCrtAction: TAction;
    SelectCrtBtn: TButton;
    SelectKeyAction: TAction;
    SelectKeyBtn: TButton;
    SelectPrivateKeyAction: TAction;
    ActionList1: TActionList;
    ConfirmBtn: TButton;
    CancelBtn: TButton;
    GroupBox4: TGroupBox;
    ListView1: TListView;
    AddMapping: TMenuItem;
    ModifyMapping: TMenuItem;
    DelMapping: TMenuItem;
    SelectCrtDialog: TOpenDialog;
    SelectKeyDialog: TOpenDialog;
    SelectPrivateKeyBtn: TButton;
    SelectPrivateKeyDialog: TOpenDialog;
    PopupMenu1: TPopupMenu;
    LocalProtocolComboBox: TComboBox;
    DnsInput: TEdit;
    LocalPort: TEdit;
    Label13: TLabel;
    Label14: TLabel;
    LocalIp: TEdit;
    Label12: TLabel;
    SshKeyInput: TEdit;
    SshKeyLabel: TLabel;
    SshPwdInput: TEdit;
    SshPwdLabel: TLabel;
    SshUserNameInput: TEdit;
    SshUserNameLabel: TLabel;
    TlsPage: TTabSheet;
    SSHPage: TTabSheet;
    KcpPage: TTabSheet;
    TTLInput: TEdit;
    GroupBox2: TGroupBox;
    GroupBox3: TGroupBox;
    Label10: TLabel;
    Label11: TLabel;
    ServerProtocolComboBox: TComboBox;
    IpInput: TEdit;
    PortInput: TEdit;
    GroupBox1: TGroupBox;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
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

