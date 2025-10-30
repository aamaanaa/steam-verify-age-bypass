package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

var (
	cookies = []interface{}{
		[]interface{}{
			int64(13370000000000000), "store.steampowered.com", "", "wants_mature_content", "1",
			[]byte{}, "/", int64(20000000000000000), 0, 0, int64(13370000000000000), 1, 1, 1, -1, 2, 443, int64(13370000000000000), 1, 1,
		},
		[]interface{}{
			int64(13370000000000000), "store.steampowered.com", "", "lastagecheckage", "20-August-1995",
			[]byte{}, "/", int64(20000000000000000), 0, 0, int64(13370000000000000), 1, 1, 1, -1, 2, 443, int64(13370000000000000), 1, 1,
		},
		[]interface{}{
			int64(13370000000000000), "store.steampowered.com", "", "birthtime", "788914801",
			[]byte{}, "/", int64(20000000000000000), 0, 0, int64(13370000000000000), 1, 1, 1, -1, 2, 443, int64(13370000000000000), 1, 1,
		},
	}

	stmt = `
	INSERT OR REPLACE INTO cookies (
		creation_utc, host_key, top_frame_site_key, name, value,
		encrypted_value, path, expires_utc, is_secure, is_httponly,
		last_access_utc, has_expires, is_persistent, priority, samesite,
		source_scheme, source_port, last_update_utc, source_type,
		has_cross_site_ancestor
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
)

func insertCookies(dbPath string) (err error) {

	fs, err := os.Stat(dbPath)
	if err != nil {
		return err
	}

	if fs.IsDir() {
		return errors.New("path points to a directory and not a file")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	for _, c := range cookies {
		cookie, ok := c.([]interface{})
		if !ok {
			return errors.New("invalid cookie format")
		}

		_, err := db.Exec(stmt, cookie...)
		if err != nil {
			return fmt.Errorf("insert cookie failed: %v", err)
		}
	}

	return nil
}

func main() {

	err := Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	select {}
}

func Run() (err error) {
	var (
		cookiesPathLinux   = filepath.Join(".steam", "steam", "config", "htmlcache", "Cookies")
		cookiesPathWindows = filepath.Join("Steam", "config", "htmlcache", "Cookies")
		path               string
	)

	filePath := flag.String("file", "", "Custom path to the Steam cookies file")
	flag.Parse()

	fmt.Println("implementing bypass...")

	if *filePath != "" {
		path = *filePath
	} else {
		switch runtime.GOOS {

		case "windows":

			configDir, err := os.UserConfigDir()
			if err != nil {
				return err
			}

			path = filepath.Join(configDir, cookiesPathWindows)
			break

		case "darwin", "linux":

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			path = filepath.Join(home, cookiesPathLinux)
			break

		default:
			fmt.Println("Unsupported OS:", runtime.GOOS)
			return
		}
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Failure: ", err)
		return
	}

	if _, err = os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf(`cookie file %s not found, make sure this is the correct path`, absPath)
	}

	fmt.Println("Found Cookies file:\n>", absPath)

	if err = insertCookies(absPath); err != nil {
		fmt.Println("Failure: ", err)
		return
	}

	fmt.Println("Bypass completed successfully!\nYou can now open Steam and access the store without age verification.\nYou may close the program.")

	return nil
}
