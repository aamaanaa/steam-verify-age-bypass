package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"steamagebypass/cookiepath"
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

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	for _, c := range cookies {
		cookie, ok := c.([]interface{})
		if !ok {
			fmt.Println("Error: Invalid cookie format")
			continue
		}

		_, err := db.Exec(stmt, cookie...)
		if err != nil {
			fmt.Println("Error: Failed to insert cookie:", err)
			continue
		}
	}

	return nil
}

func main() {

	fmt.Println("implementing bypass...")

	dbPath, err := cookiepath.GetSteamCookiePath()
	if err != nil {
		fmt.Println("Failure: ", err)
		return
	}


	if err = insertCookies(dbPath); err != nil {
		fmt.Println("Failure: ", err)
		return
	}

	fmt.Println("Bypass completed successfully! You can now open Steam and access the store without age verification.\nYou may close the program.")

	select {}
}
