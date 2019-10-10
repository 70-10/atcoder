package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		expect string
		actual bool
	}{
		{expect: "RUDLUDR", actual: true},
		{expect: "DULL", actual: false},
		{expect: "UUUUUUUUUUUUUUU", actual: true},
		{expect: "ULURU", actual: false},
		{expect: "RDULULDURURLRDULRLR", actual: true},
	}

	for _, test := range tests {
		assert.Equal(t, logic(test.expect), test.actual, test.expect)
	}
}
