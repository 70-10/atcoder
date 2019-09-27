package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.Equal(t, 3, logic([]int{10, 8, 8, 6}))
	assert.Equal(t, 1, logic([]int{15, 15, 15}))
	assert.Equal(t, 4, logic([]int{50, 30, 50, 100, 50, 80, 30}))
}
