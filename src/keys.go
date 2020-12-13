package godo

// ShortcutKey is the key value for a shortcut
type ShortcutKey struct {
	ConfigValue  string
	DisplayValue string
	Code         uint64
}

const (
	// ConfigCtrlS is the config value for CtrlS
	ConfigCtrlS = "CtrlS"
	// ConfigCtrlD is the config value for CtrlD
	ConfigCtrlD = "CtrlD"
	// ConfigEscape is the config value for Escape
	ConfigEscape = "Escape"
)

// GetConfigToCodeMap returns a map from key config values to key codes
func GetConfigToCodeMap() map[string]ShortcutKey {
	keyMap := make(map[string]ShortcutKey)
	keys := getAllShortcutKeys()

	for _, key := range keys {
		keyMap[key.ConfigValue] = key
	}

	return keyMap
}

func getAllShortcutKeys() []ShortcutKey {
	newKey := func(value string, code uint64) ShortcutKey {
		return ShortcutKey{ConfigValue: value, DisplayValue: value, Code: code}
	}

	return []ShortcutKey{
		newKey("CtrlA", 1),
		newKey("CtrlB", 2),
		newKey("CtrlC", 3),
		newKey(ConfigCtrlD, 4),
		newKey("CtrlE", 5),
		newKey("CtrlF", 6),
		newKey("CtrlG", 7),
		newKey("CtrlH", 8),
		newKey("CtrlI", 9),
		newKey("CtrlJ", 10),
		newKey("CtrlK", 11),
		newKey("CtrlL", 12),
		newKey("CtrlM", 13),
		newKey("CtrlN", 14),
		newKey("CtrlO", 15),
		newKey("CtrlP", 16),
		newKey("CtrlQ", 17),
		newKey("CtrlR", 18),
		newKey(ConfigCtrlS, 19),
		newKey("CtrlT", 20),
		newKey("CtrlU", 21),
		newKey("CtrlV", 22),
		newKey("CtrlW", 23),
		newKey("CtrlX", 24),
		newKey("CtrlY", 25),
		newKey("CtrlZ", 26),
		newKey(ConfigEscape, 27),
		newKey("CtrlBackslash", 28),
		newKey("CtrlRightSq", 29),
		newKey("CtrlCarat", 30),
		newKey("CtrlUnderscore", 31),
	}
}
