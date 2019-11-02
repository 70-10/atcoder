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
		{s: "ASSA", expect: "Yes"},
		{s: "STOP", expect: "No"},
		{s: "FFEE", expect: "Yes"},
		{s: "FREE", expect: "No"},
		{s: "FFFE", expect: "No"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.s))
	}
}
