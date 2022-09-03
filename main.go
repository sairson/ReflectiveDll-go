//go:build windows && cgo

package main

import (
	"C"
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
	"strings"
	"time"
)

func main() {}

func run() {
	// 执行方法
	timeStr := strings.ReplaceAll(time.Now().Format("2006-01-02-15-04-05"), "-", "_")
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%s.png", timeStr)
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
}

// export functions

//export Run
// is designed to work with rundll32.exe to execute a function
func Run() {
	run()
}

//export DllInstall
// is used when executing the binary with regsvr32.exe (i.e. regsvr32.exe /s /n /i example.dll)
// https://msdn.microsoft.com/en-us/library/windows/desktop/bb759846(v=vs.85).aspx
func DllInstall() {
	run()
}

//export DllRegisterServer
// is used when executing the binary with regsvr32.exe (i.e. regsvr32.exe /s example.dll)
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms682162(v=vs.85).aspx
func DllRegisterServer() {
	run()
}

//export DllUnregisterServer
// is used when executing the binary with regsvr32.exe (i.e. regsvr32.exe /s /u example.dll)
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms691457(v=vs.85).aspx
func DllUnregisterServer() {
	run()
}

//export MyFunction
// is a custom set of functions
func MyFunction() {
	run()
}
