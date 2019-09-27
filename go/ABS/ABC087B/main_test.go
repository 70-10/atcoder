package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.Equal(t, 2, logic(2, 2, 2, 100))
	assert.Equal(t, 0, logic(5, 1, 0, 150))
	assert.Equal(t, 213, logic(30, 40, 50, 6000))
}
