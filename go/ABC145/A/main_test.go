package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		r      int
		expect int
	}{
		{r: 2, expect: 4},
		{r: 100, expect: 10000},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.r))
	}
}
