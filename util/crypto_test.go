package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5Checksum(t *testing.T) {
	str := MD5Checksum("C:/testing/user2.txt")
	fmt.Println(str)
	assert.Equal(t, "d6cd6beda6171a7de3c86a4a243079fb", str, "Must be equals")
}