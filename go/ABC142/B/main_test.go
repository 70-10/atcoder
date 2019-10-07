package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.Equal(t, 2, logic([]int{150, 140, 100, 200}, 150))
	assert.Equal(t, 0, logic([]int{499}, 500))
	assert.Equal(t, 5, logic([]int{100, 200, 300, 400, 500}, 1))
}
