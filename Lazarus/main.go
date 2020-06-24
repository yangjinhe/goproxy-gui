// 由res2go自动生成。
package main

import (
    "github.com/ying32/govcl/vcl"
)

func main() {
    vcl.Application.Initialize()
    vcl.Application.CreateForm(mainFormBytes, &MainForm)
    vcl.Application.CreateForm(newProxyServerFormBytes, &NewProxyServerForm)
    vcl.Application.CreateForm(addMappingFormBytes, &AddMappingForm)
    initDb()
    initListView()
    vcl.Application.Run()
}
