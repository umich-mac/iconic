package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashFile(t *testing.T) {
	hash, err := HashFile("test.png")
	assert.Equal(t, hash, "2f5bbd5c5147380eecd94daa0cf0681c30ca43803f10868a2b9a2aedd26b0f21")
	assert.Nil(t, err)
}
