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
		{s: "abcabc", expect: "Yes"},
		{s: "abcadc", expect: "No"},
		{s: "z", expect: "No"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.s))
	}
}
