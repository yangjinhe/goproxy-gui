unit Unit1;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ComCtrls, Menus,
  ActnList, ExtCtrls, StdCtrls;

type

  { TMainForm }

  TMainForm = class(TForm)
    GroupBox1: TGroupBox;
    GroupBox2: TGroupBox;
    GroupBox3: TGroupBox;
    ImageList1: TImageList;
    ListView1: TListView;
    Memo1: TMemo;
    Memo2: TMemo;
    AddServer: TMenuItem;
    ShowMainWindow: TMenuItem;
    ExitMainWindow: TMenuItem;
    TrayPopupMenu: TPopupMenu;
    SetActivityServer: TMenuItem;
    PingServer: TMenuItem;
    PopupMenu1: TPopupMenu;
    ToolBar1: TToolBar;
    ToolButton1: TToolButton;
    ToolButton2: TToolButton;
    SettingsBtn: TToolButton;
    ToolButton4: TToolButton;
    TrayIcon1: TTrayIcon;
    UpdateBtn: TToolButton;
    ToolButton6: TToolButton;
    HelpBtn: TToolButton;
    ToolButton8: TToolButton;
    procedure AddServerClick(Sender: TObject);
    procedure ExitMainWindowClick(Sender: TObject);
    procedure PingServerClick(Sender: TObject);
    procedure SetActivityServerClick(Sender: TObject);
    procedure SettingsBtnClick(Sender: TObject);
    procedure HelpBtnClick(Sender: TObject);
    procedure ShowMainWindowClick(Sender: TObject);
    procedure UpdateBtnClick(Sender: TObject);
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }

procedure TMainForm.SettingsBtnClick(Sender: TObject);
begin

end;

procedure TMainForm.HelpBtnClick(Sender: TObject);
begin

end;

procedure TMainForm.ShowMainWindowClick(Sender: TObject);
begin

end;

procedure TMainForm.UpdateBtnClick(Sender: TObject);
begin

end;

procedure TMainForm.AddServerClick(Sender: TObject);
begin

end;

procedure TMainForm.ExitMainWindowClick(Sender: TObject);
begin

end;

procedure TMainForm.PingServerClick(Sender: TObject);
begin

end;

procedure TMainForm.SetActivityServerClick(Sender: TObject);
begin

end;

end.

