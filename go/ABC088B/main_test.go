package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.Equal(t, 2, logic([]int{3, 1}))
	assert.Equal(t, 5, logic([]int{2, 7, 4}))
	assert.Equal(t, 18, logic([]int{20, 18, 2, 18}))
}
