package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.True(t, logic("erasedream"))
	assert.True(t, logic("dreameraser"))
	assert.False(t, logic("dreamerer"))
}
