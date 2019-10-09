package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	assert.Equal(t, "Cloudy", logic("Sunny"))
	assert.Equal(t, "Rainy", logic("Cloudy"))
	assert.Equal(t, "Sunny", logic("Rainy"))
}
