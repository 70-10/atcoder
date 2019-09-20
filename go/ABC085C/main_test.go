package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	var a, b, c int
	a, b, c = logic(9, 45000)
	assert.Equal(t, 4, a)
	assert.Equal(t, 0, b)
	assert.Equal(t, 5, c)

	a, b, c = logic(20, 196000)
	assert.Equal(t, -1, a)
	assert.Equal(t, -1, b)
	assert.Equal(t, -1, c)

	a, b, c = logic(1000, 1234000)
	assert.Equal(t, 14, a)
	assert.Equal(t, 27, b)
	assert.Equal(t, 959, c)

	a, b, c = logic(2000, 20000000)
	assert.Equal(t, 2000, a)
	assert.Equal(t, 0, b)
	assert.Equal(t, 0, c)
}
