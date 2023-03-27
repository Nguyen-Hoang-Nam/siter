package env

import (
	"os"
	"path"
	"runtime"
)

func getXdgConfigHome() string {
	if value := os.Getenv(XDG_CONFIG_HOME); value != "" {
		return value
	}

	return "$HOME/.local/.config"
}

func getUserProfile() string {
	return os.Getenv(WINDOW_USER_PROFILE)
}

func GetSiterConfigDirectory() string {
	if value := os.Getenv(SITER_CONFIG_DIRECTORY); value != "" {
		return value
	}

	if runtime.GOOS == "windows" {
		return path.Join(getUserProfile(), "siter")
	} else {
		return path.Join(getXdgConfigHome(), "siter")
	}
}
