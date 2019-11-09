package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		k      int
		s      string
		expect string
	}{
		{k: 1, s: "ABC", expect: "aBC"},
		{k: 3, s: "CABA", expect: "CAbA"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.k, test.s))
	}
}
