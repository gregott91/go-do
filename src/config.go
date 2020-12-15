package godo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const configFileName = "go-do.config.json"

// ShortcutConfig is the configuration for shortcuts
type ShortcutConfig struct {
	Switch ShortcutKey
	Delete ShortcutKey
	Close  ShortcutKey
}

// Config is the master config
type Config struct {
	Shortcuts ShortcutConfig
}

// ShortcutConfig is the configuration for shortcuts
type internalShortcutConfig struct {
	SwitchBetweenInputAndList string
	DeleteListItem            string
	CloseApplication          string
}

// Config is the master config
type internalConfig struct {
	Shortcuts internalShortcutConfig
}

// GetConfig gets the configuration object
func GetConfig() (Config, error) {
	var config internalConfig
	var err error

	configFilePath, err := ConcatenateFileWithCurrentExeDir(configFileName)

	if err != nil {
		return Config{}, err
	}

	if _, err = os.Stat(configFilePath); os.IsNotExist(err) {
		config = getDefaultConfig()
		err = createConfigFile(configFilePath, config)
	} else {
		dat, err := ioutil.ReadFile(configFilePath)

		if err != nil {
			return Config{}, err
		}

		if err := json.Unmarshal(dat, &config); err != nil {
			return Config{}, err
		}
	}

	return convertConfig(config), err
}

func convertConfig(internal internalConfig) Config {
	codeMap := GetConfigToCodeMap()

	return Config{
		Shortcuts: ShortcutConfig{
			Switch: codeMap[internal.Shortcuts.SwitchBetweenInputAndList],
			Delete: codeMap[internal.Shortcuts.DeleteListItem],
			Close:  codeMap[internal.Shortcuts.CloseApplication],
		},
	}
}

func createConfigFile(filePath string, config internalConfig) error {
	f, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer f.Close()

	d2, err := json.MarshalIndent(config, "", "    ")

	if err != nil {
		return err
	}

	_, err = f.Write(d2)

	return err
}

func getDefaultConfig() internalConfig {
	return internalConfig{
		Shortcuts: internalShortcutConfig{
			SwitchBetweenInputAndList: ConfigCtrlS,
			DeleteListItem:            ConfigCtrlD,
			CloseApplication:          ConfigEscape,
		},
	}
}
