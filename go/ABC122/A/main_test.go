package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		s      string
		expect string
	}{
		{s: "A", expect: "T"},
		{s: "G", expect: "C"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.s))
	}
}
