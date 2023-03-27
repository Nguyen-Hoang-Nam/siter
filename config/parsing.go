package config

import (
	"os"
	"path"
	"siter/env"

	"github.com/BurntSushi/toml"
)

const CONFIG_FILE_NAME = "siter.toml"

func parseFile() (c *Config, err error) {
	configDir := env.GetSiterConfigDirectory()
	configPath := path.Join(configDir, CONFIG_FILE_NAME)

	if _, err = os.Stat(configDir); os.IsNotExist(err) {
		return
	}

	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		return
	}

	configData, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	_, err = toml.Decode(string(configData), &c)
	if err != nil {
		return
	}

	return c, nil
}
