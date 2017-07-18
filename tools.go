package main

import "strconv"

// parseBool parses a string to a bool.
func parseBool(value string) bool {
	boolValue, err := strconv.ParseBool(value)
	return err == nil && boolValue
}
