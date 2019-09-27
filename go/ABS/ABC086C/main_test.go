package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.True(t, logic([]int{0, 3, 6}, []int{0, 1, 1}, []int{0, 2, 1}))
	assert.False(t, logic([]int{0, 2}, []int{0, 100}, []int{0, 100}))
	assert.False(t, logic([]int{0, 5, 100}, []int{0, 1, 1}, []int{0, 1, 1}))
}
