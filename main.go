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

type Cookie struct {
	CreationUTC          int64
	HostKey              string
	TopFrameSiteKey      string
	Name                 string
	Value                string
	EncryptedValue       []byte
	Path                 string
	ExpiresUTC           int64
	IsSecure             int
	IsHTTPOnly           int
	LastAccessUTC        int64
	HasExpires           int
	IsPersistent         int
	Priority             int
	SameSite             int
	SourceScheme         int
	SourcePort           int
	LastUpdateUTC        int64
	SourceType           int
	HasCrossSiteAncestor int
}

func (c Cookie) Values() []interface{} {
	return []interface{}{
		c.CreationUTC,
		c.HostKey,
		c.TopFrameSiteKey,
		c.Name,
		c.Value,
		c.EncryptedValue,
		c.Path,
		c.ExpiresUTC,
		c.IsSecure,
		c.IsHTTPOnly,
		c.LastAccessUTC,
		c.HasExpires,
		c.IsPersistent,
		c.Priority,
		c.SameSite,
		c.SourceScheme,
		c.SourcePort,
		c.LastUpdateUTC,
		c.SourceType,
		c.HasCrossSiteAncestor,
	}
}

var cookies = []Cookie{
	{13370000000000000, "store.steampowered.com", "", "wants_mature_content", "1", []byte{}, "/", 20000000000000000, 0, 0, 13370000000000000, 1, 1, 1, -1, 2, 443, 13370000000000000, 1, 1},
	{13370000000000000, "store.steampowered.com", "", "lastagecheckage", "20-August-1995", []byte{}, "/", 20000000000000000, 0, 0, 13370000000000000, 1, 1, 1, -1, 2, 443, 13370000000000000, 1, 1},
	{13370000000000000, "store.steampowered.com", "", "birthtime", "788914801", []byte{}, "/", 20000000000000000, 0, 0, 13370000000000000, 1, 1, 1, -1, 2, 443, 13370000000000000, 1, 1},
}

const insertStmt = `
INSERT OR REPLACE INTO cookies (
	creation_utc, host_key, top_frame_site_key, name, value,
	encrypted_value, path, expires_utc, is_secure, is_httponly,
	last_access_utc, has_expires, is_persistent, priority, samesite,
	source_scheme, source_port, last_update_utc, source_type,
	has_cross_site_ancestor
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
`

func getCookiesPath(customPath string) (string, error) {
	if customPath != "" {
		return customPath, nil
	}

	switch runtime.GOOS {
	case "windows":
		configDir, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(configDir, "Steam", "config", "htmlcache", "Cookies"), nil

	case "darwin", "linux":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".steam", "steam", "config", "htmlcache", "Cookies"), nil

	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}

func insertCookies(dbPath string, dryRun bool) error {
	if fi, err := os.Stat(dbPath); err != nil {
		return fmt.Errorf("cannot access database file: %w", err)
	} else if fi.IsDir() {
		return errors.New("path points to a directory, not a file")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}

	stmt, err := db.Prepare(insertStmt)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, cookie := range cookies {
		if dryRun {
			fmt.Printf("DRY-RUN: would insert cookie '%s=%s' into '%s'\n", cookie.Name, cookie.Value, dbPath)
			continue
		}

		if _, err := stmt.Exec(cookie.Values()...); err != nil {
			return fmt.Errorf("failed to insert cookie %q: %w", cookie.Name, err)
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("\nPress Enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}

func run() error {
	filePath := flag.String("file", "", "Custom path to the Steam cookies file")
	dryRun := flag.Bool("dry-run", false, "Simulate the changes without modifying the DB")
	flag.Parse()

	fmt.Println("Implementing bypass...")

	path, err := getCookiesPath(*filePath)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	if _, err := os.Stat(absPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("cookie file '%s' not found", absPath)
		}
		return fmt.Errorf("cannot access cookie file: %w", err)
	}

	fmt.Println("Found Cookies file:\n>", absPath)

	if err := insertCookies(absPath, *dryRun); err != nil {
		return err
	}

	if *dryRun {
		fmt.Println("Dry run completed — no changes were made to the database.")
	} else {
		fmt.Println("Bypass completed successfully!")
		fmt.Println("You can now open Steam and access the store without age verification.")
	}

	return nil
}
