package config

import (
	"os"
	"path"
	"siter/env"

	"github.com/BurntSushi/toml"
)

const CONFIG_FILE_NAME = "config.toml"

func userConfig(c *Config) (err error) {
	configDir, err := env.ConfigDir()
	if err != nil {
		return
	}

	configPath := path.Join(configDir, CONFIG_FILE_NAME)

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

	return nil
}
