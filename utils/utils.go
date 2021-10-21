package utils

import "strings"

// StringContainedInSlice ...
func StringContainedInSlice(a string, list []string) (bool, string) {
	for _, b := range list {
		if strings.Contains(a, b) {
			return true, b
		}
	}
	return false, ""
}
