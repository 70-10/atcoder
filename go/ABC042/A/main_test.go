package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.True(t, logic(5, 5, 7))
	assert.False(t, logic(7, 7, 5))
}
