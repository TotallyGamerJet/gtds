// +build windows

package gtds

import (
	"github.com/totallygamerjet/w32"
	"golang.org/x/sys/windows"
	"unsafe"
)

const (
	className = "GTDS_CLASSNAME"
)

var instance = w32.GetModuleHandle("GTDS")

func platformCreateWindow(w Window) {
	handle := w32.CreateWindowEx(
		0,
		windows.StringToUTF16Ptr(className),
		windows.StringToUTF16Ptr(w.Title),
		w32.WS_CAPTION|w32.WS_MINIMIZE|w32.WS_SYSMENU,
		0, 0, w.Width, w.Height,
		0, 0, instance, nil)
	w32.ShowWindow(handle, w32.SW_SHOW)
}

func platformRun() {
	wc := w32.WNDCLASSEX{}
	wc.Size = uint32(unsafe.Sizeof(wc))
	wc.Style = w32.CS_OWNDC
	wc.WndProc = windows.NewCallback(wndProc)
	wc.ClsExtra = 0
	wc.WndExtra = 0
	wc.Instance = instance
	wc.Icon = 0
	wc.Cursor = 0
	wc.Background = 0
	wc.MenuName = nil
	wc.ClassName = windows.StringToUTF16Ptr(className)
	wc.Icon = 0
	if atom := w32.RegisterClassEx(&wc); atom == 0 {
		panic("failed to register window class")
	}
	appUpdate() //try to catch a call to window. if this fails then the program freezes.
	//this is a race condition and should probably be fixed
	for !processMessages() {
	}
}

// processMessages return if the app should terimate
func processMessages() bool {
	msg := w32.MSG{}
	for w32.PeekMessage(&msg, 0, 0, 0, w32.PM_REMOVE) {
		if msg.Message == w32.WM_QUIT {
			return true
		}
		appUpdate()
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
	return false
}

func wndProc(hwnd w32.HWND, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_CLOSE:
		w32.PostQuitMessage(0) //TODO: check if there are other windows open
	}
	return w32.DefWindowProc(hwnd, msg, wparam, lparam)
}
