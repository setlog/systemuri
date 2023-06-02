package systemuri

import (
	"fmt"
	"regexp"
	"strings"
)

var validateScheme = regexp.MustCompile("^[a-zA-Z0-9-]*$")

// RegisterURLHandlerPassingAll registers a custom URL handler for the specified schema and application
func RegisterURLHandlerPassingAll(name string, scheme string, applicationPath string) error {
	return registerURLHandler(name, scheme, applicationPath, "%s")
}

// RegisterURLHandler registers a custom URL handler for the specified schema and application. argumentsPattern the %s is replaced by the system-specific WINDOWS_PLACEHOLDER
func RegisterURLHandler(name string, scheme string, applicationPath string, argumentsPattern string) error {
	if err := validateURLHandleParameters(name, scheme, applicationPath, argumentsPattern); err != nil {
		return err
	}
	return registerURLHandler(name, scheme, applicationPath, argumentsPattern)
}

// UnregisterURLHandler removes all registered entries based on the scheme
func UnregisterURLHandler(scheme string) error {
	return unregisterURLHandler(scheme)
}

// UnregisterURLHandlerByPath removes all registered entries based on the application path
func UnregisterURLHandlerByPath(applicationPath string) error {
	return unregisterURLHandlerByPath(applicationPath)
}

func validateURLHandleParameters(name string, scheme string, applicationPath string, argumentsPattern string) error {
	if name == "" {
		return fmt.Errorf("`name` is required but empty")
	}
	if scheme == "" {
		return fmt.Errorf("`scheme` is required but empty")
	}
	if !validateScheme.MatchString(scheme) {
		return fmt.Errorf("`scheme` must match the regex %v", validateScheme.String())
	}
	if applicationPath == "" {
		return fmt.Errorf("`applicationPath` is required but empty")
	}
	if argumentsPattern != "" && !strings.Contains(argumentsPattern, "%s") {
		return fmt.Errorf("`argumentsPattern` has to contain %%s to be replaced by os-specific placeholder")
	}
	return nil
}
