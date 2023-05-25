package systemuri

// RegisterURLHandler registers a custom URL handler for the specified schema and application.
// noDisplay is only supported by Linux
func RegisterURLHandler(name string, schema string, applicationPath string, noDisplay bool) error {
	return registerURLHandler(name, schema, applicationPath)
}

func UnregisterURLHandler(applicationPath string) error {
	return unregisterURLHandler(applicationPath)
}
