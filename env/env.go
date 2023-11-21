package env

import (
	"os"
	"path"
)

const CONFIG_DIRECTORY = "SITER_CONFIG_DIRECTORY"

func ConfigDir() (string, error) {
	if value := os.Getenv(CONFIG_DIRECTORY); value != "" {
		return value, nil
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(configDir, "siter"), nil
}
