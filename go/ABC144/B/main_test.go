package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		a      int
		expect string
	}{
		{a: 10, expect: "Yes"},
		{a: 50, expect: "No"},
		{a: 81, expect: "Yes"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.a))
	}
}
