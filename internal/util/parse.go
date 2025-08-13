package util

import "strconv"

func ParseIntDefault(v string, def int) int {
	if i, err := strconv.Atoi(v); err == nil && i > 0 {
		return i
	}
	return def
}
func ParseUintDefault(v string, def uint64) uint64 {
	if i, err := strconv.ParseUint(v, 10, 64); err == nil {
		return i
	}
	return def
}
