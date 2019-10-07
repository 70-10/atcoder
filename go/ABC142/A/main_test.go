package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.Equal(t, 0.5000000000, logic(4))
	assert.Equal(t, 0.6000000000, logic(5))
	assert.Equal(t, 1.0000000000, logic(1))
}
