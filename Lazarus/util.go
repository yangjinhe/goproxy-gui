package main

import (
	"database/sql"
	"fmt"
	"github.com/ProtonMail/go-autostart"
	"github.com/getlantern/sysproxy"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var mainDbFileName = "goproxy.db"

var proxyInfoSql = `CREATE TABLE "proxy_info" (
		  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		  "name" TEXT(128),
		  "proxyType" TEXT(10),
		  "remoteIp" TEXT(20),
		  "remotePort" TEXT(6),
		  "remoteProtocol" TEXT(4),
		  "localIp" TEXT(20),
		  "localPort" TEXT(6),
		  "localProtocol" TEXT(4),
		  "crtLocation" TEXT(128),
		  "keyLocation" TEXT(128),
		  "sshUserName" TEXT(128),
		  "sshPwd" TEXT(128),
		  "sshPrivateKey" TEXT(128),
		  "kcpPwd" TEXT(128),
		  "dns" TEXT(20),
		  "dnsTtl" TEXT(20)
		);`
var proxyInfoIdxSql = `CREATE UNIQUE INDEX "idx_proxy_info_id"
			ON "proxy_info" (
			  "id"
			);`

var tunnelMappingSql = `CREATE TABLE "tunnel_mapping" (
		  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		  "info_id" INTEGER,
		  "type" TEXT(10),
		  "localPort" TEXT(128),
		  "remotePort" TEXT(128)
		);`
var tunnelMappingIdxSql = `CREATE UNIQUE INDEX "idx_mapping_id"
		ON "tunnel_mapping" (
		  "id"
		);`

var insertProxyInfoSql = `insert into "proxy_info" ("name", "proxyType",
                          "remoteIp", "remotePort", "remoteProtocol", "localIp", "localPort", "localProtocol",
                          "crtLocation", "keyLocation", "sshUserName", "sshPwd", "sshPrivateKey", "kcpPwd",
                          "dns", "dnsTtl") values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

var insertTunnelMappingSql = `insert into "tunnel_mapping" ("info_id", "protocol_type", "localPort", "remotePort") values (?, ?, ?, ?)`

func initDb() {

	b, _ := PathExists(mainDbFileName)
	if !b {
		db, err := sql.Open("sqlite3", mainDbFileName)
		checkErr(err)
		_, err = db.Exec(proxyInfoSql)
		checkErr(err)
		_, err = db.Exec(proxyInfoIdxSql)
		checkErr(err)
		_, err = db.Exec(tunnelMappingSql)
		checkErr(err)
		_, err = db.Exec(tunnelMappingIdxSql)
		checkErr(err)
	}
}

type ProxyInfo struct {
	id             int64
	name           string
	proxyType      string
	remoteIp       string
	remotePort     string
	remoteProtocol string
	localIp        string
	localPort      string
	localProtocol  string
	crtLocation    string
	keyLocation    string
	sshUserName    string
	sshPwd         string
	sshPrivateKey  string
	kcpPwd         string
	dns            string
	dnsTtl         string
}

type TunnelMapping struct {
	id           int64
	infoId       int64
	protocolType string
	localPort    string
	remotePort   string
}

func initListView() {
	MainForm.ListView1.SetRowSelect(true)
	b, err := PathExists(mainDbFileName)
	if err != nil {
		panic("错误")
	}
	if !b {
		initDb()
	} else {
		db, err := sql.Open("sqlite3", mainDbFileName)
		checkErr(err)
		result, _ := db.Query("select * from proxy_info")
		MainForm.ListView1.Items().BeginUpdate()
		for result.Next() {
			p := new(ProxyInfo)
			err = result.Scan(&p.id, &p.name, &p.proxyType, &p.remoteIp, &p.remotePort,
				&p.remoteProtocol, &p.localIp, &p.localPort, &p.localProtocol, &p.crtLocation,
				&p.keyLocation, &p.sshUserName, &p.sshPwd, &p.sshPrivateKey, &p.kcpPwd, &p.dns, &p.dnsTtl)
			checkErr(err)
			listItem := MainForm.ListView1.Items().Add()
			listItem.SetCaption(string(p.id))
			listItem.SubItems().Add(p.name)
			listItem.SubItems().Add(p.remoteIp)
			listItem.SubItems().Add(p.remotePort)
			listItem.SubItems().Add(p.remoteProtocol)
			listItem.SubItems().Add(p.localProtocol)
			listItem.SubItems().Add(p.localPort)
			listItem.SubItems().Add("未启动")
		}
		MainForm.ListView1.Items().EndUpdate()
	}
}

func addNewProxyInfo(proxyInfo ProxyInfo) (id int64) {
	db, err := sql.Open("sqlite3", mainDbFileName)
	checkErr(err)
	stmt, err := db.Prepare(insertProxyInfoSql)
	checkErr(err)
	result, err := stmt.Exec(proxyInfo.name, proxyInfo.proxyType, proxyInfo.remoteIp, proxyInfo.remotePort, proxyInfo.remoteProtocol,
		proxyInfo.localIp, proxyInfo.localPort, proxyInfo.localProtocol, proxyInfo.crtLocation, proxyInfo.keyLocation,
		proxyInfo.sshUserName, proxyInfo.sshPwd, proxyInfo.sshPrivateKey, proxyInfo.kcpPwd, proxyInfo.dns,
		proxyInfo.dnsTtl)
	checkErr(err)
	id, _ = result.LastInsertId()
	return id
}

func addNewTunnelMapping(tunnelMapping TunnelMapping) (id int64) {
	db, err := sql.Open("sqlite3", mainDbFileName)
	checkErr(err)
	stmt, err := db.Prepare(insertTunnelMappingSql)
	checkErr(err)
	result, err := stmt.Exec(tunnelMapping.infoId, tunnelMapping.protocolType, tunnelMapping.localPort, tunnelMapping.remotePort)
	checkErr(err)
	id, _ = result.LastInsertId()
	return id
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func setAutoRun() bool {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(dir)
	split := strings.Split(os.Args[0], string(os.PathSeparator))
	exeName := ""
	if len(split) > 0 {
		fmt.Println(split[len(split)-1])
		exeName = split[len(split)-1]
	} else {
		exeName = os.Args[0]
	}
	fmt.Println(exeName)

	var exec = dir + string(os.PathSeparator) + exeName
	app := &autostart.App{
		Name:        "goproxy-gui",
		DisplayName: "goproxy的图形界面",
		Exec:        []string{exec},
	}

	if app.IsEnabled() {
		log.Println("App is already enabled, removing it...")
	} else {
		log.Println("Enabling app...")

		if err := app.Enable(); err != nil {
			log.Fatal(err)
			return false
		}
	}
	return true
}

func unSetAutoRun() bool {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(dir)
	split := strings.Split(os.Args[0], string(os.PathSeparator))
	exeName := ""
	if len(split) > 0 {
		fmt.Println(split[len(split)-1])
		exeName = split[len(split)-1]
	} else {
		exeName = os.Args[0]
	}
	fmt.Println(exeName)

	var exec = dir + string(os.PathSeparator) + exeName
	app := &autostart.App{
		Name:        "goproxy-gui",
		DisplayName: "goproxy的图形界面",
		Exec:        []string{exec},
	}

	if app.IsEnabled() {
		log.Println("App is already enabled, removing it...")

		if err := app.Disable(); err != nil {
			log.Fatal(err)
			return false
		}
	}
	return true
}

func setSysProxy(addr string) bool {
	helperFullPath := "sysproxy-cmd"
	iconFullPath, _ := filepath.Abs("./icon.png")
	log.Println("Using icon at %v", iconFullPath)
	err := sysproxy.EnsureHelperToolPresent(helperFullPath, "Input your password and save the world!", iconFullPath)
	if err != nil {
		fmt.Printf("Error EnsureHelperToolPresent: %s\n", err)
		return false
	}
	_, err1 := sysproxy.On(addr)
	if err1 != nil {
		fmt.Printf("Error set proxy: %s\n", err)
		return false
	}
	return true
}

func unsetSysProxy(addr string) bool {
	helperFullPath := "sysproxy-cmd"
	iconFullPath, _ := filepath.Abs("./icon.png")
	//log.Println("Using icon at %v", iconFullPath)
	err := sysproxy.EnsureHelperToolPresent(helperFullPath, "Input your password and save the world!", iconFullPath)
	if err != nil {
		fmt.Printf("Error EnsureHelperToolPresent: %s\n", err)
		return false
	}
	off, err1 := sysproxy.On(addr)
	if err1 != nil {
		fmt.Printf("Error set proxy: %s\n", err1)
		return false
	}
	err2 := off()
	if err2 != nil {
		fmt.Printf("Error unset proxy: %s\n", err2)
		return false
	}
	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
