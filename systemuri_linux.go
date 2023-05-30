//go:build linux

package systemuri

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func registerURLHandler(name string, schema string, applicationPath string) error {
	if applicationPath == "" {
		return fmt.Errorf("applicationPath is empty")
	}

	desktopFileContent := fmt.Sprintf(`[Desktop Entry]
Version=1.0
Name=%s
Exec=%s %%u
Terminal=false
Type=Application
NoDisplay=true
MimeType=x-scheme-handler/%s;
`, name, filepath.Clean(applicationPath), schema)

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

func unregisterURLHandler(schema string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	if xdgDataHome == "" {
		xdgDataHome = filepath.Join(usr.HomeDir, ".local", "share")
	}

	applicationsDir := filepath.Join(xdgDataHome, "applications")
	desktopFilePath := filepath.Join(applicationsDir, fmt.Sprintf("%s-url-handler.desktop", schema))

	err = os.Remove(desktopFilePath)
	if err != nil {
		return fmt.Errorf("failed to remove .desktop file: %w", err)
	}

	return nil
}
