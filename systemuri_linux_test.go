//go:build linux

package systemuri

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterURLHandler(t *testing.T) {
	testCases := []struct {
		name            string
		schema          string
		applicationPath string
		wantErr         bool
	}{
		{
			name:            "Valid case",
			schema:          "test",
			applicationPath: "/usr/bin/testapp",
			wantErr:         false,
		},
		{
			name:            "Invalid application path",
			schema:          "test",
			applicationPath: "",
			wantErr:         true,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			err := registerURLHandler(testcase.name, testcase.schema, testcase.applicationPath)
			if testcase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				usr, _ := user.Current()
				xdgDataHome := os.Getenv("XDG_DATA_HOME")
				if xdgDataHome == "" {
					xdgDataHome = filepath.Join(usr.HomeDir, ".local", "share")
				}
				desktopFilePath := filepath.Join(xdgDataHome, "applications", fmt.Sprintf("%s-url-handler.desktop", testcase.schema))
				_, err := os.Stat(desktopFilePath)
				assert.False(t, errors.Is(err, os.ErrNotExist))
			}
		})
	}
}
