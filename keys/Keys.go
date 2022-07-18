package keys

const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
)

type Key uint16

// All kinds of Keys
const (
	KeySpace Key = 0x20
	Key0     Key = 0x30
	Key1     Key = 0x31
	Key2     Key = 0x32
	Key3     Key = 0x33
	Key4     Key = 0x34
	Key5     Key = 0x35
	Key6     Key = 0x36
	Key7     Key = 0x37
	Key8     Key = 0x38
	Key9     Key = 0x39
	KeyA     Key = 0x41
	KeyB     Key = 0x42
	KeyC     Key = 0x43
	KeyD     Key = 0x44
	KeyE     Key = 0x45
	KeyF     Key = 0x46
	KeyG     Key = 0x47
	KeyH     Key = 0x48
	KeyI     Key = 0x49
	KeyJ     Key = 0x4A
	KeyK     Key = 0x4B
	KeyL     Key = 0x4C
	KeyM     Key = 0x4D
	KeyN     Key = 0x4E
	KeyO     Key = 0x4F
	KeyP     Key = 0x50
	KeyQ     Key = 0x51
	KeyR     Key = 0x52
	KeyS     Key = 0x53
	KeyT     Key = 0x54
	KeyU     Key = 0x55
	KeyV     Key = 0x56
	KeyW     Key = 0x57
	KeyX     Key = 0x58
	KeyY     Key = 0x59
	KeyZ     Key = 0x5A

	KeyReturn Key = 0x0D
	KeyEscape Key = 0x1B
	KeyDelete Key = 0x2E
	KeyTab    Key = 0x09

	KeyLeft  Key = 0x25
	KeyRight Key = 0x27
	KeyUp    Key = 0x26
	KeyDown  Key = 0x28

	KeyF1  Key = 0x70
	KeyF2  Key = 0x71
	KeyF3  Key = 0x72
	KeyF4  Key = 0x73
	KeyF5  Key = 0x74
	KeyF6  Key = 0x75
	KeyF7  Key = 0x76
	KeyF8  Key = 0x77
	KeyF9  Key = 0x78
	KeyF10 Key = 0x79
	KeyF11 Key = 0x7A
	KeyF12 Key = 0x7B
	KeyF13 Key = 0x7C
	KeyF14 Key = 0x7D
	KeyF15 Key = 0x7E
	KeyF16 Key = 0x7F
	KeyF17 Key = 0x80
	KeyF18 Key = 0x81
	KeyF19 Key = 0x82
	KeyF20 Key = 0x83
)
