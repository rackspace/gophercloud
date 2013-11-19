package gophercloud

import (
	"strings"
)

// cimap is a map with case-insensitive keys.
// This map is used to support custom metadata retrieval.
type cimap struct {
	m map[string]string
}

// rawMap() returns the raw map implementing the cimap type.  Use with care.
func (m cimap) rawMap() map[string]string {
	return m.m
}

// get() returns two data: the value mapped from the corresponding key, and
// a boolean indication of whether or not the key existed at all.
func (m cimap) get(key string) (string, bool) {
	lowercaseKey := strings.ToLower(key)
	value, ok := m.m[lowercaseKey]
	return value, ok
}
