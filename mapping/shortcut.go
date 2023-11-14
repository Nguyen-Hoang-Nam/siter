package mapping

import (
	"errors"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var modifierKeys = map[string]fyne.KeyModifier{
	"ctrl":  fyne.KeyModifierControl,
	"alt":   fyne.KeyModifierAlt,
	"shift": fyne.KeyModifierShift,
}

var specialKeys = map[string]fyne.KeyName{
	"ESC":       fyne.KeyEscape,
	"RETURN":    fyne.KeyReturn,
	"TAB":       fyne.KeyTab,
	"BACKSPACE": fyne.KeyBackspace,
	"INSERT":    fyne.KeyInsert,
	"DELETE":    fyne.KeyDelete,
	"RIGHT":     fyne.KeyRight,
	"LEFT":      fyne.KeyLeft,
	"DOWN":      fyne.KeyDown,
	"UP":        fyne.KeyUp,
	"PAGE_UP":   fyne.KeyPageUp,
	"PAGE_DOWN": fyne.KeyPageDown,
	"HOME":      fyne.KeyHome,
	"END":       fyne.KeyEnd,
	"ENTER":     fyne.KeyEnter,
	"SPACE":     fyne.KeySpace,
}

var overrideShortcut = map[string]fyne.Shortcut{
	"ctrl+c": &fyne.ShortcutCopy{},
	"ctrl+v": &fyne.ShortcutPaste{},
	"ctrl+x": &fyne.ShortcutCut{},
	"ctrl+a": &fyne.ShortcutSelectAll{},
}

func (m *mappingStruct) functions() map[string]func([]string) func(fyne.Shortcut) {
	return map[string]func([]string) func(fyne.Shortcut){
		"write": func(args []string) func(fyne.Shortcut) {
			return func(_ fyne.Shortcut) {
				m.process.Write([]byte(args[0]))
			}
		},
	}
}

func parseShortcut(text string) (fyne.Shortcut, error) {
	if val, ok := overrideShortcut[text]; ok {
		return val, nil
	}

	keys := strings.Split(text, "+")
	if len(keys) != 2 {
		return nil, errors.New("shortcut much be 2 keys")
	}

	modifierKeyStr := strings.ToLower(keys[0])
	var modifierKey fyne.KeyModifier
	if val, ok := modifierKeys[modifierKeyStr]; ok {
		modifierKey = val
	} else {
		return nil, errors.New("modifier key not found")
	}

	keyStr := strings.ToUpper(keys[1])
	var key fyne.KeyName
	if val, ok := specialKeys[keyStr]; ok {
		key = val
	} else if len(keyStr) == 1 {
		key = fyne.KeyName(keyStr)
	} else {
		return nil, errors.New("key not found")
	}

	return &desktop.CustomShortcut{KeyName: key, Modifier: modifierKey}, nil
}

func (m *mappingStruct) loadShortcut() {
	functionMap := m.functions()

	for shortcutStr, functionStr := range m.conf.Map {
		shortcut, err := parseShortcut(shortcutStr)
		if err != nil {
			continue
		}

		params := strings.Split(functionStr, " ")
		functionName := params[0]
		args := params[1:]
		if val, ok := functionMap[functionName]; ok {
			m.canvas.AddShortcut(shortcut, val(args))
		}
	}
}
