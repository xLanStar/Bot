package winapi

import (
	"syscall"
)

var (
	user32           = syscall.MustLoadDLL("user32")
	RegisterHotkey   = user32.MustFindProc("RegisterHotKey")
	UnregisterHotkey = user32.MustFindProc("UnregisterHotKey")
	GetMessage       = user32.MustFindProc("GetMessageW")
	PeekMessage      = user32.MustFindProc("PeekMessageA")
	SendMessage      = user32.MustFindProc("SendMessageW")
	GetAsyncKeyState = user32.MustFindProc("GetAsyncKeyState")
	QuitMessage      = user32.MustFindProc("PostQuitMessage")
)

// func RegisterHotKey(hwnd, id uintptr, mod uintptr, k uintptr) (bool, error) {
// 	ret, _, err := registerHotkey.Call(
// 		hwnd, id, mod, k,
// 	)
// 	return ret != 0, err
// }

// func UnregisterHotKey(hwnd, id uintptr) (bool, error) {
// 	ret, _, err := unregisterHotkey.Call(hwnd, id)
// 	return ret != 0, err
// }

type MSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}

// func SendMessage(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
// 	ret, _, _ := sendMessage.Call(
// 		hwnd,
// 		uintptr(msg),
// 		wParam,
// 		lParam,
// 	)

// 	return ret
// }

// func GetMessage(msg *MSG, hWnd uintptr, msgFilterMin, msgFilterMax uint32) bool {
// 	ret, _, _ := getMessage.Call(
// 		uintptr(unsafe.Pointer(msg)),
// 		hWnd,
// 		uintptr(msgFilterMin),
// 		uintptr(msgFilterMax),
// 	)

// 	return ret != 0
// }

// func PeekMessage(msg *MSG, hWnd uintptr, msgFilterMin, msgFilterMax uint32) bool {
// 	ret, _, _ := peekMessage.Call(
// 		uintptr(unsafe.Pointer(msg)),
// 		hWnd,
// 		uintptr(msgFilterMin),
// 		uintptr(msgFilterMax),
// 		0, // PM_NOREMOVE
// 	)

// 	return ret != 0
// }

// func PostQuitMessage(exitCode int) {
// 	quitMessage.Call(uintptr(exitCode))
// }

// func GetAsyncKeyState(keycode int) uintptr {
// 	ret, _, _ := getAsyncKeyState.Call(uintptr(keycode))
// 	return ret
// }
