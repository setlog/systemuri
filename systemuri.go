package systemuri

// RegisterURLHandler registers a custom URL handler for the specified schema and application.
func RegisterURLHandler(name string, schema string, applicationPath string) error {
	return registerURLHandler(name, schema, applicationPath)
}

func UnregisterURLHandler(schema string) error {
	return unregisterURLHandler(applicationPath)
}
