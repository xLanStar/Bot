package bot

import (
	"bytes"
	"fmt"

	"github.com/xLanStar/bot/keys"
)

type Keybind struct {
	Id          int16
	Modifiers   uint8
	Key         keys.Key
	Function    func(interface{})
	Description string
}

// String returns a human-friendly display name of the hotkey
// such as "Hotkey[Id: 1, Alt+Ctrl+O]"
func (h *Keybind) String() string {
	mod := &bytes.Buffer{}
	if h.Modifiers&keys.ModAlt != 0 {
		mod.WriteString("Alt+")
	}
	if h.Modifiers&keys.ModCtrl != 0 {
		mod.WriteString("Ctrl+")
	}
	if h.Modifiers&keys.ModShift != 0 {
		mod.WriteString("Shift+")
	}
	if h.Modifiers&keys.ModWin != 0 {
		mod.WriteString("Win+")
	}
	return fmt.Sprintf("Keybind[Id: %d, %s%v]", h.Id, mod, h.Key)
}
