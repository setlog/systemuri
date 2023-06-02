//go:build windows

package systemuri

import (
	"fmt"
	winreg "golang.org/x/sys/windows/registry"
	"io"
	"log"
	"strings"
)

const (
	windowsPlaceholder = "%1"
	KeySoftware        = "Software"
	KeyClasses         = "Classes"
	KeyShell           = "shell"
	KeyOpen            = "open"
	KeyCommand         = "command"
)

type CreateKeyFn func(k winreg.Key, path string, access uint32) (winreg.Key, bool, error)

func registerURLHandler(name string, scheme string, applicationPath string, argumentsPattern string) error {
	return registerURLHandlerWithRegistry(winreg.CreateKey, name, scheme, applicationPath, argumentsPattern)
}

func registerURLHandlerWithRegistry(createKey CreateKeyFn, name string, scheme string, applicationPath string, argumentsPattern string) error {

	log.Printf("Registering URL handler name:%s, scheme:%s, applicationPath:%s, argumentsPattern:%s on Windows\n", name, scheme, applicationPath, argumentsPattern)

	k, _, err := createKey(winreg.CURRENT_USER, strings.Join([]string{KeySoftware, KeyClasses, scheme}, "\\"), winreg.WRITE)
	if err != nil {
		return fmt.Errorf("failed to create key: %w", err)
	}
	defer closeSafely(k)

	err = k.SetStringValue("", "URL:"+name+" Protocol")
	if err != nil {
		return fmt.Errorf("failed to set default value: %w", err)
	}

	err = k.SetStringValue("URL Protocol", "")
	if err != nil {
		return fmt.Errorf("failed to set URL Protocol value: %w", err)
	}

	shellKey, _, err := createKey(k, strings.Join([]string{KeyShell, KeyOpen, KeyCommand}, "\\"), winreg.WRITE)
	if err != nil {
		return fmt.Errorf("failed to create shell\\open\\command key: %w", err)
	}
	defer closeSafely(shellKey)

	if argumentsPattern == "" {
		argumentsPattern = windowsPlaceholder
	} else {
		argumentsPattern = strings.ReplaceAll(argumentsPattern, "%s", windowsPlaceholder)
	}
	command := fmt.Sprintf(`"%s" "%s"`, applicationPath, argumentsPattern)
	err = shellKey.SetStringValue("", command)
	if err != nil {
		return fmt.Errorf("failed to set command value: %w", err)
	}
	return nil
}

func unregisterURLHandler(scheme string) error {
	registryPathsToDelete := []string{
		strings.Join([]string{KeySoftware, KeyClasses, scheme, KeyShell, KeyOpen, KeyCommand}, "\\"),
		strings.Join([]string{KeySoftware, KeyClasses, scheme, KeyShell, KeyOpen}, "\\"),
		strings.Join([]string{KeySoftware, KeyClasses, scheme, KeyShell}, "\\"),
		strings.Join([]string{KeySoftware, KeyClasses, scheme}, "\\"),
	}
	for _, registryPath := range registryPathsToDelete {
		if err := winreg.DeleteKey(winreg.CURRENT_USER, registryPath); err != nil {
			return fmt.Errorf("failed to remove registry key HKEY_CURRENT_USER\\%v: %w", registryPath, err)
		}
	}
	return nil
}

func unregisterURLHandlerByPath(applicationPath string) error {
	log.Println("Not yet implemented")
	return nil
}

func closeSafely(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Printf("error closing resource: %v\n", err)
	}
}
