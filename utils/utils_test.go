package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetContents(t *testing.T) {
	b := Box{Contents: "hello"}
	b.SetContents("world")
	assert.Equal(t, "world", b.Contents)
}
