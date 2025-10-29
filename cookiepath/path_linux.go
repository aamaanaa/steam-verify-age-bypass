//go:build linux

package cookiepath

import (
	"errors"
	"os"
	"path/filepath"
)

var cookiesPath = filepath.Join(".steam", "steam", "config", "htmlcache", "Cookies")

func GetSteamCookiePath() (path string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	cookiePath := filepath.Join(home, cookiesPath)

	if _, err = os.Stat(cookiePath); os.IsNotExist(err) {
		return "", errors.New("steam cookies file not found at: Make sure Steam has been opened at least once")
	}

	return cookiePath, nil
}
