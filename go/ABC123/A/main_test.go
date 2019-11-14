package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		n      []int
		k      int
		expect string
	}{
		{n: []int{1, 2, 3, 8, 9}, k: 15, expect: "Yay!"},
		{n: []int{15, 18, 26, 35, 36}, k: 18, expect: ":("},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.n, test.k))
	}
}
