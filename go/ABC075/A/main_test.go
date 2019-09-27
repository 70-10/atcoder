package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.Equal(t, 7, logic(5, 7, 5))
	assert.Equal(t, 7, logic(1, 1, 7))
	assert.Equal(t, -100, logic(-100, 100, 100))
}
