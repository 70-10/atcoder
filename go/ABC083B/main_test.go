package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.Equal(t, 84, logic(20, 2, 5))
	assert.Equal(t, 13, logic(10, 1, 2))
	assert.Equal(t, 4554, logic(100, 4, 16))
}
