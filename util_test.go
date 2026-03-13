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

func Test_escapeStrVal(t *testing.T) {
	input := "No special characters here."
	expected := "No special characters here."
	result := escapeStrVal(input)
	assert.Equal(t, expected, result)

	input = `New line\n\t
	and tab	characters.`
	expected = "New line\\n\\t\\n\tand tab\tcharacters."
	result = escapeStrVal(input)
	assert.Equal(t, expected, result)

	input = "New line\nand tab\tcharacters."
	expected = "New line\\nand tab\tcharacters."
	result = escapeStrVal(input)
	assert.Equal(t, expected, result)

	input = `He said, "Hello, World!"`
	expected = `He said, \"Hello, World!\"`
	result = escapeStrVal(input)
	assert.Equal(t, expected, result)
}
