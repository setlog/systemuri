//go:build linux

package systemuri

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func registerURLHandler(name string, schema string, applicationPath string, noDisplay bool) error {
	if applicationPath == "" {
		return fmt.Errorf("applicationPath is empty")
	}

	noDisplayString := "false"
	if noDisplay {
		noDisplayString = "true"
	}
	desktopFileContent := fmt.Sprintf(`[Desktop Entry]
Version=1.0
Name=%s
Exec=%s %%u
Terminal=false
Type=Application
NoDisplay=%s
MimeType=x-scheme-handler/%s;
`, name, filepath.Clean(applicationPath), noDisplayString, schema)

	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	if xdgDataHome == "" {
		xdgDataHome = filepath.Join(usr.HomeDir, ".local", "share")
	}

	applicationsDir := filepath.Join(xdgDataHome, "applications")
	err = os.MkdirAll(applicationsDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create applications directory: %w", err)
	}

	desktopFilePath := filepath.Join(applicationsDir, fmt.Sprintf("%s-url-handler.desktop", schema))
	err = os.WriteFile(desktopFilePath, []byte(desktopFileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to create .desktop file: %w", err)
	}
	return nil
}

func unregisterURLHandler(applicationPath string) error {
	log.Println("Registering URL handler on linux is not implemented")
	return nil
}
