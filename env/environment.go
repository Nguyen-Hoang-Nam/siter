package env

import (
	"os"
	"path"
)

func getXdgConfigHome() string {
	if value := os.Getenv(XDG_CONFIG_HOME); value != "" {
		return value
	}

	return "$HOME/.local/.config"
}

func GetSiterConfigDirectory() string {
	if value := os.Getenv(SITER_CONFIG_DIRECTORY); value != "" {
		return value
	}

	return path.Join(getXdgConfigHome(), "siter")
}
