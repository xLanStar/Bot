package bot

import (
	"fmt"
	"unsafe"

	"github.com/xLanStar/bot/keys"
	"github.com/xLanStar/bot/winapi"
)

var (
	keybinds map[int16]Keybind

	idCounter int16
)

func Register(modifiers uint8, key keys.Key, function func(interface{}), description string) {
	idCounter++
	keybind := Keybind{Id: idCounter, Modifiers: modifiers, Key: key, Function: function, Description: description}
	keybinds[keybind.Id] = keybind
	result, _, err := winapi.RegisterHotkey.Call(0, uintptr(keybind.Id), uintptr(keybind.Modifiers), uintptr(keybind.Key))

	fmt.Println("[bot] Register:", keybind, result, err)
}

func Run() {
	// c := make(chan bool)
	for {
		var msg = &winapi.MSG{}
		winapi.PeekMessage.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0, 1)

		// Registered id is in the WPARAM field:
		if id := msg.WPARAM; id != 0 {
			fmt.Println("Hotkey pressed:", id, keybinds[id])
			if id == 1 { // CTRL+ALT+X = Exit
				fmt.Println("id 1 pressed, goodbye...")
				return
			}
		}
	}
}

func Stop(a interface{}) {
	fmt.Println("[bot] Stop")
}

func init() {
	fmt.Println("[bot] init")

	keybinds = map[int16]Keybind{}

	idCounter = 0

	Register(0, keys.KeyF10, Stop, "關閉")
}
