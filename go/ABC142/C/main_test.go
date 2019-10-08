package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.Equal(t, "3 1 2", logic([]int{2, 3, 1}))
	assert.Equal(t, "1 2 3 4 5", logic([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, "8 2 4 5 6 7 3 1", logic([]int{8, 2, 7, 3, 4, 5, 6, 1}))
}
