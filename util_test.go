package ent

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsNebulaRetryableError(t *testing.T) {
	e := errors.New("no")
	assert.False(t, IsNebulaRetryableError(e))
	e = errors.New("Storage Error: More than one request trying to add/update/delete one edge/vertex at the same time.")
	assert.True(t, IsNebulaRetryableError(e))
}
