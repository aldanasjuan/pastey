package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/atotto/clipboard"
	"github.com/micmonay/keybd_event"
)

const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
	KEY_EXIT   = -1
	KEY_RELOAD = -2
)

type Hotkey struct {
	Id        int // Unique id
	Modifiers int // Mask of modifiers
	KeyCode   int // Key code, e.g. 'A'
	Value     []byte
	Text      string
}

// String returns a human-friendly display name of the hotkey
// such as "Hotkey[Id: 1, Alt+Ctrl+O]"
func (h *Hotkey) String() string {
	mod := &bytes.Buffer{}
	if h.Modifiers&ModAlt != 0 {
		mod.WriteString("Alt+")
	}
	if h.Modifiers&ModCtrl != 0 {
		mod.WriteString("Ctrl+")
	}
	if h.Modifiers&ModShift != 0 {
		mod.WriteString("Shift+")
	}
	if h.Modifiers&ModWin != 0 {
		mod.WriteString("Win+")
	}
	return fmt.Sprintf("Hotkey[Id: %d, %s%c]", h.Id, mod, h.KeyCode)
}

func main() {
	runtime.LockOSThread()
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Println(err)
		return
	}
	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	path := "./keys.config"
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
		if _, err := os.Stat(os.Args[1]); err == nil {
			path = os.Args[1]
		} else {
			_, err := os.Stat(path)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	keys, err := ParseFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := make(chan int, 2)
	go func() {
		for s := range c {
			for _, key := range keys {
				if key.Id == s {
					if key.Id == KEY_RELOAD {

					} else {
						clipboard.WriteAll(string(key.Value))
						kb.Clear()
						time.Sleep(150 * time.Millisecond)
						kb.AddKey(keybd_event.VK_V)
						kb.HasCTRL(true)
						err = kb.Press()
						if err != nil {
							fmt.Println(err)
							continue
						}
						time.Sleep(10 * time.Millisecond)
						err = kb.Release()
						if err != nil {
							fmt.Println(err)
							continue
						}
					}
					// clipboard.Write(clipboard.FmtText, key.Value)
				}
			}

		}
	}()
	user32 := syscall.MustLoadDLL("user32")
	defer user32.Release()

	reghotkey := user32.MustFindProc("RegisterHotKey")

	// Hotkeys to listen to:
	// keys := map[int16]*Hotkey{
	// 	1: {1, ModAlt, '1', 0},                        // ALT+CTRL+O
	// 	2: {2, ModAlt, '2', 0},                        // ALT+SHIFT+M
	// 	3: {3, ModCtrl + ModAlt, '3', 0},              // ALT+CTRL+X
	// 	4: {-1, ModAlt, codes.GetKey("backspace"), 0}, // ALT+CTRL+X
	// }

	// Register hotkeys:
	for _, v := range keys {
		r1, _, err := reghotkey.Call(
			0, uintptr(v.Id), uintptr(v.Modifiers), uintptr(v.KeyCode))
		if r1 == 1 {
			if v.Id == KEY_EXIT {
				fmt.Printf("Registered %v (exit program)\n", v.Text)
			} else if v.Id == KEY_RELOAD {
				fmt.Printf("Registered %v (reload config file)\n", v.Text)
			} else {
				fmt.Printf("Registered %v with %v\n", v.Text, string(v.Value))
			}
		} else {
			fmt.Println("Failed to register", v.Text, ", error:", err)
		}
	}

	getmsg := user32.MustFindProc("GetMessageW")

	for {
		var msg = &MSG{}
		getmsg.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0)

		// Registered id is in the WPARAM field:
		if id := msg.WPARAM; id != 0 {
			if id == -1 { // CTRL+ALT+X = Exit
				fmt.Println("Closing...")
				return
			}
			c <- int(id)
		}
	}
}

type MSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}
