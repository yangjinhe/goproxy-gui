unit Unit3;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls;

type

  { TAddMappingForm }

  TAddMappingForm = class(TForm)
    ConfirmBtn: TButton;
    CancelBtn: TButton;
    TypeComboBox: TComboBox;
    RemotePort: TEdit;
    Label3: TLabel;
    LocalPort: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    procedure CancelBtnClick(Sender: TObject);
    procedure TypeComboBoxChange(Sender: TObject);
    procedure ConfirmBtnClick(Sender: TObject);
  private

  public

  end;

var
  AddMappingForm: TAddMappingForm;

implementation

{$R *.lfm}

{ TAddMappingForm }

procedure TAddMappingForm.ConfirmBtnClick(Sender: TObject);
begin

end;

procedure TAddMappingForm.CancelBtnClick(Sender: TObject);
begin

end;

procedure TAddMappingForm.TypeComboBoxChange(Sender: TObject);
begin

end;

end.

