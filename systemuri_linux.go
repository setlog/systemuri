//go:build linux

package systemuri

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func registerURLHandler(name string, scheme string, applicationPath string, argumentsPattern string) error {
	if applicationPath == "" {
		return fmt.Errorf("applicationPath is empty")
	}

	exec := applicationPath + " " + strings.Replace(argumentsPattern, "%s", "%u", -1)

	desktopFileContent := fmt.Sprintf(`[Desktop Entry]
Version=1.0
Name=%s
Exec=%s
Terminal=false
Type=Application
NoDisplay=true
MimeType=x-scheme-handler/%s;
`, name, exec, scheme)

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

	desktopFilePath := filepath.Join(applicationsDir, fmt.Sprintf("%s-url-handler.desktop", scheme))
	err = os.WriteFile(desktopFilePath, []byte(desktopFileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to create .desktop file: %w", err)
	}
	return nil
}

func unregisterURLHandler(scheme string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	if xdgDataHome == "" {
		xdgDataHome = filepath.Join(usr.HomeDir, ".local", "share")
	}

	applicationsDir := filepath.Join(xdgDataHome, "applications")
	desktopFilePath := filepath.Join(applicationsDir, fmt.Sprintf("%s-url-handler.desktop", scheme))

	err = os.Remove(desktopFilePath)
	if err != nil {
		return fmt.Errorf("failed to remove .desktop file: %w", err)
	}

	return nil
}
