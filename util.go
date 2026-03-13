package ent

import (
	"strings"
)

func IsNebulaRetryableError(err error) bool {
	s := strings.ToLower(err.Error())
	return strings.Contains(s, "storage error: more than one request trying to")
}

func escapeStrVal(s string) string {
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, "\n", "\\n")
	return s
}
