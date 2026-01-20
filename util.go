package ent

import (
	"strings"
)

func IsNebulaRetryableError(err error) bool {
	s := strings.ToLower(err.Error())
	return strings.Contains(s, "storage error: more than one request trying to")
}
