//go:build windows

package cookiepath

import (
	"errors"
	"os"
	"path/filepath"
)

func GetSteamCookiePath() (path string, err error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	cookiePath := filepath.Join(configDir, "Steam", "htmlcache", "Cookies")

	if _, err = os.Stat(cookiePath); os.IsNotExist(err) {
		return "", errors.New("steam cookies file not found: Make sure Steam has been opened at least once")
	}
	return cookiePath, nil
}
