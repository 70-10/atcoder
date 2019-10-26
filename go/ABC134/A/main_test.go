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
		{r: 4, expect: 48},
		{r: 15, expect: 675},
		{r: 80, expect: 19200},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.r))
	}
}
