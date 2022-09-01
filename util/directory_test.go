package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	strOrigin = "test/path"
	strExpect = "test/path/"
)
func TestAddEscapeInLastDir(t *testing.T) {
	assert.Equal(t, strExpect, AddEscapeInLastDir(strOrigin), "Must be equals")
}