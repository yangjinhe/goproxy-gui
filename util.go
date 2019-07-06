package main

import (
	"fmt"
	"github.com/ProtonMail/go-autostart"
	"github.com/getlantern/sysproxy"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
